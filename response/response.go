package response

import (
	"encoding/json"
	"net/http"
)

type Error struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}
type Success struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

func ErrorResponse(w http.ResponseWriter, r *http.Request, statusCode int, errorMessage string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	e := &Error{false, errorMessage}
	json.NewEncoder(w).Encode(e)
}

func SuccessMessageResponse(w http.ResponseWriter, r *http.Request, statusCode int, data string) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	e := &Error{true, data}
	json.NewEncoder(w).Encode(e)
}
