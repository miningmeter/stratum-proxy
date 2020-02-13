/*
Класс воркера.
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
	"strings"
	"sync"
	"time"

	rpc2 "github.com/miningmeter/rpc2"
	"github.com/miningmeter/rpc2/stratumrpc"
)

/*
Worker - воркер.
*/
type Worker struct {
	mutex      sync.RWMutex
	ua         string                 // User Agent воркера.
	id         string                 // Идентификатор воркера.
	addr       string                 // ip:port воркера.
	user       string                 // Пользователь воркера.
	hash       string                 // Хэш воркера.
	divider    float64                // Делитель хэша для вычисления хэшрейта.
	difficulty float64                // Значение сложности.
	window     map[int64]float64      // Счетчик шар за пять минут.
	client     *rpc2.Client           // Указатель на соединение, которое обслуживает воркер.
	extensions map[string]interface{} // Поддерживаемые расширения.
	pool       struct {
		addr            string                 // IP:port пула.
		user            string                 // Имя пользователя пула.
		password        string                 // Пароль пула.
		subscription    string                 // Notify-id соединения.
		ua              string                 // User Agent передаваемый пулу.
		extranonce1     string                 // Extranonce1 соединения.
		extranonce2size int                    // Размер Extranonce2.
		client          *rpc2.Client           // Соединение, обслуживающее пул.
		extensions      map[string]interface{} // Поддерживаемые расширения
	}
}

/*
GetAddr -  получение адреса воркера.

@return string адрес воркера
*/
func (w *Worker) GetAddr() string {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.addr
}

/*
GetID -  получение идентификатора воркера.

@return string идентификатор воркера
*/
func (w *Worker) GetID() string {
	w.mutex.RLock()
	defer w.mutex.RUnlock()

	return w.id
}

/*
Init - инициализация воркера.

@return error если произошла ошибка
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

	// Генерируем идентификатор воркера.
	h := md5.New()
	h.Write([]byte(w.addr + string(time.Now().Unix())))

	w.mutex.Lock()
	id := h.Sum(nil)
	w.id = hex.EncodeToString(id[0:8])
	w.extensions = extensions
	w.client = client
	// Задаем майнеру большую сложность. Если этого не сделать, то при тормозах авторизации
	// и подключения майнер начнет заваливать прокси шарами с низкой сложностью и после некоторого
	// количества отбивок от прокси майнер отключается.
	w.difficulty = 2097152.0
	w.pool.extranonce2size = 4
	w.pool.extensions = make(map[string]interface{})
	w.mutex.Unlock()

	w.ResetHashrate()
	go w.UpdateHashrate()

	return nil
}

/*
Auth - запрос аутентификации.

@param string user имя пользователя воркера.
@param string password пароль пользователя воркера. Не
используется.

@return error Если произошла ошибка.
*/
func (w *Worker) Auth(user, password string) error {
	var pClient *rpc2.Client

	us, err := db.GetUser(user)
	if err != nil {
		return err
	}

	w.mutex.Lock()
	if w.user == "" {
		w.user = us.name
		w.pool.addr = us.pool
		w.pool.user = us.user
		w.pool.password = us.password
		w.hash = us.hash
		w.divider = us.divider
	}
	pClient = w.pool.client
	w.mutex.Unlock()

	if pClient == nil {
		go w.Connect()
	}

	return nil
}

/*
Connect - подключение воркера к пулу.

@return error Если произошла ошибка.
*/
func (w *Worker) Connect() error {
	var status bool

	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wUser := w.user
	wDifficulty := w.difficulty
	wHash := w.hash
	wClient := w.client
	pAddr := w.pool.addr
	pClient := w.pool.client
	pUser := w.pool.user
	pPassword := w.pool.password
	pSubscription := w.pool.subscription
	w.mutex.RUnlock()

	pUA := "miningmeter-proxy/1.0"

	// Проверка существования подключения.
	LogInfo("%s : check connect", sID, pAddr)
	if pClient != nil {
		LogInfo("%s : already connected.", sID, pAddr)
		return nil
	}
	LogInfo("%s : connecting", sID, pAddr)

	// Подключаемся к пулу.
	conn, err := net.DialTimeout("tcp", pAddr, 5*time.Second)

	if err != nil {
		LogError("%s : connection error: %s", sID, pAddr, err.Error())
		w.Disconnect()
		return fmt.Errorf("connection error: %s", err.Error())
	}

	// Инициализируем JSON-RPC клиент.
	client := rpc2.NewClientWithCodec(stratumrpc.NewStratumCodec(conn))

	client.Handle("mining.notify", mining.Notify)
	client.Handle("mining.set_difficulty", mining.SetDifficulty)

	// Привязываем воркер к соединению.
	state := rpc2.NewState()
	state.Set("worker", w)
	client.State = state

	// Стартуем клиент в отдельном потоке.
	go client.Run()

	// Привязываем соединение к пулу воркера.
	w.mutex.Lock()
	w.pool.client = client
	w.mutex.Unlock()

	// Запускаем монитор состояния соединения.
	go w.DisconnectNotify()

	var params []interface{}
	var reply []interface{}

	// Шлем subscribe пулу.
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

	LogInfo("%s > mining.subscribe: %s, %s, %d", sID, pAddr,
		response.subscriptions["mining.notify"], response.extranonce1, response.extranonce2size)

	// Шлем authorize пулу.
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
	if breply == false {
		LogError("%s : access denied to pool", sID, pAddr)
		w.Disconnect()
		return err
	}

	// Активируем метрики.
	mWorkerUp.WithLabelValues(stratumAddr, wAddr, wUser).Set(1)
	mPoolUp.WithLabelValues(stratumAddr, wHash, pAddr).Set(1)
	mDifficulty.WithLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr).Set(0)

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

	wClient.Notify("mining.set_difficulty", wDifficulty)
	LogInfo("%s < mining.set_difficulty: %f", sID, wAddr, wDifficulty)

	return nil
}

