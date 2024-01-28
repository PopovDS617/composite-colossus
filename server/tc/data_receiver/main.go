package main

import (
	"fmt"
	"log"
	"net/http"
	"receiver/middleware"
	"receiver/producer"
	"receiver/types"

	"github.com/gorilla/websocket"
)

type DataReceiver struct {
	msgCh    chan types.OBUData
	wsConn   *websocket.Conn
	producer producer.DataProducer
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
		var data types.OBUData
		if err := receiver.wsConn.ReadJSON(&data); err != nil {
			log.Println("read error", err)
			continue
		}
		// fmt.Printf("received OBU data from [%d] :: lat %.2f | long %.2f\n", data.OBUID, data.Lat, data.Long)
		// receiver.msgCh <- data
		err := receiver.produceData(data)

		if err != nil {
			fmt.Println("produce error:", err)
		}

	}

}

func newDataReceiver(producer producer.DataProducer) (*DataReceiver, error) {

	producer = middleware.NewLogMiddleware(producer)

	return &DataReceiver{
		msgCh:    make(chan types.OBUData, 128),
		producer: producer,
	}, nil
}

func (receiver *DataReceiver) produceData(data types.OBUData) error {
	return receiver.producer.ProduceData(data)
}

func main() {

	var kafkaTopic = "obudata"

	producer, err := producer.NewKafkaProducer(kafkaTopic)

	if err != nil {
		log.Fatal(err)
	}

	receiver, err := newDataReceiver(producer)

	if err != nil {
		log.Fatal(err)
	}
	// defer receiver.producer.Close()

	http.HandleFunc("/ws", receiver.handleWS)
	http.ListenAndServe(":30000", nil)

	// receiver.producer.Flush(15 * 1000)

}
