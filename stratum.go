/*
Методы для работы со stratum протоколом.
*/
package main

import (
	"errors"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

/*
MiningSubscribeRequest - запрос подписки.
*/
type MiningSubscribeRequest struct {
	ua          string
	extranonce1 string
}

/*
Encode - кодирование запроса на подписку.
*/
func (s *MiningSubscribeRequest) Encode() ([]interface{}, error) {
	if s.extranonce1 != "" {
		if !ValidateHexString(s.extranonce1) {
			return nil, fmt.Errorf("Invalid extranonce1 = %s in mining.subscribe", s.extranonce1)
		}
	}

	return []interface{}{s.ua, s.extranonce1}, nil
}

/*
Decode - декодирование запроса на подписку.
["cpuminer/2.3.2"]
*/
func (s *MiningSubscribeRequest) Decode(data []interface{}) error {
	count := len(data)
	if count == 0 || count > 2 {
		return errors.New("Invalid params count in mining.subscribe")
	}
	if count > 0 {
		ua, ok := data[0].(string)
		if !ok {
			return errors.New("Invalid useragent type in mining.subscribe")
		}
		s.ua = ua
	}

	if count > 1 {
		extranonce1, ok := data[1].(string)
		if !ok {
			return errors.New("Invalid extranonce1 type in mining.subscribe")
		}
		if extranonce1 != "" && !ValidateHexString(extranonce1) {
			return fmt.Errorf("Invalid param extranonce1 = %s in mining.subscribe", extranonce1)
		}
		s.extranonce1 = extranonce1
	}

	return nil
}

/*
MiningSubscribeResponse - ответ на запрос подписки.
*/
type MiningSubscribeResponse struct {
	subscriptions   map[string]string
	extranonce1     string
	extranonce2size int
}

/*
Encode - кодирование ответа на запрос подписки.
*/
func (s *MiningSubscribeResponse) Encode() ([]interface{}, error) {
	if s.extranonce1 != "" && !ValidateHexString(s.extranonce1) {
		return nil, fmt.Errorf("Invalid extranonce1 = %s in mining.subscribe response", s.extranonce1)
	}
	if s.extranonce2size == 0 {
		return nil, errors.New("No extranonce2_size in mining.subscribe response")
	}

	var subscriptions []interface{}

	keys := make([]string, 0, len(s.subscriptions))
	for k := range s.subscriptions {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		subscriptions = append(subscriptions, []interface{}{k, s.subscriptions[k]})
	}

	return []interface{}{subscriptions, s.extranonce1, s.extranonce2size}, nil
}

/*
Decode - декодирование ответа на запрос подписки.
[[["mining.set_difficulty","deadbeefcafebabee20a000000000000"],["mining.notify","deadbeefcafebabee20a000000000000"]],"3ffffdfe",4]
[[["mining.notify","5c567cf8"]],"bf5c565c",4]
[["mining.notify", "ae6812eb4cd7735a302a8a9dd95cf71f"], "a006868b", 4]
*/
func (s *MiningSubscribeResponse) Decode(data []interface{}) error {
	count := len(data)
	if count != 3 {
		return errors.New("Invalid count of params in mining.subscribe response")
	}
	subscriptions, ok := data[0].([]interface{})
	if !ok {
		return errors.New("Invalid subscriptions type in mining.subscribe response")
	}
	count = len(subscriptions)
	if count == 0 {
		return errors.New("No subscriptions in mining.subscribe response")
	}
	// Коррекция нарушения стандарта subscribe, использование вместо массива подписок
	// массива частей подписок. То есть если в массиве подписок 1 элемент, этот
	// элемент выносится за пределы массива. Сравните:
	// [["mining.notify","5c567cf8"]] - 1 элемент массива подписок по стандарту.
	// ["mining.notify", "ae6812eb"] - 1 элемент массива подписок с нарушением стандарта.
	_, ok = subscriptions[0].(string)
	if ok {
		LogInfo("proxy : stratum standart violation in mining.subscribe response. Incorrect subscription array. Correcting...", "")
		subscriptions = []interface{}{subscriptions}
	}
	s.subscriptions = make(map[string]string)
	for _, sb := range subscriptions {
		subscription, ok := sb.([]interface{})
		if !ok {
			return errors.New("Invalid subscription type in mining.subscribe response")
		}
		count = len(subscription)
		if count != 2 {
			return errors.New("Invalid subscription param count in mining.subscribe response")
		}
		name, ok := subscription[0].(string)
		if !ok {
			return errors.New("Invalid subscription name type in mining.subscribe response")
		}
		value, ok := subscription[1].(string)
		if !ok {
			return errors.New("Invalid subscription value type in mining.subscribe response")
		}
		matched, rErr := regexp.MatchString("^mining\\..+$", name)
		if !matched || rErr != nil {
			err := fmt.Errorf("Invalid subscription name = %s in mining.subscribe response", name)
			if rErr != nil {
				err = fmt.Errorf("%s. Regexp error: %s", err.Error(), rErr.Error())
			}
			return err
		}
		s.subscriptions[name] = value
	}

	extranonce1, ok := data[1].(string)
	if !ok {
		return errors.New("Invalid extranonce1 type in mining.subscribe response")
	}
	if extranonce1 != "" && (!ValidateHexString(extranonce1)) {
		return fmt.Errorf("Invalid extranonce1 = %s in mining.subscribe response", extranonce1)
	}
	s.extranonce1 = extranonce1

	var extranonce2size int
	if d, ok := data[2].(float64); ok {
		extranonce2size = int(d)
	} else if d, ok := data[2].(int); ok {
		extranonce2size = d
	} else {
		return errors.New("Invalid extranonce2_size type in mining.subscribe response")
	}
	if extranonce2size == 0 {
		return errors.New("Invalid extranonce2_size in mining.subscribe response")
	}
	s.extranonce2size = extranonce2size

	return nil
}

/*
MiningConfigureRequest - запрос конфигурации майнером.

{"method": "mining.configure",
  "id": 1,
  "params": [["minimum-difficulty", "version-rolling"],
	     {"minimum-difficulty.value": 2048,
	      "version-rolling.mask": "1fffe000", "version-rolling.min-bit-count": 2}]}
*/
type MiningConfigureRequest struct {
	extensions map[string]interface{}
}

/*
Encode - кодирование запроса конфигурации.
*/
func (s *MiningConfigureRequest) Encode() ([]interface{}, error) {
	if s.extensions == nil {
		return nil, errors.New("No extensions in mining.configure request")
	}
	extensions := make([]string, 0)
	params := make(map[string]interface{})
	for ke, ve := range s.extensions {
		if !strings.Contains(ke, ".") {
			if ve.(bool) {
				extensions = append(extensions, ke)
			}
		} else {
			params[ke] = ve
		}
	}

	if len(extensions) == 0 {
		return nil, nil
	}

	return []interface{}{extensions, params}, nil
}

/*
Decode - декодирование запроса конфигурации.
*/
func (s *MiningConfigureRequest) Decode(data []interface{}) error {
	if len(data) != 2 {
		return errors.New("Invalid count of params in mining.configure request")
	}
	extensions, ok := data[0].([]interface{})
	if !ok {
		return errors.New("Invalid extensions type in mining.configure request")
	}
	if len(extensions) == 0 {
		return errors.New("No extensions in mining.configure request")
	}
	params, ok := data[1].(map[string]interface{})
	if !ok {
		return errors.New("Invalid params type in mining.configure request")
	}
	s.extensions = make(map[string]interface{})

	for _, ve := range extensions {
		extension, ok := ve.(string)
		if ok || len(extension) > 0 {
			s.extensions[extension] = true
		}
	}
	for kp, vp := range params {
		if !strings.Contains(kp, ".") {
			continue
		}
		chunks := strings.Split(kp, ".")
		if _, ok := s.extensions[chunks[0]]; len(chunks) == 2 && len(chunks[0]) > 0 && len(chunks[1]) > 0 || ok {
			s.extensions[kp] = vp
		}
	}
	if len(s.extensions) == 0 {
		s.extensions = nil
		return errors.New("No extensions in mining.configure request")
	}

	return nil
}

/*
MiningConfigureResponse - ответ на запрос конфигурации.

{"error": null,
  "id": 1,
  "result": {"version-rolling": true,
	     "version-rolling.mask": "18000000",
	     "minimum-difficulty": true}}
*/
type MiningConfigureResponse struct {
	extensions map[string]interface{}
}

/*
Encode - кодирование ответа на запрос конфигурации.
*/
func (s *MiningConfigureResponse) Encode() (map[string]interface{}, error) {
	if s.extensions == nil {
		return nil, errors.New("No extensions in mining.configure response")
	}
	count := len(s.extensions)
	if count == 0 {
		return nil, errors.New("Empty extensions array in mining.configure response")
	}

	return s.extensions, nil
}

/*
Decode - декодирование ответа на запрос конфигурации.
*/
func (s *MiningConfigureResponse) Decode(data interface{}) error {
	result, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("Invalid data in mining.configure response")
	}
	if len(result) == 0 {
		return errors.New("Empty data in mining.configure response")
	}
	s.extensions = make(map[string]interface{})
	for kp, vp := range result {
		if !strings.Contains(kp, ".") {
			s.extensions[kp] = vp
		} else {
			chunks := strings.Split(kp, ".")
			if _, ok := s.extensions[chunks[0]]; len(chunks) == 2 && len(chunks[0]) > 0 && len(chunks[1]) > 0 && ok {
				s.extensions[kp] = vp
			}
		}
	}
	if len(s.extensions) == 0 {
		s.extensions = nil
		return errors.New("No extensions in mining.configure response")
	}

	return nil
}

