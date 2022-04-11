/*
Worker.
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"net"
	"strconv"
	"strings"
	"sync"
	"time"

	rpc2 "github.com/miningmeter/rpc2"
	"github.com/miningmeter/rpc2/stratumrpc"
)

/*
Worker - it is worker. It's all.
*/
type Worker struct {
	mutex      sync.RWMutex
	ua         string                 // User Agent.
	id         string                 // ID.
	addr       string                 // ip:port.
	user       string                 // User.
	hash       string                 // Name of hash.
	divider    float64                // Divider of hash for computing of hashrate.
	difficulty float64                // Difficulty of job.
	window     map[int64]float64      // Shares counter window.
	client     *rpc2.Client           // Pointer on connection to worker.
	extensions map[string]interface{} // Extensions of the worker.
	pool       WorkerPool
}

type WorkerPool struct {
	addr            string                 // ip:port.
	user            string                 // User.
	password        string                 // Password.
	subscription    string                 // Notify-id of connection to pool.
	ua              string                 // User Agent, that sended to pool.
	extranonce1     string                 // Extranonce1of connection to pool.
	extranonce2size int                    // Size of Extranonce2.
	client          *rpc2.Client           // Pointer on connection to pool.
	extensions      map[string]interface{} // Extensions of the pool.
	job             []interface{}          // Last job from pool.
}

/*
GetAddr - getting addr of the worker.

@return string addr
*/
func (w *Worker) GetAddr() string {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.addr
}

/*
GetID -  getting id of the worker.

@return string id
*/
func (w *Worker) GetID() string {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.id
}

/*
Init - init of the worker.

@return error
*/
func (w *Worker) Init(client *rpc2.Client) error {
	if w.GetID() != "" {
		return nil
	}
	if client == nil {
		return errors.New("unable link with non-exist rpc2 client")
	}

	extensions := make(map[string]interface{})
	for _, v := range sExtensions {
		extensions[v] = false
	}

	// Generate new id.
	h := md5.New()
	h.Write([]byte(w.addr + fmt.Sprint(time.Now().Unix())))

	w.mutex.Lock()
	id := h.Sum(nil)
	w.id = hex.EncodeToString(id[0:8])
	w.extensions = extensions
	w.client = client
	// Send high difficulty to the worker. Else on slow Auth worker send to proxy big amount
	// of shares with low difficulty. The proxy will reject these shares and worker will disconnect.
	w.difficulty = 2097152.0
	w.pool.addr = workers.poolAddr
	w.pool.extranonce2size = 8
	w.pool.extensions = make(map[string]interface{})
	// w.pool.ua = ""
	w.mutex.Unlock()

	workers.Add(w)

	w.ResetHashrate()
	go w.UpdateHashrate()

	return nil
}

func (w *Worker) Reset(username string, pass string, pool string) {

	us := &User{user: username, password: pass}
	w.mutex.Lock()
	w.pool.password = pass
	w.pool.user = username
	w.pool.addr = pool
	w.pool.client.Close()
	w.user = us.GetName()
	w.mutex.Unlock()
}

/*
Auth - Auth request.

@param string user username of the worker.
@param string password password of the worker. Not using.

@return error
*/
func (w *Worker) Auth(user, password string) error {
	// us, err := db.GetUser(user)
	// if err != nil {
	// 	return err
	// }
	if user == "" {
		err := fmt.Errorf("invalid User: '%v'", user)
		LogError("%w", w.GetID(), err)

	}

	us := &User{}

	LogInfo("Initialializing user... worker user: %v; pool user: %v; auth user: %v", w.id, w.user, w.pool.user, user)
	err := us.Init(workers.poolAddr, user, password)

	if err != nil {
		LogError("User Init error: %w", w.GetID(), err)
	}

	LogInfo("User initialized... %+v\n", w.id, us)

	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wUser := w.user
	pClient := w.pool.client
	w.mutex.RUnlock()

	LogInfo("before reauth - user: %v; name: %v; p", sID, wUser, us.GetName())
	reauth := w.user != "" && w.pool.user == user && workers.poolAddr == w.pool.addr && wUser != us.GetName()
	LogInfo("reauth: %v", sID, strconv.FormatBool(reauth))
	if reauth {
		LogInfo("%s : change session from user %s to user %s", sID, wAddr, wUser, us.name)
		if pClient != nil {
			w.DisconnectPool()
		}
	}

	w.mutex.Lock()

	if w.user == "" || reauth {
		w.user = us.GetName()
		w.pool.addr = us.pool
		w.pool.user = us.user
		w.pool.password = us.password
		w.hash = "sha256" //us.hash
		w.divider = 1.0   //us.divider
	}
	pClient = w.pool.client
	w.mutex.Unlock()

	LogInfo("reauth: %v; worker user after reauth: %s", sID, reauth, w.user)

	if pClient == nil {
		go w.Connect()
	}

	return nil
}

