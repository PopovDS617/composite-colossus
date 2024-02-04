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
	errCounter prometheus.Counter
}

type HTTPFunc = func(w http.ResponseWriter, r *http.Request) error

type APIError struct {
	Code int
	Err  error
}

func (e APIError) Error() string {
	return e.Err.Error()
}

func MakeHTTPHandlerFunc(fn HTTPFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {

			if apiErr, ok := err.(APIError); ok {
				writeJSON(w, apiErr.Code, apiErr)
			}

		}
	}
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

	errCounter := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: fmt.Sprintf("http_%s_%s", reqName, "error_counter"),
		Name:      "aggregator",
	})

	return &HTTPMetricHandler{
		reqCounter: reqCounter,
		reqLatency: reqLatency,
		errCounter: errCounter,
	}
}

func (mh *HTTPMetricHandler) Instrument(next HTTPFunc) HTTPFunc {

	return func(w http.ResponseWriter, r *http.Request) error {
		var err error

		defer func(start time.Time) {
			lat := float64(time.Since(start).Seconds())

			logrus.WithFields(logrus.Fields{
				"latency": lat,
				"request": r.RequestURI,
				"err":     err,
			}).Info()
			mh.reqLatency.Observe(lat)
			mh.reqCounter.Inc()
		}(time.Now())

		err = next(w, r)

		if err != nil {
			mh.errCounter.Inc()
		}
		return nil
	}
}

func HandleAggregate(svc service.Aggregator) HTTPFunc {
	return func(w http.ResponseWriter, r *http.Request) error {
		if r.Method != "POST" {

			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})

			return APIError{
				Code: http.StatusMethodNotAllowed,
				Err:  fmt.Errorf("invalid HTTP method %s", r.Method),
			}
		}

		var distance types.Distance
		if err := json.NewDecoder(r.Body).Decode(&distance); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})

			return APIError{
				Code: http.StatusBadRequest,
				Err:  fmt.Errorf("bad request"),
			}

		}
		if err := svc.AggregateDistance(distance); err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})

			return APIError{
				Code: http.StatusBadRequest,
				Err:  fmt.Errorf("bad request"),
			}

		}
		return nil
	}
}
func HandleGetInvoice(svc service.Aggregator) HTTPFunc {

	return func(w http.ResponseWriter, r *http.Request) error {

		if r.Method != "GET" {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "method not allowed"})
			return APIError{
				Code: http.StatusBadRequest,
				Err:  fmt.Errorf("invalid HTTP method %s", r.Method),
			}

		}
		values, ok := r.URL.Query()["obu"]

		if !ok || len(values) == 0 || len(values[0]) == 0 {

			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "missing OBU ID"})

			return APIError{
				Code: http.StatusBadRequest,
				Err:  fmt.Errorf("missing OBU ID"),
			}

		}

		obuID, err := strconv.Atoi(values[0])

		if err != nil {

			writeJSON(w, http.StatusBadRequest, map[string]string{"error": "wrong OBU ID format"})
			return APIError{
				Code: http.StatusBadRequest,
				Err:  fmt.Errorf("invalid OBU ID %s", values[0]),
			}
		}

		invoice, err := svc.GenerateInvoice(int(obuID))

		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
			return APIError{
				Code: http.StatusBadRequest,
				Err:  err,
			}

		}

		return writeJSON(w, http.StatusOK, invoice)

	}
}

func writeJSON(rw http.ResponseWriter, status int, v any) error {
	rw.WriteHeader(status)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}
