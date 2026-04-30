package api

import (
	"database"
	"net/http"
)

type getKeysHandler struct {
	db database.KvDatabase
}

func GetKeys(db database.KvDatabase) getKeysHandler {
	return getKeysHandler{db: db}
}

func (h getKeysHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	keys, err := h.db.GetKeys()
	if err != nil {
		writeError(w, err)
	}

	res, err := createResponse(keys)
	if err != nil {
		writeError(w, err)
	}

	w.Write([]byte(res))
}
