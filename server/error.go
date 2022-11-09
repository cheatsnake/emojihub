package server

import (
	"encoding/json"
	"net/http"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func HandleError(w http.ResponseWriter, code int, msg string) {
	e := ErrorMessage{msg}
	jsonErr, _ := json.Marshal(e)

	w.WriteHeader(code)
	w.Write(jsonErr)
}