/*
SyncExtensions - Синхронизация поддерживаемых расширений с воркером.

@return bool синхронизированы ли расширения.
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

	if len(pe) > 0 {
		return true
	}

	// Грязный хак. Удаляет из запроса к пулу расширение subscribe-extranonce. Не
	// все пулы адекватно реагируют на это расширение. Например, viabtc.com при
	// наличии этого расширения в mining.configure просто не отвечает на сообшение.
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

	LogInfo("%s : sync extensions", sID, a)
	r := new(MiningConfigureRequest)
	r.extensions = he
	params, err := r.Encode()
	if err != nil {
		LogError("%s : encode extensions error: %s", sID, a, err.Error())
		return true
	}
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
		LogInfo("%s : sync not required ", sID, a)
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
UpdateData - обновление данных worker-а.

@param bool force принудительное переподключение воркера.

@return bool успешное ли обновление.
*/
func (w *Worker) UpdateData(force bool) bool {
	w.mutex.RLock()
	sID := w.id
	a := w.addr
	c := w.client
	e := w.pool.extranonce1
	e2 := w.pool.extranonce2size
	v, u := w.extensions["subscribe-extranonce"]
	w.mutex.RUnlock()

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
		workers.Add(w)
		// w.Disconnect тут не вызываем, он будет автоматически вызван при закрытии соединения.
		c.Close()

		return false
	}

	return true
}

/*
Restore - восстановление подключения к воркеру.

@param string id идентификатор сессии.

@return error если произошла ошибка
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
	workers.remove(id)

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

	wAddr := worker.addr
	wUser := worker.user
	wDifficulty := worker.difficulty
	wHash := worker.hash
	pAddr := worker.pool.addr

	worker.mutex.Unlock()
	w.mutex.Unlock()

	// Активируем метрики.
	mWorkerUp.WithLabelValues(stratumAddr, wAddr, wUser).Set(1)
	mDifficulty.WithLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr).Set(wDifficulty)

	go w.Death()

	return nil
}

/*
ResetHashrate - сброс счетчика шар.
*/
func (w *Worker) ResetHashrate() {
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.window = make(map[int64]float64)
}

/*
IncShares - увеличение счетчика шар.
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
ComputeHashrate - получение хэшрейта майнера.
*/
func (w *Worker) ComputeHashrate() float64 {
	wShares := 0.0

	stamp := time.Now().Unix() - 300 // Граница пяти минут.
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
UpdateHashrate - обновление хэшрейта.
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
		wUser := w.user
		wHash := w.hash
		wClient := w.client
		pAddr := w.pool.addr
		w.mutex.RUnlock()

		if wClient == nil {
			mSpeed.DeleteLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr)
			break
		}
		if pAddr != "" {
			mSpeed.WithLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr).Set(hashrate)
		}
		LogInfo("%s : hashrate: %.0f h/s", sID, wAddr, hashrate)

	}
	LogInfo("%s : stop hashrate calculation", sID, wAddr)
}

/*
DisconnectNotify - мониторинг состояния соединения.
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

	if p != nil {
		w.Disconnect()
	}
}

/*
Disconnect - отключение воркера от прокси.
*/
func (w *Worker) Disconnect() {
	// Получаем данные о пуле.
	w.mutex.Lock()
	sID := w.id
	wAddr := w.addr
	wHash := w.hash
	wClient := w.client
	pAddr := w.pool.addr
	pClient := w.pool.client

	w.client = nil
	w.mutex.Unlock()

	worker, _ := workers.Get(sID)

	if pClient != nil && worker == nil {
		w.mutex.Lock()
		w.pool.client = nil
		w.mutex.Unlock()

		LogInfo("%s : disconnecting", sID, pAddr)
		// Закрываем соединение с пулом.
		pClient.Close()

		// Удаляем метрики.
		ok := mPoolUp.DeleteLabelValues(stratumAddr, wHash, pAddr)
		if !ok {
			LogError("%s : error delete proxy_pool_up metric", sID, pAddr)
		}

		LogInfo("%s : disconnected from proxy", sID, pAddr)
	}

	if wClient != nil {
		LogInfo("%s : disconnecting", sID, wAddr)

		wClient.Close()

		go w.Death()

		LogInfo("%s : disconnected from proxy", sID, wAddr)
	}
}

/*
Death - подготовка к удалению воркера.
*/
func (w *Worker) Death() {
	<-time.After(1 * time.Minute)

	w.mutex.RLock()
	sID := w.id
	wAddr := w.addr
	wUser := w.user
	wHash := w.hash
	wClient := w.client
	pAddr := w.pool.addr
	w.mutex.RUnlock()

	if wClient == nil {
		LogInfo("%s : deleting metrics", sID, wAddr)

		if pAddr != "" {
			// Удаление метрик.

			mSended.DeleteLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr)
			mAccepted.DeleteLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr)
			ok := mDifficulty.DeleteLabelValues(stratumAddr, wAddr, wUser, wHash, pAddr)
			if !ok {
				LogError("%s : error delete proxy_worker_difficulty metric", sID, pAddr)
			}
			ok = mWorkerUp.DeleteLabelValues(stratumAddr, wAddr, wUser)
			if !ok {
				LogError("%s : error delete proxy_worker_up metric.", sID, wAddr)
			}
		}
		LogInfo("%s : removed from proxy", sID, wAddr)
	}
}
