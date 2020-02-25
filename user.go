/*
Класс пользователя.
*/
package main

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"sync"
	"time"
)

/*
User - учетная запись пользователя.
*/
type User struct {
	mutex    sync.RWMutex
	name     string
	pool     string
	user     string
	password string
	hash     string
	divider  float64
	touch    int64
}

/*
GetName -  получение имени пользователя.

@return string имя пользователя
*/
func (u *User) GetName() string {
	u.mutex.RLock()
	defer u.mutex.RUnlock()

	return u.name
}

/*
Touch - обновление времени последнего использования.
*/
func (u *User) Touch() {
	u.mutex.Lock()
	u.touch = time.Now().Unix()
	u.mutex.Unlock()
}

/*
Init - инициализация пользователя.

@return error если произошла ошибка
*/
func (u *User) Init(pool string, user string, password string) error {
	if !ValidateAddr(pool, true) {
		return fmt.Errorf("invalid format pool = %s on init user", pool)
	}
	if user == "" {
		return errors.New("empty user on init user")
	}
	// Генерируем имя пользователя.
	name := ""
	for t := time.Now().Unix(); t > 0; t-- {
		h := md5.New()
		h.Write([]byte(pool + user + password + string(t)))
		id := h.Sum(nil)
		name = hex.EncodeToString(id[0:8])
		us, _ := db.GetUser(name)
		if us == nil {
			break
		}
	}

	u.mutex.Lock()
	u.name = name
	u.pool = pool
	u.user = user
	u.password = password
	u.touch = 0
	u.mutex.Unlock()

	if err := db.AddUser(u); err != nil {
		return err
	}

	LogInfo("proxy : created new user %s for account %s on pool %s", "", name, user, pool)

	return nil
}
