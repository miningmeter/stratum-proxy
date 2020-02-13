package main

import (
	"fmt"
	"sync"
)

/*
Workers - массив подключенных воркеров.
*/
type Workers struct {
	mutex   sync.RWMutex
	workers map[string]*Worker
}

func (w *Workers) add(worker *Worker) bool {
	id := worker.GetID()

	if wr := w.get(id); wr == nil {
		w.workers[id] = worker
	} else {
		return false
	}

	return true
}

func (w *Workers) get(id string) *Worker {
	if worker, ok := w.workers[id]; ok {
		return worker
	}

	return nil
}

func (w *Workers) remove(id string) {
	if worker := w.get(id); worker != nil {
		delete(w.workers, id)
	}

	return
}

/*
Init - инициализация массива воркеров.
*/
func (w *Workers) Init() {
	w.mutex.Lock()
	w.workers = make(map[string]*Worker)
	w.mutex.Unlock()
}

/*
Add - добавление воркера в массив.

@param *Worker worker указатель на добавляемый воркер.

@return error Если произошла ошибка
*/
func (w *Workers) Add(worker *Worker) error {
	wid := worker.GetID()

	if !ValidateHexString(wid) {
		return fmt.Errorf("invalid format worker.id = %s on add to workers table", wid)
	}

	w.mutex.Lock()
	result := w.add(worker)
	w.mutex.Unlock()

	if !result {
		return fmt.Errorf("worker.id = %s already exist in workers table", wid)
	}

	return nil
}

/*
Get - получение воркера из списка по идентификатору.

@param string id идентификатор воркера.

@return *Worker указатель на найденный воркер.
        error  ошибка.
*/
func (w *Workers) Get(id string) (*Worker, error) {
	if !ValidateHexString(id) {
		return nil, fmt.Errorf("invalid format id = %s on get from workers table", id)
	}

	w.mutex.RLock()
	worker := w.get(id)
	w.mutex.RUnlock()

	if worker == nil {
		return nil, fmt.Errorf("worker with id = %s not found in workers table", id)
	}

	return worker, nil
}

/*
Remove - удаление воркера из списка.

@param string id адрес воркера.

@return error Если произошла ошибка удаления.
*/
func (w *Workers) Remove(id string) error {
	if !ValidateHexString(id) {
		return fmt.Errorf("invalid format id = %s on remove worker from workers table", id)
	}

	w.mutex.Lock()
	w.remove(id)
	w.mutex.Unlock()

	return nil
}
