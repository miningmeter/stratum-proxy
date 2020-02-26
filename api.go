/*
API - API functions.
*/
package main

import (
	"encoding/json"
	"net/http"
)

/*
Pool - the pool data for creating of the user.
*/
type Pool struct {
	Pool     string `json:"pool"`
	User     string `json:"user"`
	Password string `json:"password"`
}

/*
API - API functions.
*/
type API struct{}

/*
ServeHTTP - web handler.
*/
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	err = nil
	u := r.URL
	w.Header().Set("Content-Type", "application/json")
	if u.Path == "/api/v1/users" && (r.Method == "POST" || r.Method == "PUT") {
		var p Pool
		decoder := json.NewDecoder(r.Body)
		err = decoder.Decode(&p)
		if err == nil {
			LogInfo("proxy : API request to add user with pool %s and credentials %s:%s", "", p.Pool, p.User, p.Password)
			var user *User
			user, err = db.GetUserByPool(p.Pool, p.User)
			if err == nil {
				if user == nil {
					LogInfo("proxy : user not found and will be added", "")
					user = new(User)
					err = user.Init(p.Pool, p.User, p.Password)
				}
				if user != nil && err == nil {
					LogInfo("proxy : user successfully created with name %s", "", user.name)
					w.WriteHeader(http.StatusCreated)
					w.Write([]byte(`{"name": "` + user.name + `", "error": ""}`))
				}
			}
		}
	} else if u.Path == "/api/v1/pools" && r.Method == "GET" {
		LogInfo("proxy : API request to get pools", "")
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(`{"error": "command not found"}`))
	}
	if err != nil {
		LogError("proxy : API error: %s", "", err.Error())
		w.WriteHeader(http.StatusServiceUnavailable)
		w.Write([]byte(`{"error": "` + err.Error() + `"}`))
	}
}
