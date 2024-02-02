package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
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

		defer func(start time.Time) {

			logrus.WithFields(logrus.Fields{
				"took": time.Since(start),
				"uri":  r.RequestURI,
			}).Info("req :: ")

		}(time.Now())

		if err := fn(rw, r); err != nil {
			writeJSON(rw, http.StatusInternalServerError, Map{"error": err.Error()})
		}
	}
}
