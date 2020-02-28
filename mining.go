/*
Handlers for JSON-RPC server and client.
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
Mining - the communicating with worker and pool.
*/
type Mining struct{}

/*
Subscribe - the handler of mining.subscribe from the worker. He's connecting worker to proxy.

@param *rpc2.Client  client
@param []interface{} params
@param *interface{}  res    result, sent to the worker

@return error|nil
*/
func (*Mining) Subscribe(client *rpc2.Client, params []interface{}, res *interface{}) error {
	// The getting of the worker, linked with the connection.
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

	// The sending mining.subscribe response message to the worker.
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
Authorize - the handler of mining.authorize from the worker. He's detect the pool
for connect to the worker, connecting to the pool, sending the last difficulty
and the job.

@param *rpc2.Client  client
@param []interface{} params
@param *bool         res    result, sent to the worker

@return error|nil
*/
func (*Mining) Authorize(client *rpc2.Client, params []interface{}, res *bool) error {
	*res = false
	// The getting of the worker, linked with the connection.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wDifficulty := w.difficulty
	w.mutex.RUnlock()

	LogInfo("%s > mining.authorize", sID, wAddr)

	// The checking of the subscription and the double authorization.
	sErr, err := mining.checkAuthorized(w)
	if err != nil && sID == "" {
		LogError("%s : authorize error: %s", sID, wAddr, err.Error())
		return errors.New(sErr)
	}

	// The sending of high difficulty to the worker. If this is not done, then with authorization and
	// connection brakes, the worker will begin to flood proxy with shares with low difficulty and after
	// a certain number of errors from the proxy, the worker will disconnect.
	client.Notify("mining.set_difficulty", wDifficulty)
	LogInfo("%s < mining.set_difficulty: %f", sID, wAddr, wDifficulty)

	auth := new(MiningAuthorizeRequest)
	auth.Decode(params)

	// The authorizing of the miner.
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
Submit - the handler of mining.submit from the worker. The taking the nonce from the worker and
sending it to the pool. Return to the worker a validation status.

@param *rpc2.Client  client
@param []interface{} params
@param *bool         res    result, sent to the worker

@return error|nil
*/
func (*Mining) Submit(client *rpc2.Client, params []interface{}, res *bool) error {
	var sErr error
	*res = false
	// The getting of the worker, linked with the connection.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)

	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wUser := w.user
	wHash := w.hash
	wDivider := w.divider
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

	// The getting of the data of the worker and the pool.
	if pAddr == "" {
		LogError("%s : ignore share without pool", sID, wAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	if pClient == nil {
		LogError("%s : ignore share to disconnected pool", sID, wAddr)
		return errors.New("[20, \"Other/Unknown\", null]")
	}

	s := new(MiningSubmitRequest)
	err := s.Decode(params)
	if err != nil {
		LogError("%s : %s", sID, wAddr, err.Error())
	}
	isRoll := s.versionbits != ""

	if isRoll {
		LogInfo("%s > mining.submit: %s, %s, %s", sID, wAddr, s.job, s.nonce, s.versionbits)
	} else {
		LogInfo("%s > mining.submit: %s, %s", sID, wAddr, s.job, s.nonce)
	}

	// The checking compatability of the share and the extensions of the worker.
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

	if *res {
		LogInfo("%s > %s", sID, pAddr, strconv.FormatBool(*res))
	} else {
		LogError("%s > %s", sID, pAddr, strconv.FormatBool(*res))
	}

	// The increasing the counter of the accepted shares.
	mSended.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Inc()
	w.mutex.RLock()
	wDifficulty := w.difficulty
	w.mutex.RUnlock()
	mOneSended.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Add(wDifficulty / wDivider)

	// If the pool has validated the work - we are increasing
	// the counter of the accepted shares.
	if *res {
		mAccepted.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Inc()
		mOneAccepted.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Add(wDifficulty / wDivider)
		w.IncShares()
		LogInfo("%s < %s", sID, wAddr, strconv.FormatBool(*res))
	} else {
		LogError("%s < %s", sID, wAddr, strconv.FormatBool(*res))
	}

	return sErr
}

/*
ExtranonceSubscribe - the handler of mining.extranonce.subscribe from the worker.

@param *rpc2.Client  client
@param []interface{} params
@param *bool  res           result, sent to the worker

@return error|nil
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

	// The adding a subscribe-extranonce extension on the worker.
	w.mutex.Lock()
	if _, ok := w.extensions["subscribe-extranonce"]; ok {
		w.extensions["subscribe-extranonce"] = true
	}
	w.mutex.Unlock()

	return nil
}

/*
Notify - the handler of mining.notify from the pool. He's receiving job.

@param *rpc2.Client
@param []interface{}
@param *[]interface{} res result, sent to the pool

@return error|nil
*/
func (*Mining) Notify(client *rpc2.Client, params []interface{}, res *interface{}) error {
	// The getting of the pool, linked with the connection.
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
SetDifficulty -  the handler of mining.set_difficulty from the pool. He's receiving difficulty
from the pool and sending it to the worker.

@param *rpc2.Client   client
@param []interface{}  params
@param *[]interface{} res    result, sent to the pool

@return error|nil
*/
func (*Mining) SetDifficulty(client *rpc2.Client, params []interface{}, res *interface{}) error {
	// The getting of the pool, linked with the connection.
	temp, _ := client.State.Get("worker")
	w := temp.(*Worker)
	difficulty := params[0].(float64)

	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wUser := w.user
	wDifficulty := w.difficulty
	wHash := w.hash
	pAddr := w.pool.addr
	w.mutex.RUnlock()

	// The saving of difficulty in the metrics.
	if wDifficulty != difficulty {
		LogInfo("%s > mining.set_difficulty: %f", sID, pAddr, difficulty)
		w.mutex.Lock()
		w.difficulty = difficulty
		wClient := w.client
		w.mutex.Unlock()

		// The sending of difficulty to the linked worker.
		if wClient != nil {
			wClient.Notify("mining.set_difficulty", params)
			LogInfo("%s < mining.set_difficulty: %f", sID, wAddr, difficulty)
			mDifficulty.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Set(difficulty)
		}
	}

	return nil
}

/*
Configure - the handler of mining.configure from the worker. He's sent to the pool
the extensions that the worker is supporting.

@param *rpc2.Client  client
@param []interface{} params
@param *interface{}  res    result, sent to the worker

@return error|nil
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
