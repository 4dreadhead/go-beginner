package response

import (
	"encoding/json"
	"net/http"
)

type ErrorBody struct {
	Message string `json:"message"`
}

func Json(w http.ResponseWriter, data any, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	obj := ErrorBody{
		Message: err.Error(),
	}
	json.NewEncoder(w).Encode(obj)
}
