package transport

import (
	"bytes"
	"context"
	"data_aggregator/service"
	aggendpoint "data_aggregator/service/aggregate-endpoint"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/go-kit/kit/circuitbreaker"
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/ratelimit"
	"github.com/go-kit/kit/transport"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/go-kit/log"
	"github.com/sony/gobreaker"
	"golang.org/x/time/rate"
)

func errorEncoder(ctx context.Context, err error, w http.ResponseWriter) {
	fmt.Println("--- error encoder ---", err)

}

func NewHTTPHandler(endpoints aggendpoint.Set, logger log.Logger) http.Handler {
	options := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(errorEncoder),
		httptransport.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
	}

	m := http.NewServeMux()
	m.Handle("/aggregator", httptransport.NewServer(
		endpoints.AggregateEndpoint,
		decodeHTTPAggregateRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	m.Handle("/invoice", httptransport.NewServer(
		endpoints.CalculateEndpoint,
		decodeHTTPCalculateRequest,
		encodeHTTPGenericResponse,
		options...,
	))
	return m

}

func decodeHTTPAggregateRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req aggendpoint.AggregateRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	return req, err

}

func decodeHTTPCalculateRequest(_ context.Context, r *http.Request) (interface{}, error) {

	var req aggendpoint.CalculateRequest

	values, ok := r.URL.Query()["obu"]

	if !ok {
		return nil, errors.New("no obu id")
	}

	obuID, err := strconv.Atoi(values[0])

	if err != nil {
		return nil, errors.New("cannot parse obu id")
	}

	req.OBUID = obuID

	return req, nil

}

func encodeHTTPGenericResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	// if f, ok := response.(endpoint.Failer); ok && f.Failed() != nil {
	// 	errorEncoder(ctx, f.Failed(), w)
	// 	return nil
	// }
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeHTTPGenericRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = io.NopCloser(&buf)
	return nil
}

func NewHTTPClient(instance string, logger log.Logger) (service.Service, error) {

	if !strings.HasPrefix(instance, "http") {
		instance = "http://" + instance
	}

	u, err := url.Parse(instance)

	if err != nil {
		return nil, err
	}

	limiter := ratelimit.NewErroringLimiter(rate.NewLimiter(rate.Every(time.Second), 100))

	// var options []httptransport.ClientOption

	var aggregateEndpoint endpoint.Endpoint
	var calculateEndpoint endpoint.Endpoint

	aggregateEndpoint = httptransport.NewClient(
		"POST",
		copyURL(u, "/aggregator"),
		encodeHTTPGenericRequest,
		decodeHTTPAggregateResponse,
	).Endpoint()
	aggregateEndpoint = limiter(aggregateEndpoint)
	aggregateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "Aggregate",
		Timeout: 30 * time.Second,
	}))(aggregateEndpoint)

	calculateEndpoint = httptransport.NewClient(
		"GET",
		copyURL(u, "/invoice"),
		encodeHTTPGenericRequest,
		decodeHTTPCalculateResponse,
	).Endpoint()
	calculateEndpoint = limiter(calculateEndpoint)
	calculateEndpoint = circuitbreaker.Gobreaker(gobreaker.NewCircuitBreaker(gobreaker.Settings{
		Name:    "Calculate",
		Timeout: 30 * time.Second,
	}))(calculateEndpoint)

	return aggendpoint.Set{
		AggregateEndpoint: aggregateEndpoint,
		CalculateEndpoint: calculateEndpoint,
	}, nil

}

func copyURL(base *url.URL, path string) *url.URL {
	next := *base
	next.Path = path
	return &next
}

func decodeHTTPAggregateResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp aggendpoint.AggregateResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
func decodeHTTPCalculateResponse(_ context.Context, r *http.Response) (interface{}, error) {
	if r.StatusCode != http.StatusOK {
		return nil, errors.New(r.Status)
	}
	var resp aggendpoint.CalculateResponse
	err := json.NewDecoder(r.Body).Decode(&resp)
	return resp, err
}
