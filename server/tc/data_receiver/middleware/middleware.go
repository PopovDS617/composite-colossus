package middleware

import (
	"receiver/producer"
	"receiver/types"
	"time"

	"github.com/sirupsen/logrus"
)

type LogMiddleware struct {
	next producer.DataProducer
}

func NewLogMiddleware(next producer.DataProducer) *LogMiddleware {
	return &LogMiddleware{
		next,
	}
}

func (l *LogMiddleware) ProduceData(data types.OBUData) error {

	start := time.Now()

	defer func(start time.Time) {
		logrus.WithFields(logrus.Fields{"obuID": data.OBUID,
			"lat":  data.Lat,
			"long": data.Long,
			"took": time.Since(start)}).Info("producing to kafka")
	}(start)

	return l.next.ProduceData(data)

}
