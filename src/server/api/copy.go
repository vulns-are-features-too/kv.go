package api

import (
	"database"
	"io"
	"net/http"
	"strings"
)

type copyHandler struct {
	db database.KvDatabase
}

func Copy(db database.KvDatabase) copyHandler {
	return copyHandler{db: db}
}

func (h copyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	bodyParts := strings.Split(string(body), " ")
	if len(bodyParts) != 2 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("request body must be in the form: srcKey dstKey"))
		return
	}

	src := bodyParts[0]
	dst := bodyParts[1]

	err := h.db.Copy(src, dst)
	if err != nil {
		writeError(w, err)
	}
}
