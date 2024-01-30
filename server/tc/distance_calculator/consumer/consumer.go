package consumer

import (
	"dist_calc/client"
	"dist_calc/service"
	"dist_calc/types"
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type DataConsumer struct {
	consumer *kafka.Consumer
	isUp     bool
	service  service.Calculator
	client   *client.Client
}

func NewDataConsumer(topic string, svc service.Calculator, client *client.Client) (*DataConsumer, error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "kafka:29092",
		// "bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &DataConsumer{
		consumer: c,
		service:  svc,
		client:   client,
	}, nil
}

func (c *DataConsumer) Start() {
	c.isUp = true
	logrus.Info("kafka transport is up")
	c.readMessageLoop()
}

func (c *DataConsumer) readMessageLoop() {
	for c.isUp {
		msg, err := c.consumer.ReadMessage(-1)

		if err != nil {
			logrus.Errorf("kafka consume error %s", err)

		} else {
			var data types.OBUData
			if err := json.Unmarshal(msg.Value, &data); err != nil {
				logrus.Errorf("JSON serialization error: %s", err)

			}
			distance, err := c.service.CalculateDistance(data)
			if err != nil {
				logrus.Errorf("calculation error %s:", err)
			}

			result := types.Distance{
				Value: distance,
				Unix:  time.Now().Unix(),
				OBUID: data.OBUID,
			}

			if err := c.client.AggregateInvoice(result); err != nil {
				logrus.Error("aggregate error:", err)
			}
		}
	}

}
