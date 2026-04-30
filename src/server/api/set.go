package api

import (
	"database"
	"io"
	"net/http"
	"strings"
)

type setHandler struct {
	db database.KvDatabase
}

func Set(db database.KvDatabase) setHandler {
	return setHandler{db: db}
}

func (h setHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	bodyParts := strings.Split(string(body), "=")
	if len(bodyParts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request body must be in the form: key=value"))
		return
	}

	key := bodyParts[0]
	val := bodyParts[1]

	err := h.db.Set(key, val)
	if err != nil {
		writeError(w, err)
	}
}