/*
MiningAuthorizeRequest - запрос авторизации.
*/
type MiningAuthorizeRequest struct {
	user     string
	password string
}

/*
Encode - кодирование запроса на авторизацию.
*/
func (s *MiningAuthorizeRequest) Encode() ([]interface{}, error) {
	if s.user == "" {
		return nil, fmt.Errorf("Empty user in mining.authorize")
	}
	return []interface{}{s.user, s.password}, nil
}

/*
Decode - декодирование запроса на авторизацию.
["1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG", "X"]
*/
func (s *MiningAuthorizeRequest) Decode(data []interface{}) error {
	count := len(data)
	if count != 2 {
		return errors.New("Invalid params count in mining.authorize")
	}
	user, ok := data[0].(string)
	if !ok {
		return errors.New("Invalid user type in mining.authorize")
	}
	if user == "" {
		return errors.New("Empty user in mining.authorize")
	}
	s.user = user
	password, ok := data[1].(string)
	if !ok {
		return errors.New("Invalid password type in mining.authorize")
	}
	s.password = password

	return nil
}

/*
MiningSubmitRequest - отправка шары.
*/
type MiningSubmitRequest struct {
	user        string
	job         string
	extranonce2 string
	ntime       string
	nonce       string
	versionbits string
}

/*
Encode - кодирование шары.
*/
func (s *MiningSubmitRequest) Encode() ([]interface{}, error) {
	if s.user == "" {
		return nil, fmt.Errorf("Empty user in mining.submit request")
	}
	if s.job == "" {
		return nil, fmt.Errorf("Empty job in mining.submit request")
	}
	if !ValidateHexString(s.job) {
		return nil, fmt.Errorf("Invalid job = %s in mining.submit request", s.job)
	}
	if err := s.validateDword(s.extranonce2, "extranonce2"); err != nil {
		return nil, fmt.Errorf("%s in mining.submit request", err.Error())
	}
	if err := s.validateDword(s.ntime, "ntime"); err != nil {
		return nil, fmt.Errorf("%s in mining.submit request", err.Error())
	}
	if err := s.validateDword(s.nonce, "nonce"); err != nil {
		return nil, fmt.Errorf("%s in mining.submit request", err.Error())
	}
	if s.versionbits != "" {
		if err := s.validateDword(s.versionbits, "versionbits"); err != nil {
			return nil, fmt.Errorf("%s in mining.submit request", err.Error())
		}
	}
	out := []interface{}{s.user, s.job, s.extranonce2, s.ntime, s.nonce}
	if s.versionbits != "" {
		out = append(out, s.versionbits)
	}

	return out, nil
}

