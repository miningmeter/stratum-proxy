/*
Методы для JSON-RPC сервера и клиента.
*/
package main

import (
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	rpc2 "github.com/miningmeter/rpc2"
)

/*
Mining - класс взаимодействия с майнером и пулом.
*/
type Mining struct{}

/*
Subscribe - обработка метода mining.subscribe от майнера. Подключает майнер к прокси.

@param *rpc2.Client  client клиент
@param []interface{} params переданные параметры
@param *interface{}  res    результат, отправляемый майнеру

@return error|nil ошибка
*/
func (*Mining) Subscribe(client *rpc2.Client, params []interface{}, res *interface{}) error {
	// Находим воркера, связанного с данным подключением.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	sID := w.GetID()
	wAddr := w.GetAddr()

	if _, err := mining.checkSubscribed(w); err == nil {
		LogError("%s : worker already subscribed", sID, wAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	Connect(client, w)
	sID = w.GetID()

	request := new(MiningSubscribeRequest)
	if err := request.Decode(params); err != nil {
		LogError("%s : invalid mining.subscribe: %s", sID, wAddr, err.Error())
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	if request.extranonce1 != "" {
		LogInfo("%s > mining.subscribe: %s, %s", sID, wAddr, request.ua, request.extranonce1)
		LogInfo("%s : restoring old connection", sID, wAddr)

		if err := w.Restore(request.extranonce1); err != nil {
			LogError("%s : restore error: %s", sID, wAddr, err.Error())
		}

		temp, _ = client.State.Get("worker")
		w = temp.(*Worker)
	} else {
		LogInfo("%s > mining.subscribe: %s", sID, wAddr, request.ua)
	}

	w.ua = request.ua

	w.mutex.RLock()
	wID := w.id
	wAddr = w.addr
	extranonce1 := w.pool.extranonce1
	extranonce2size := w.pool.extranonce2size
	w.mutex.RUnlock()

	// Отправляем майнеру mining.subscribe response.
	response := MiningSubscribeResponse{map[string]string{"mining.notify": wID}, extranonce1, extranonce2size}
	data, err := response.Encode()
	*res = &data
	if err != nil {
		LogError("%s : invalid mining.subscribe response: %s", sID, wAddr, err.Error())
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	if extranonce1 == "" {
		LogInfo("%s < %s %d", sID, wAddr, wID, extranonce2size)
	} else {
		LogInfo("%s < %s %s %d", sID, wAddr, wID, extranonce1, extranonce2size)
	}

	return nil
}

/*
Authorize - обработка метода mining.authorize от майнера. Получает данные,
к какому пулу подключить майнера, подключает к пулу, передает последнюю
сложность и работу майнеру.

@param *rpc2.Client  client клиент
@param []interface{} params переданные параметры
@param *bool         res    результат, отправляемый майнеру

@return error|nil ошибка
*/
func (*Mining) Authorize(client *rpc2.Client, params []interface{}, res *bool) error {
	*res = false
	// Получаем майнера, связанного с подключением.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wDifficulty := w.difficulty
	w.mutex.RUnlock()

	LogInfo("%s > mining.authorize", sID, wAddr)

	// Скрещивание ежа с ужом. Проверка сразу и подписки и повторной авторизации.
	sErr, err := mining.checkAuthorized(w)
	if err != nil && sID == "" {
		LogError("%s : authorize error: %s", sID, wAddr, err.Error())
		return errors.New(sErr)
	}

	// Отправляем майнеру большую сложность. Если этого не сделать, то при тормозах авторизации
	// и подключения майнер начнет заваливать прокси шарами с низкой сложностью и после некоторого
	// количества отбивок от прокси майнер отключается.
	client.Notify("mining.set_difficulty", wDifficulty)
	LogInfo("%s < mining.set_difficulty: %f", sID, wAddr, wDifficulty)

	auth := new(MiningAuthorizeRequest)
	auth.Decode(params)

	// Авторизуем майнера.
	err = w.Auth(auth.user, auth.password)
	if err != nil {
		LogError("%s < false: %s", sID, wAddr, err.Error())
		w.Disconnect()
		return err
	}

	*res = true
	LogInfo("%s < true", sID, wAddr)

	return nil
}

/*
Submit - обработка метода mining.submit от майнера. Принимает от майнера
найденный nonce и отправляет пулу. Возвращает майнеру статус валидации.

@param *rpc2.Client  client клиент
@param []interface{} params переданные параметры
@param *bool         res    результат, отправляемый майнеру

@return error|nil ошибка
*/
func (*Mining) Submit(client *rpc2.Client, params []interface{}, res *bool) error {
	var sErr error
	*res = false
	// Получаем майнера, связанного с подключением.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)

	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wUser := w.user
	wHash := w.hash
	wExt := w.extensions
	pAddr := w.pool.addr
	pUser := w.pool.user
	pClient := w.pool.client
	pExt := w.pool.extensions
	w.mutex.RUnlock()

	if sErr, err := mining.checkAuthorized(w); err != nil {
		LogError("%s : ignore share from %s", sID, wAddr, err.Error())
		return errors.New(sErr)
	}

	// получаем данные майнера и пула.
	if pAddr == "" {
		LogError("%s : ignore share without pool", sID, wAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	if pClient == nil {
		LogError("%s : ignore share to disconnected pool", sID, wAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	s := new(MiningSubmitRequest)
	s.Decode(params)
	isRoll := s.versionbits != ""

	if isRoll {
		LogInfo("%s > mining.submit: %s, %s, %s", sID, wAddr, s.job, s.nonce, s.versionbits)
	} else {
		LogInfo("%s > mining.submit: %s, %s", sID, wAddr, s.job, s.nonce)
	}

	// Проверяем шару на совместимость с расширениями.
	wRoll, wIsRoll := wExt["version-rolling"]
	pRoll, pIsRoll := pExt["version-rolling"]
	if isRoll && (!wIsRoll || !wRoll.(bool)) {
		LogError("%s : ignore share from miner without version rolling", sID, wAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}
	if isRoll && (!pIsRoll || !pRoll.(bool)) {
		LogError("%s : ignore share to pool without version rolling", sID, pAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}
	if !isRoll && (wIsRoll && wRoll.(bool)) {
		LogError("%s : ignore share from miner with version rolling", sID, wAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}
	if !isRoll && (pIsRoll && pRoll.(bool)) {
		LogError("%s : ignore share to pool with version rolling", sID, pAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	params[0] = pUser
	pClient.Call("mining.submit", params, res)

	LogInfo("%s < mining.submit: %s, %s", sID, pAddr, s.job, s.nonce)

	// Если работа валидна по мнению пула - увеличиваем счетчик
	// принятых шар.
	if *res {
		LogInfo("%s > %s", sID, pAddr, strconv.FormatBool(*res))
	} else {
		LogError("%s > %s", sID, pAddr, strconv.FormatBool(*res))
	}

	// Увеличиваем счетчик отправленных шар.
	mSended.WithLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr).Inc()

	// Если работа валидна - увеличиваем счетчик валидных шар, принятых
	// прокси.
	if *res {
		mAccepted.WithLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr).Inc()
		w.IncShares()
		LogInfo("%s < %s", sID, wAddr, strconv.FormatBool(*res))
	} else {
		LogError("%s < %s", sID, wAddr, strconv.FormatBool(*res))
	}

	return sErr
}

/*
ExtranonceSubscribe - обработка метода mining.extranonce.subscribe от майнера.

@param *rpc2.Client  client клиент
@param []interface{} params переданные параметры
@param *bool  res    результат, отправляемый майнеру

@return error|nil ошибка
*/
func (*Mining) ExtranonceSubscribe(client *rpc2.Client, params []interface{}, res *bool) error {
	*res = true
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	sID := w.GetID()
	wAddr := w.GetAddr()

	LogInfo("%s > mining.extranonce.subscribe", sID, wAddr)

	if sErr, err := mining.checkAuthorized(w); err != nil {
		LogError("%s : extranonce.subscribe error: %s", sID, wAddr, err.Error())
		return errors.New(sErr)
	}

	LogInfo("%s < true", sID, wAddr)

	// Взводим соответствующий флаг у майнера.
	w.mutex.Lock()
	if _, ok := w.extensions["subscribe-extranonce"]; ok {
		w.extensions["subscribe-extranonce"] = true
	}
	w.mutex.Unlock()

	return nil
}

/*
Notify - обработка метода mining.notify от пула. Принимает от пула job.

@param *rpc2.Client   client клиент
@param []interface{}  params переданные параметры
@param *[]interface{} res    ответ, отправляемый пулу

@return error|nil ошибка
*/
func (*Mining) Notify(client *rpc2.Client, params []interface{}, res *interface{}) error {
	// Получаем пул, связанный с соединением.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	jobID := params[0].(string)

	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wClient := w.client
	pAddr := w.pool.addr
	w.mutex.RUnlock()

	LogInfo("%s > mining.notify: %s", sID, pAddr, jobID)
	if wClient != nil {
		LogInfo("%s < mining.notify: %s", sID, wAddr, jobID)
		wClient.Notify("mining.notify", params)
	}

	return nil
}

/*
SetDifficulty - обработка метода mining.set_difficulty от пула. Принимает от пула
сложность и раздает его подключенным майнерам.

@param *rpc2.Client   client     клиент
@param []interface{}  params     переданные параметры
@param *[]interface{} res ответ, отправляемый пулу

@return error|nil ошибка
*/
func (*Mining) SetDifficulty(client *rpc2.Client, params []interface{}, res *interface{}) error {
	// Получаем пул, связанный с соединением.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	difficulty := params[0].(float64)

	// Получаем список майнеров, подключенных к пулу, и алгоритм.
	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wUser := w.user
	wDifficulty := w.difficulty
	wHash := w.hash
	pAddr := w.pool.addr
	w.mutex.RUnlock()

	// Устанавливаем метрику сложность на соответствующем пуле и
	// алгоритме.

	if wDifficulty != difficulty {
		LogInfo("%s > mining.set_difficulty: %f", sID, pAddr, difficulty)
		// Сохраняем сложность.
		w.mutex.Lock()
		w.difficulty = difficulty
		wClient := w.client
		w.mutex.Unlock()

		if wClient != nil {
			wClient.Notify("mining.set_difficulty", params)
			LogInfo("%s < mining.set_difficulty: %f", sID, wAddr, difficulty)
			mDifficulty.WithLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr).Set(difficulty)
		}
	}

	return nil
}

/*
Configure - обработка метода mining.configure от майнера. Передает пулу поддерживаемые расширения.

@param *rpc2.Client  client клиент
@param []interface{} params переданные параметры
@param *interface{}  res    результат, отправляемый майнеру

@return error|nil ошибка
*/
func (*Mining) Configure(client *rpc2.Client, params []interface{}, res *interface{}) error {
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)

	Connect(client, w)
	sID := w.GetID()
	wAddr := w.GetAddr()

	LogInfo("%s > mining.configure", sID, wAddr)

	e := new(MiningConfigureRequest)
	if err := e.Decode(params); err != nil {
		LogError("%s : invalid mining.configure: %s", sID, wAddr, err.Error())
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	old := make(map[string]interface{})
	w.mutex.Lock()
	for k, v := range w.extensions {
		old[k] = v
		delete(w.extensions, k)
	}
	for k, v := range e.extensions {
		w.extensions[k] = v
	}
	w.mutex.Unlock()

	for ke := range e.extensions {
		if _, ok := old[ke]; !ok {
			if strings.Contains(ke, ".") {
				parts := strings.Split(ke, ".")
				if _, ok := old[parts[0]]; !ok {
					delete(e.extensions, ke)
				}
			} else {
				e.extensions[ke] = false
			}
		}
	}

	a := new(MiningConfigureResponse)
	a.extensions = w.pool.extensions

	data, err := a.Encode()
	*res = &data
	if err != nil {
		LogError("%s : invalid mining.configure response: %s", sID, wAddr, err.Error())
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	j, _ := json.Marshal(data)
	LogInfo("%s < %s", sID, wAddr, j)

	return nil
}

func (*Mining) checkSubscribed(w *Worker) (string, error) {
	if w.GetID() == "" {
		return "[25, \"Not subscribed\", null]", errors.New("unsubscribed worker")
	}

	return "", nil
}

func (*Mining) checkAuthorized(w *Worker) (string, error) {
	if str, err := mining.checkSubscribed(w); err != nil {
		return str, err
	}
	w.mutex.RLock()
	user := w.user
	w.mutex.RUnlock()
	if user == "" {
		return "[24, \"Unauthorized worker\", null]", errors.New("unauthorized worker")
	}

	return "", nil
}
