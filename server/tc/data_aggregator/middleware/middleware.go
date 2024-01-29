package middleware

import (
	"data_aggregator/service"
	"data_aggregator/types"
	"time"

	"github.com/sirupsen/logrus"
)

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
		}).Info()
	}(time.Now())

	err = lg.next.AggregateDistance(distance)

	return
}
