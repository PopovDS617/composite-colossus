package handlers

import (
	"encoding/json"
	"net/http"
)

type Map map[string]any
type APIFunc func(rw http.ResponseWriter, r *http.Request) error

func writeJSON(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}

func MakeAPIFunc(fn APIFunc) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		if err := fn(rw, r); err != nil {
			writeJSON(rw, http.StatusInternalServerError, Map{"error": err.Error()})
		}
	}
}
