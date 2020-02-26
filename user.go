/*
User.
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
User - User.
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
GetName - getting of user name.

@return string user name
*/
func (u *User) GetName() string {
	u.mutex.RLock()
	defer u.mutex.RUnlock()

	return u.name
}

/*
Touch - updating of last time used.
*/
func (u *User) Touch() {
	u.mutex.Lock()
	u.touch = time.Now().Unix()
	u.mutex.Unlock()
}

/*
Init - initializing of user.

@return error
*/
func (u *User) Init(pool string, user string, password string) error {
	if !ValidateAddr(pool, true) {
		return fmt.Errorf("invalid format pool = %s on init user", pool)
	}
	if user == "" {
		return errors.New("empty user on init user")
	}
	// Generating of user name.
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
