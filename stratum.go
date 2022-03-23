/*
Methods of stratum protocol.
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
MiningSubscribeRequest - subscribe request.
*/
type MiningSubscribeRequest struct {
	ua          string
	extranonce1 string
}

/*
Encode - encoding of subscribe request.
*/
func (s *MiningSubscribeRequest) Encode() ([]interface{}, error) {
	if s.extranonce1 != "" {
		if !ValidateHexString(s.extranonce1) {
			return nil, fmt.Errorf("invalid extranonce1 = %s in mining.subscribe", s.extranonce1)
		}
	}

	return []interface{}{s.ua, s.extranonce1}, nil
}

/*
Decode - decoding of subscribe request.
["cpuminer/2.3.2"]
*/
func (s *MiningSubscribeRequest) Decode(data []interface{}) error {
	count := len(data)
	if count == 0 || count > 2 {
		return errors.New("invalid params count in mining.subscribe")
	}
	if count > 0 {
		ua, ok := data[0].(string)
		if !ok {
			return errors.New("invalid useragent type in mining.subscribe")
		}
		s.ua = ua
	}

	if count > 1 {
		extranonce1, ok := data[1].(string)
		if !ok {
			return errors.New("invalid extranonce1 type in mining.subscribe")
		}
		if extranonce1 != "" && !ValidateHexString(extranonce1) {
			return fmt.Errorf("invalid param extranonce1 = %s in mining.subscribe", extranonce1)
		}
		s.extranonce1 = extranonce1
	}

	return nil
}

/*
MiningSubscribeResponse - subscribe response.
*/
type MiningSubscribeResponse struct {
	subscriptions   map[string]string
	extranonce1     string
	extranonce2size int
}