/*
Connect - connecting worker to the pool.

@return error
*/
func (w *Worker) Connect() error {
	var status bool

	w.mutex.RLock()

	sID := w.id
	wAddr := w.addr
	wUser := w.user
	// wDivider := w.divider
	// wHash := w.hash
	pAddr := w.pool.addr
	pClient := w.pool.client
	pUser := w.pool.user
	pPassword := w.pool.password
	pSubscription := w.pool.subscription
	w.mutex.RUnlock()

	pUA := "miningmeter-proxy/1.0"

	// Check existence of the connection.
	//LogInfo("%s : check connect", sID, pAddr)
	if pClient != nil {
		//LogInfo("%s : already connected.", sID, pAddr)
		return nil
	}

	LogInfo("connecting to the pool - thread unsafe pool info: %+v\n; threadsafe pool address: %v; user: %v", sID, w.pool, pAddr, wUser)

	// Connecting to the pool.
	conn, err := net.DialTimeout("tcp", pAddr, 5*time.Second)

	if err != nil {
		LogError("%s : connection error: %s", sID, pAddr, err.Error())
		w.Disconnect()
		return fmt.Errorf("connection error: %s", err.Error())
	}

	// Init JSON-RPC client.
	client := rpc2.NewClientWithCodec(stratumrpc.NewStratumCodec(conn))

	client.Handle("mining.notify", mining.Notify)
	client.Handle("mining.set_difficulty", mining.SetDifficulty)

	// Linking worker to connection.
	state := rpc2.NewState()
	state.Set("worker", w)
	client.State = state

	// Run client in other goroutine.
	go client.Run()

	// Linking pool to the connection.
	w.mutex.Lock()
	w.pool.client = client
	w.mutex.Unlock()

	// Starting monitoring of the connection.
	go w.DisconnectNotify()

	var params []interface{}
	var reply []interface{}

	// Sending subscribe command to the pool.
	msg := MiningSubscribeRequest{pUA, pSubscription}
	params, err = msg.Encode()
	if err != nil {
		LogError("%s : %s", sID, pAddr, err.Error())
		w.Disconnect()
		return err
	}
	if pSubscription == "" {
		LogInfo("%s < mining.subscribe: %s", sID, pAddr, pUA)
	} else {
		LogInfo("%s < mining.subscribe: %s, %s", sID, pAddr, pUA, pSubscription)
	}
	err = client.Call("mining.subscribe", params, &reply)
	if err != nil {
		LogError("%s > mining.subscribe error: %s", sID, pAddr, err.Error())
		w.Disconnect()
		return err
	}
	response := new(MiningSubscribeResponse)
	err = response.Decode(reply)
	if err != nil {
		LogError("%s : %s", sID, err.Error())
		w.Disconnect()
		return err
	}
	w.mutex.Lock()
	w.pool.subscription = response.subscriptions["mining.notify"]
	w.pool.extranonce1 = response.extranonce1
	w.pool.extranonce2size = response.extranonce2size
	w.mutex.Unlock()

	LogInfo("%s > mining.subscribe: %s, %s, %d", sID, pAddr, response.subscriptions["mining.notify"], response.extranonce1, response.extranonce2size)

	// Sending authorize command to the pool.
	msgAuth := MiningAuthorizeRequest{pUser, pPassword}
	params, err = msgAuth.Encode()
	if err != nil {
		LogError("%s : %s", sID, pAddr, err.Error())
		w.Disconnect()
		return err
	}
	var breply bool
	LogInfo("%s < mining.authorize: %s, %s", sID, pAddr, pUser, pPassword)
	err = client.Call("mining.authorize", params, &breply)
	if err != nil {
		LogError("%s > mining.authorize error: %s", sID, pAddr, err.Error())
		w.Disconnect()
		return err
	}
	LogInfo("%s > mining.authorize: %t", sID, pAddr, breply)
	if !breply {
		LogError("%s : access denied to pool", sID, pAddr)
		w.Disconnect()
		return err
	}

	// Activating of the metrics.
	// mWorkerUp.WithLabelValues(tag, wAddr, wUser).Set(1)
	// mPoolUp.WithLabelValues(tag, wHash, pAddr).Set(1)
	// mPoolDivider.WithLabelValues(tag, wHash, pAddr).Set(wDivider)
	// mDifficulty.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Set(0)

	LogInfo("%s : sync extensions to pool %s", sID, wAddr, pAddr)
	status = w.SyncExtensions()
	if status {
		LogInfo("%s : extensions already synced", sID, wAddr)
	} else {
		LogInfo("%s : extensions not synced. Forcing reconnect.", sID, wAddr)
	}
	status = w.UpdateData(!status)

	if !status {
		return nil
	}

	LogInfo("%s : connected to pool %s", sID, wAddr, pAddr)

	return nil
}

