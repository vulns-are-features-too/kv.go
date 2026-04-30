package server

import (
	"api"
	"database"
	"fmt"
	"net/http"
)

const (
	DefaultHost = "127.0.0.1"
	DefaultPort = 43219
)

func MakeServer(db database.KvDatabase) *http.Server {
	registerHandlers(db)

	return &http.Server{
		Addr: fmt.Sprintf("%s:%d", DefaultHost, DefaultPort),
	}
}

func registerHandlers(db database.KvDatabase) {
	http.Handle("/get", api.Get(db))
	http.Handle("/getkeys", api.GetKeys(db))
	http.Handle("/set", api.Set(db))
}
