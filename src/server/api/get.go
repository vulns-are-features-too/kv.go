package api

import (
	"database"
	"io"
	"net/http"
)

type getHandler struct {
	db database.KvDatabase
}

func Get(db database.KvDatabase) getHandler {
	return getHandler{db: db}
}

func (h getHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	key := string(body)

	val, err := h.db.Get(key)
	if err != nil {
		writeError(w, err)
	}

	res, err := createResponse(val)
	if err != nil {
		writeError(w, err)
	}

	w.Write([]byte(res))
}