/*
SyncExtensions - Syncing extensions between worker and pool.

@return bool sync status.
*/
func (w *Worker) SyncExtensions() bool {
	var reply interface{}

	extensions := make(map[string]interface{})

	w.mutex.RLock()
	sID := w.id
	a := w.pool.addr
	c := w.pool.client
	e := w.extensions
	pe := w.pool.extensions
	w.mutex.RUnlock()

	LogInfo("%s : sync extensions - worker: %+v\n; pool:%+v\n;", sID, a, e, pe)
	if len(pe) > 0 {
		return true
	}

	// Dirty hack. It removes from pool request the subscribe-extranonce extension.
	// Some pools incorrect processing of this extension. For example,
	// if this extension existing in mining. configures the pool viabtc.com not
	// responding to this request.
	he := make(map[string]interface{})
	for k, v := range e {
		if k != "subscribe-extranonce" {
			he[k] = v
		}
	}

	if c == nil {
		LogError("%s : connection closed unexpectedly", sID, a)
		return true
	}
	r := new(MiningConfigureRequest)
	r.extensions = he
	params, err := r.Encode()
	if err != nil {
		LogError("%s : encode extensions error: %s", sID, a, err.Error())
		return true
	}
	LogInfo("sync params: %+v\n ", sID, params)
	if params != nil {
		j, _ := json.Marshal(params)
		LogInfo("%s < mining.configure: %s", sID, a, j)
		err = c.Call("mining.configure", params, &reply)
		if err != nil {
			LogError("%s : sync extensions error: %s", sID, a, err.Error())
			return true
		}
		j, _ = json.Marshal(reply)
		LogInfo("%s > %s", sID, a, j)

		if reply == nil {
			LogInfo("%s : sync empty response. Set default values", sID, a)
			return true
		}

		re := new(MiningConfigureResponse)
		if err := re.Decode(reply); err != nil {
			LogError("%s : decode extensions error: %s", sID, a, err.Error())
			return true
		}

		LogInfo("%s : fix extensions with non-exist response", sID, a)
		for k, v := range re.extensions {
			extensions[k] = v
		}
		for k := range e {
			if _, ok := extensions[k]; !strings.Contains(k, ".") && !ok {
				extensions[k] = false
			}
		}

		j, _ = json.Marshal(extensions)
		LogInfo("%s > computed pool extensions: %s", sID, a, j)
	} else {
		LogInfo("%s : sync not required - params: %+v\n ", sID, a, params)
		extensions = e
	}

	w.mutex.Lock()
	w.pool.extensions = extensions
	w.mutex.Unlock()

	valid := 0
	for _, v := range sExtensions {
		if s, ok := extensions[v]; ok && (!s.(bool) || v == "subscribe-extranonce") {
			valid++
		}
	}
	_, ok := extensions["subscribe-extranonce"]
	if (len(extensions) == 1 && ok) || (len(extensions) == len(sExtensions) && len(sExtensions) == valid) {
		return true
	}

	return false
}