/*
Encode - encoding of subscribe response.
*/
func (s *MiningSubscribeResponse) Encode() ([]interface{}, error) {
	if s.extranonce1 != "" && !ValidateHexString(s.extranonce1) {
		return nil, fmt.Errorf("invalid extranonce1 = %s in mining.subscribe response", s.extranonce1)
	}
	if s.extranonce2size == 0 {
		return nil, errors.New("no extranonce2_size in mining.subscribe response")
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
Decode - decoding of subscribe response.
[[["mining.set_difficulty","deadbeefcafebabee20a000000000000"],["mining.notify","deadbeefcafebabee20a000000000000"]],"3ffffdfe",4]
[[["mining.notify","5c567cf8"]],"bf5c565c",4]
[["mining.notify", "ae6812eb4cd7735a302a8a9dd95cf71f"], "a006868b", 4]
*/
func (s *MiningSubscribeResponse) Decode(data []interface{}) error {
	count := len(data)
	if count != 3 {
		return errors.New("invalid count of params in mining.subscribe response")
	}
	subscriptions, ok := data[0].([]interface{})
	if !ok {
		return errors.New("invalid subscriptions type in mining.subscribe response")
	}
	count = len(subscriptions)
	if count == 0 {
		return errors.New("no subscriptions in mining.subscribe response")
	}
	// Correction of violation of subscribing message standard, using an array of subscription
	// parts instead of an array of subscriptions. If subscriptions array has a one element,
	// this element placing out of subscription array. Compare:
	// [["mining.notify","5c567cf8"]] - 1 element of the subscription array with standard.
	// ["mining.notify", "ae6812eb"] - 1 element of the subscription array without standard.
	_, ok = subscriptions[0].(string)
	if ok {
		LogInfo("proxy : stratum standart violation in mining.subscribe response. Incorrect subscription array. Correcting...", "")
		subscriptions = []interface{}{subscriptions}
	}
	s.subscriptions = make(map[string]string)

	miningMatcher, rErr := regexp.Compile(`^mining\..+$`)

	if rErr != nil {
		err := fmt.Errorf("regexp error: %s", rErr.Error())

		return err
	}

	for _, sb := range subscriptions {
		subscription, ok := sb.([]interface{})
		if !ok {
			return errors.New("invalid subscription type in mining.subscribe response")
		}
		count = len(subscription)
		if count != 2 {
			return errors.New("invalid subscription param count in mining.subscribe response")
		}
		name, ok := subscription[0].(string)
		if !ok {
			return errors.New("invalid subscription name type in mining.subscribe response")
		}
		value, ok := subscription[1].(string)
		if !ok {
			return errors.New("invalid subscription value type in mining.subscribe response")
		}
		matched := miningMatcher.MatchString(name)

		if !matched {
			err := fmt.Errorf("invalid subscription name = %s in mining.subscribe response", name)

			return err
		}
		s.subscriptions[name] = value
	}

	extranonce1, ok := data[1].(string)
	if !ok {
		return errors.New("invalid extranonce1 type in mining.subscribe response")
	}
	if extranonce1 != "" && (!ValidateHexString(extranonce1)) {
		return fmt.Errorf("invalid extranonce1 = %s in mining.subscribe response", extranonce1)
	}
	s.extranonce1 = extranonce1

	var extranonce2size int
	if d, ok := data[2].(float64); ok {
		extranonce2size = int(d)
	} else if d, ok := data[2].(int); ok {
		extranonce2size = d
	} else {
		return errors.New("invalid extranonce2_size type in mining.subscribe response")
	}
	if extranonce2size == 0 {
		return errors.New("invalid extranonce2_size in mining.subscribe response")
	}
	s.extranonce2size = extranonce2size

	return nil
}

/*
MiningConfigureRequest - configure request.

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
Encode - encoding of configure request.
*/
func (s *MiningConfigureRequest) Encode() ([]interface{}, error) {
	if s.extensions == nil {
		return nil, errors.New("no extensions in mining.configure request")
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
Decode - decoding of configure request.
*/
func (s *MiningConfigureRequest) Decode(data []interface{}) error {
	if len(data) != 2 {
		return errors.New("invalid count of params in mining.configure request")
	}
	extensions, ok := data[0].([]interface{})
	if !ok {
		return errors.New("invalid extensions type in mining.configure request")
	}
	if len(extensions) == 0 {
		return errors.New("no extensions in mining.configure request")
	}
	params, ok := data[1].(map[string]interface{})
	if !ok {
		return errors.New("invalid params type in mining.configure request")
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
		return errors.New("no extensions in mining.configure request")
	}

	return nil
}

/*
MiningConfigureResponse - configure response.

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
Encode - encoding of configure response.
*/
func (s *MiningConfigureResponse) Encode() (map[string]interface{}, error) {
	if s.extensions == nil {
		return nil, errors.New("no extensions in mining.configure response")
	}
	count := len(s.extensions)
	if count == 0 {
		return nil, errors.New("empty extensions array in mining.configure response")
	}

	return s.extensions, nil
}

/*
Decode - decoding of configure response.
*/
func (s *MiningConfigureResponse) Decode(data interface{}) error {
	result, ok := data.(map[string]interface{})
	if !ok {
		return errors.New("invalid data in mining.configure response")
	}
	if len(result) == 0 {
		return errors.New("empty data in mining.configure response")
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
		return errors.New("no extensions in mining.configure response")
	}

	return nil
}

/*
MiningAuthorizeRequest - authorize request.
*/
type MiningAuthorizeRequest struct {
	user     string
	password string
}

/*
Encode - encoding of authorize request.
*/
func (s *MiningAuthorizeRequest) Encode() ([]interface{}, error) {
	if s.user == "" {
		return nil, fmt.Errorf("empty user in mining.authorize")
	}
	return []interface{}{s.user, s.password}, nil
}

/*
Decode - decoding of authorize request.
["1.1CvpvjwJTV5ob6HCUtsA5QZUwTbSQCj6iG", "X"]
*/
func (s *MiningAuthorizeRequest) Decode(data []interface{}) error {
	count := len(data)
	if count != 2 {
		return errors.New("invalid params count in mining.authorize")
	}
	user, ok := data[0].(string)
	if !ok {
		return errors.New("invalid user type in mining.authorize")
	}
	if user == "" {
		return errors.New("empty user in mining.authorize")
	}
	s.user = user
	password, ok := data[1].(string)
	if !ok {
		return errors.New("invalid password type in mining.authorize")
	}
	s.password = password

	return nil
}

/*
MiningSubmitRequest - share submit request.
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
Encode - encoding of share.
*/
func (s *MiningSubmitRequest) Encode() ([]interface{}, error) {
	if s.user == "" {
		return nil, fmt.Errorf("empty user in mining.submit request")
	}
	if s.job == "" {
		return nil, fmt.Errorf("empty job in mining.submit request")
	}
	if !ValidateHexString(s.job) {
		return nil, fmt.Errorf("invalid job = %s in mining.submit request", s.job)
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
Decode - decoding of share.
["814d7ad2fc212263","125","28b9040000000000","5e590495","275d26e4"]
["81c0792ed5ccda13","B8vnoaHIg","ea72010000000000","5e590735","0cbe52ca"]
*/
func (s *MiningSubmitRequest) Decode(data []interface{}) error {
	count := len(data)
	if count < 5 || count > 6 {
		return errors.New("invalid params count in mining.submit request")
	}
	user, ok := data[0].(string)
	if !ok {
		return errors.New("invalid user type in mining.submit request")
	}
	if user == "" {
		return errors.New("empty user in mining.submit request")
	}
	s.user = user
	job, ok := data[1].(string)
	if !ok {
		return errors.New("invalid job type in mining.submit request")
	}
	if job == "" {
		return errors.New("empty job in mining.submit request")
	}
	s.job = job
	extranonce2, ok := data[2].(string)
	if !ok {
		return errors.New("invalid extranonce2 type in mining.submit request")
	}
	if !ValidateHexString(extranonce2) {
		return fmt.Errorf("invalid extranonce2 in mining.submit request")
	}
	s.extranonce2 = extranonce2
	ntime, ok := data[3].(string)
	if !ok {
		return errors.New("invalid ntime type in mining.submit request")
	}
	if err := s.validateDword(ntime, "ntime"); err != nil {
		return fmt.Errorf("%s in mining.submit request", err.Error())
	}
	s.ntime = ntime
	nonce, ok := data[4].(string)
	if !ok {
		return errors.New("invalid nonce type in mining.submit request")
	}
	if err := s.validateDword(nonce, "nonce"); err != nil {
		return fmt.Errorf("%s in mining.submit request", err.Error())
	}
	s.nonce = nonce
	s.versionbits = ""
	if count == 6 {
		versionbits, ok := data[5].(string)
		if !ok {
			return errors.New("invalid versionbits type in mining.submit request")
		}
		if err := s.validateDword(versionbits, "versionbits"); err != nil {
			return fmt.Errorf("%s in mining.submit request", err.Error())
		}
		s.versionbits = versionbits
	}

	return nil
}

/*
validateDword - validating of dword string representation.

@param string dword validated string
@param string name  name of validated value for error message

@return error
*/
func (*MiningSubmitRequest) validateDword(dword string, name string) error {
	if dword == "" {
		return fmt.Errorf("empty %s", name)
	}
	count := len(dword)
	if count != 8 {
		return fmt.Errorf("invalid length = %d of %s = %s", count, name, dword)
	}
	if !ValidateHexString(dword) {
		return fmt.Errorf("invalid %s = %s", name, dword)
	}

	return nil
}
