package handlers

import (
	"data_aggregator/service"
	"data_aggregator/types"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
)

type HTTPMetricHandler struct {
	reqCounter prometheus.Counter
	reqLatency prometheus.Histogram
}

func NewHTTPMetricHandler(reqName string) *HTTPMetricHandler {

	reqCounter := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: fmt.Sprintf("http_%s_%s", reqName, "request_counter"),
		Name:      "aggregator",
	})

	reqLatency := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: fmt.Sprintf("http_%s_%s", reqName, "request_latency"),
		Name:      "aggregator",
		Buckets:   []float64{0.1, 0.5, 1.0},
	})

	return &HTTPMetricHandler{
		reqCounter: reqCounter,
		reqLatency: reqLatency,
	}
}

func (mh *HTTPMetricHandler) Instrument(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		defer func(start time.Time) {
			lat := float64(time.Since(start).Seconds())

			logrus.WithFields(logrus.Fields{
				"latency": lat,
				"request": r.RequestURI,
			}).Info()
			mh.reqLatency.Observe(lat)
		}(time.Now())

		mh.reqCounter.Inc()
		next(w, r)
	}
}

func HandleAggregate(svc service.Aggregator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
		}

		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
	}
}
func HandleGetInvoice(svc service.Aggregator) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method != "GET" {
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		values, ok := r.URL.Query()["obu"]

		if !ok || len(values) == 0 || len(values[0]) == 0 {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing OBU ID"})
			return
		}

		obuID, err := strconv.Atoi(values[0])

		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong OBU ID format"})
			return
		}

		invoice, err := svc.GenerateInvoice(int(obuID))

		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, invoice)

	}
}

func writeJSON(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}