/*
UpdateData - updating of worker data.

@param bool force forced reconnection of worker.

@return bool if update is successfull.
*/
func (w *Worker) UpdateData(force bool) bool {
	w.mutex.RLock()
	sID := w.id
	a := w.addr
	c := w.pool.client
	e := w.pool.extranonce1
	e2 := w.pool.extranonce2size
	v, u := w.extensions["subscribe-extranonce"]
	w.mutex.RUnlock()
	LogInfo("worker subscribe-extranonce - exists: %v; value: %v", sID, u, v)
	u = u && v.(bool) && e != "" && !force

	if c == nil {
		LogError("%s : connection closed unexpectedly", sID, a)
		return false
	}

	if u {
		LogInfo("%s : update extranonce1: %s", sID, a, e)
		c.Notify("mining.set_extranonce", []interface{}{e, e2})
		LogInfo("%s < mining.set_extranonce: %s, %d", sID, a, e, e2)
	} else {

		LogInfo("%s : reconnect to proxy", sID, a)
		// w.Disconnect not requesting here, he will requested on closing connection.
		c.Close()

		return false
	}

	return true
}

/*
Restore - restoring connection to worker.

@param string id session id.

@return error
*/
func (w *Worker) Restore(id string) error {
	if status := ValidateHexString(id); !status {
		return fmt.Errorf("invalid format id = %s on worker restore", id)
	}

	worker, err := workers.Get(id)
	if err != nil {
		return err
	}
	if worker == nil {
		return nil
	}

	w.mutex.Lock()
	worker.mutex.Lock()

	temp := worker.addr
	worker.addr = w.addr
	w.addr = temp
	w.user = worker.user
	w.hash = worker.hash
	w.pool.addr = worker.pool.addr
	worker.client = w.client
	w.client = nil
	worker.client.State.Set("worker", worker)

	// wAddr := worker.addr
	// wUser := worker.user
	// wDifficulty := worker.difficulty
	// wHash := worker.hash
	// pAddr := worker.pool.addr

	worker.mutex.Unlock()
	w.mutex.Unlock()

	// Activating of metrics.
	// mWorkerUp.WithLabelValues(tag, wAddr, wUser).Set(1)
	// mDifficulty.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Set(wDifficulty)

	// <-time.After(1 * time.Minute)
	go w.Death()

	return nil
}

/*
ResetHashrate - resetting counter of shares.
*/
func (w *Worker) ResetHashrate() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.window = make(map[int64]float64)
}

/*
IncShares - incrementing counter of shares.
*/
func (w *Worker) IncShares() {
	w.mutex.RLock()
	d := w.difficulty
	dv := w.divider
	p := w.pool.client
	w.mutex.RUnlock()

	if p != nil {
		if d != 0 {
			stamp := time.Now().Unix()
			w.mutex.Lock()
			if s, ok := w.window[stamp]; ok {
				w.window[stamp] = s + (d / dv)
			} else {
				w.window[stamp] = d / dv
			}
			w.mutex.Unlock()
		}
	}
}

/*
ComputeHashrate - computing of worker hashrate.
*/
func (w *Worker) ComputeHashrate() float64 {
	wShares := 0.0

	stamp := time.Now().Unix() - 300 // 5 minutes.
	w.mutex.Lock()
	for i, s := range w.window {
		if i >= stamp {
			wShares += s
		} else {
			delete(w.window, stamp)
		}
	}
	w.mutex.Unlock()

	hashrate := (math.Pow(2, 32) * wShares) / 300

	return hashrate
}

/*
UpdateHashrate - updating of hashrate.
*/
func (w *Worker) UpdateHashrate() {
	var hashrate float64

	sID := w.GetID()
	wAddr := w.GetAddr()
	LogInfo("%s : init hashrate calculation", sID, wAddr)

	for {
		for i := 0; i <= 1; i++ {
			<-time.After(1 * time.Minute)
			if w == nil {
				break
			}

			hashrate = w.ComputeHashrate()
			if hashrate != 0 {
				break
			}
		}

		w.mutex.RLock()
		wAddr = w.addr
		// wUser := w.user
		// wHash := w.hash
		// wClient := w.client
		// pAddr := w.pool.addr
		w.mutex.RUnlock()

		// if wClient == nil {
		// 	// mSpeed.DeleteLabelValues(tag, wAddr, wUser, wHash, pAddr)
		// 	break
		// }
		// if pAddr != "" {
		// 	// mSpeed.WithLabelValues(tag, wAddr, wUser, wHash, pAddr).Set(hashrate)
		// }
		LogInfo("%s : hashrate: %.0f h/s", sID, wAddr, hashrate)

	}
	LogInfo("%s : stop hashrate calculation", sID, wAddr)
}

