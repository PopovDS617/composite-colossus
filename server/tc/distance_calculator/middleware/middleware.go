package middleware

import (
	"dist_calc/service"
	"dist_calc/types"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next service.Calculator
}

func NewLogMiddleware(next service.Calculator) service.Calculator {

	return &LogMiddleware{
		next: next,
	}
}

func (l *LogMiddleware) CalculateDistance(data types.OBUData) (distance float64, err error) {

	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{
			"took": time.Since(start),
			"err":  err,
			"dist": fmt.Sprintf("%.2f", distance),
		}).Info("calculate distance")
	}(time.Now())

	distance, err = l.next.CalculateDistance(data)

	return
}
