package consumer

import (
	"context"
	"dist_calc/client"
	"dist_calc/pb"
	"dist_calc/service"
	"dist_calc/types"
	"encoding/json"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/sirupsen/logrus"
)

type DataConsumer struct {
	consumer   *kafka.Consumer
	isUp       bool
	service    service.Calculator
	httpClient *client.HTTPClient
	grpcClient *client.GRPCClient
}

func NewDataConsumer(topic string, svc service.Calculator, httpClient *client.HTTPClient, grpcClient *client.GRPCClient) (*DataConsumer, error) {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		// "bootstrap.servers": "kafka:29092",
		"bootstrap.servers": "localhost",
		"group.id":          "myGroup",
		"auto.offset.reset": "earliest",
	})

	if err != nil {
		panic(err)
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &DataConsumer{
		consumer:   c,
		service:    svc,
		httpClient: httpClient,
		grpcClient: grpcClient,
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

			httpReq := types.Distance{
				Value: distance,
				Unix:  time.Now().Unix(),
				OBUID: data.OBUID,
			}

			if err := c.httpClient.AggregateInvoice(httpReq); err != nil {
				logrus.Error("aggregate error:", err)
			}

			grpcReq := &pb.AggregateRequest{OBUID: int32(data.OBUID), Value: distance, Unix: time.Now().Unix()}

			if _, err := c.grpcClient.AggregateDistance(context.Background(), grpcReq); err != nil {
				logrus.Error("aggregate error:", err)
			}
		}
	}

}