/*
DisconnectNotify - monitoring of connection status.
*/
func (w *Worker) DisconnectNotify() {
	w.mutex.RLock()
	sID := w.id
	a := w.pool.addr
	p := w.pool.client
	w.mutex.RUnlock()

	LogInfo("%s : set monitoring rpc connection", sID, a)
	<-p.DisconnectNotify()
	LogInfo("%s : rpc connection gone", sID, a)

	w.mutex.RLock()
	p = w.pool.client
	w.mutex.RUnlock()

	if p != nil {
		w.Disconnect()
	}
}

/*
Disconnect - disconnecting of worker.
*/
func (w *Worker) Disconnect() {
	w.mutex.Lock()
	sID := w.id
	wAddr := w.addr
	wClient := w.client

	w.client = nil
	w.mutex.Unlock()

	if wClient != nil {
		LogInfo("%s : disconnecting", sID, wAddr)

		wClient.Close()

		go w.Death()

		LogInfo("%s : disconnected from proxy", sID, wAddr)
	}
}

func (w *Worker) DisconnectNoWait() {
	w.mutex.Lock()
	sID := w.id
	wAddr := w.pool.addr
	wUser := w.user
	wClient := w.client

	w.client = nil
	w.mutex.Unlock()

	if wClient != nil {
		//LogInfo("%s : disconnecting", sID, wAddr)

		wClient.Close()

		go w.DeathNoWait()

		LogInfo("%s : disconnected from proxy; address: %v", sID, wUser, wAddr)
	}
}

/*
Death - preparing of worker delete.
*/
func (w *Worker) Death() {
	<-time.After(1 * time.Minute)

	// Removing of metrics.
	w.DeathNoWait()
}

func (w *Worker) DeathNoWait() {
	w.mutex.RLock()
	sID := w.id
	// wAddr := w.addr
	// wUser := w.user
	// wHash := w.hash
	wClient := w.client
	// pAddr := w.pool.addr
	w.mutex.RUnlock()

	if wClient == nil {
		workers.remove(sID)
		w.DisconnectPool()

		//LogInfo("%s : deleting metrics", sID, wAddr)

		// if pAddr != "" {

		// mSended.DeleteLabelValues(tag, wAddr, wUser, wHash, pAddr)
		// mOneSended.DeleteLabelValues(tag, wAddr, wUser, wHash, pAddr)
		// mAccepted.DeleteLabelValues(tag, wAddr, wUser, wHash, pAddr)
		// mOneAccepted.DeleteLabelValues(tag, wAddr, wUser, wHash, pAddr)
		// ok := mDifficulty.DeleteLabelValues(tag, wAddr, wUser, wHash, pAddr)
		// if !ok {
		// 	LogError("%s : error delete proxy_worker_difficulty metric", sID, pAddr)
		// }
		// ok = mWorkerUp.DeleteLabelValues(tag, wAddr, wUser)
		// if !ok {
		// 	LogError("%s : error delete proxy_worker_up metric.", sID, wAddr)
		// }
		// }
		//LogInfo("%s : removed from proxy", sID, wAddr)
	}
}

/*
DisconnectPool - disconnect pool.
*/
func (w *Worker) DisconnectPool() {
	w.mutex.RLock()
	sID := w.id
	// wHash := w.hash
	pAddr := w.pool.addr
	pClient := w.pool.client
	w.mutex.RUnlock()

	if pClient != nil {
		w.mutex.Lock()
		w.pool.client = nil
		w.pool.job = nil
		w.difficulty = 2097152.0
		w.mutex.Unlock()

		LogInfo("%s : disconnecting", sID, pAddr)
		// Closing of pool connection.
		pClient.Close()

		// The deleting of metrics.
		// ok := mPoolDivider.DeleteLabelValues(tag, wHash, pAddr)
		// if !ok {
		// 	LogError("%s : error delete proxy_pool_divider metric", sID, pAddr)
		// }
		// ok = mPoolUp.DeleteLabelValues(tag, wHash, pAddr)
		// if !ok {
		// 	LogError("%s : error delete proxy_pool_up metric", sID, pAddr)
		// }

		//LogInfo("%s : disconnected from proxy", sID, pAddr)
	}
}
