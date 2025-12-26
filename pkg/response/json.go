package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func MethodNotAllowed(w http.ResponseWriter) {
	JSON(w, http.StatusMethodNotAllowed, map[string]string{
		"error": "method not allowed",
	})
}
