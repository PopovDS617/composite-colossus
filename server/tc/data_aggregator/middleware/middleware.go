package middleware

import (
	"data_aggregator/service"
	"data_aggregator/types"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/sirupsen/logrus"
)

type MetricsMiddleware struct {
	errCounterAggregate prometheus.Counter
	errCounterCalculate prometheus.Counter
	reqCounterAggregate prometheus.Counter
	reqCounterCalculate prometheus.Counter
	reqLatencyAggregate prometheus.Histogram
	reqLatencyCalculate prometheus.Histogram
	next                service.Aggregator
}

func NewMetricsMiddleware(next service.Aggregator) *MetricsMiddleware {

	errCounterAggregate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "aggregate_error_counter",
		Name:      "aggregate"})

	errCounterCalculate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "calculate_error_counter",
		Name:      "calculate"})

	reqCounterAggregate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "aggregate_request_counter",
		Name:      "aggregate"})
	reqCounterCalculate := promauto.NewCounter(prometheus.CounterOpts{
		Namespace: "calculate_request_counter",
		Name:      "calculate"})

	reqLatencyAggregate := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "aggregate_request_latency",
		Name:      "aggregate", Buckets: []float64{0.2, 0.4, 0.6, 0.8, 1.0}})
	reqLatencyCalculate := promauto.NewHistogram(prometheus.HistogramOpts{
		Namespace: "calculate_request_latency",
		Name:      "calculate", Buckets: []float64{0.2, 0.4, 0.6, 0.8, 1.0}})

	return &MetricsMiddleware{
		next:                next,
		reqCounterAggregate: reqCounterAggregate,
		reqCounterCalculate: reqCounterCalculate,
		reqLatencyAggregate: reqLatencyAggregate,
		reqLatencyCalculate: reqLatencyCalculate,
		errCounterAggregate: errCounterAggregate,
		errCounterCalculate: errCounterCalculate,
	}
}

type LogMiddleware struct {
	next service.Aggregator
}

func NewLogMiddleware(next service.Aggregator) service.Aggregator {
	return &LogMiddleware{
		next: next,
	}
}

func (lg *LogMiddleware) AggregateDistance(distance types.Distance) (err error) {
	defer func(start time.Time) {

		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("aggregating distance")
	}(time.Now())

	err = lg.next.AggregateDistance(distance)

	return
}
func (lg *LogMiddleware) GenerateInvoice(id int) (invoice *types.Invoice, err error) {
	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
		}).Info("generating invoice")
	}(time.Now())

	invoice, err = lg.next.GenerateInvoice(id)

	return
}

func (mm *MetricsMiddleware) AggregateDistance(distance types.Distance) (err error) {

	defer func(start time.Time) {
		mm.reqLatencyAggregate.Observe(float64(time.Since(start).Seconds()))
		mm.reqCounterAggregate.Inc()
		if err != nil {
			mm.errCounterAggregate.Inc()
		}
	}(time.Now())

	err = mm.next.AggregateDistance(distance)

	return
}
func (mm *MetricsMiddleware) GenerateInvoice(id int) (invoice *types.Invoice, err error) {

	defer func(start time.Time) {
		mm.reqLatencyCalculate.Observe(float64(time.Since(start).Seconds()))
		mm.reqCounterCalculate.Inc()
		if err != nil {
			mm.errCounterCalculate.Inc()
		}
	}(time.Now())

	invoice, err = mm.next.GenerateInvoice(id)

	return
}
