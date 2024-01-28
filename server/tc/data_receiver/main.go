package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	"github.com/gorilla/websocket"
)

var kafkaTopic = "obudata"

type OBUData struct {
	OBUID int     `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

type DataReceiver struct {
	msgCh    chan OBUData
	wsConn   *websocket.Conn
	producer *kafka.Producer
}

func (receiver *DataReceiver) handleWS(w http.ResponseWriter, r *http.Request) {

	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	wsConn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Fatal(err)
	}

	receiver.wsConn = wsConn

	go receiver.handleReceiveWS()

}

func (receiver *DataReceiver) handleReceiveWS() {
	fmt.Println("--- new obu connected")
	for {
		var data OBUData
		if err := receiver.wsConn.ReadJSON(&data); err != nil {
			log.Println("read error", err)
			continue
		}
		// fmt.Printf("received OBU data from [%d] :: lat %.2f | long %.2f\n", data.OBUID, data.Lat, data.Long)
		// receiver.msgCh <- data
		err := receiver.produceData(data)

		if err != nil {
			fmt.Println("kafka produce error:", err)
		}

	}

}

func newDataReceiver() (*DataReceiver, error) {
	p, err := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if err != nil {
		panic(err)
	}

	if err != nil {
		return nil, err
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafka.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Delivery failed: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Delivered message to %v\n", ev.TopicPartition)
				}
			}
		}
	}()

	return &DataReceiver{
		msgCh:    make(chan OBUData, 128),
		producer: p,
	}, nil
}

func (receiver *DataReceiver) produceData(data OBUData) error {

	b, err := json.Marshal(data)

	if err != nil {
		return err
	}

	err = receiver.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &kafkaTopic,
			Partition: kafka.PartitionAny},
		Value: b,
	}, nil)

	return err
}

func main() {

	receiver, err := newDataReceiver()

	if err != nil {
		log.Fatal(err)
	}

	defer receiver.producer.Close()

	http.HandleFunc("/ws", receiver.handleWS)
	http.ListenAndServe(":30000", nil)

	// receiver.producer.Flush(15 * 1000)

}