/*
Decode - декодирование шары.
*/
func (s *MiningSubmitRequest) Decode(data []interface{}) error {
	count := len(data)
	if count < 5 || count > 6 {
		return errors.New("Invalid params count in mining.submit request")
	}
	user, ok := data[0].(string)
	if !ok {
		return errors.New("Invalid user type in mining.submit request")
	}
	if user == "" {
		return errors.New("Empty user in mining.submit request")
	}
	s.user = user
	job, ok := data[1].(string)
	if !ok {
		return errors.New("Invalid job type in mining.submit request")
	}
	if job == "" {
		return errors.New("Empty job in mining.submit request")
	}
	if !ValidateHexString(job) {
		return fmt.Errorf("Invalid job = %s in mining.submit request", job)
	}
	s.job = job
	extranonce2, ok := data[2].(string)
	if !ok {
		return errors.New("Invalid extranonce2 type in mining.submit request")
	}
	if err := s.validateDword(extranonce2, "extranonce2"); err != nil {
		return fmt.Errorf("%s in mining.submit request", err.Error())
	}
	s.extranonce2 = extranonce2
	ntime, ok := data[3].(string)
	if !ok {
		return errors.New("Invalid ntime type in mining.submit request")
	}
	if err := s.validateDword(ntime, "ntime"); err != nil {
		return fmt.Errorf("%s in mining.submit request", err.Error())
	}
	s.ntime = ntime
	nonce, ok := data[4].(string)
	if !ok {
		return errors.New("Invalid nonce type in mining.submit request")
	}
	if err := s.validateDword(nonce, "nonce"); err != nil {
		return fmt.Errorf("%s in mining.submit request", err.Error())
	}
	s.nonce = nonce
	s.versionbits = ""
	if count == 6 {
		versionbits, ok := data[5].(string)
		if !ok {
			return errors.New("Invalid versionbits type in mining.submit request")
		}
		if err := s.validateDword(versionbits, "versionbits"); err != nil {
			return fmt.Errorf("%s in mining.submit request", err.Error())
		}
		s.versionbits = versionbits
	}

	return nil
}

/*
validateDword - валидация строчного представления dword.

@param string dword валидируемое значение
@param string name  название значения для сообщения об ошибке

@return error при ошибке валидации
*/
func (*MiningSubmitRequest) validateDword(dword string, name string) error {
	if dword == "" {
		return fmt.Errorf("Empty %s", name)
	}
	count := len(dword)
	if count != 8 {
		return fmt.Errorf("Invalid length = %d of %s = %s", count, name, dword)
	}
	if !ValidateHexString(dword) {
		return fmt.Errorf("Invalid %s = %s", name, dword)
	}

	return nil
}
