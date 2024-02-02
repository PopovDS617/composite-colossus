package main

import (
	"log"
	"math"
	"math/rand"
	"obu/types"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

//const wsEndpoint = "ws://data_receiver:30000/ws"

const wsEndpoint = "ws://localhost:30000/ws"

var updateInterval = time.Second * 3

func getCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}

func sendOBUData(wsConn *websocket.Conn, data types.OBUData) error {
	return wsConn.WriteJSON(data)
}

func generateLatLong() (float64, float64) {
	return getCoord(), getCoord()
}

func generateOBUIDs(n int) []int {
	ids := make([]int, n)

	for i := 0; i < n; i++ {
		ids[i] = rand.Intn(math.MaxInt)
	}

	return ids
}

func main() {

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	wsConn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		log.Fatal(err)
	}

	obuIDs := generateOBUIDs(15)

	for {

		for i := 0; i < len(obuIDs); i++ {

			lat, long := generateLatLong()

			data := types.OBUData{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}

			logrus.Info(data)

			if err := sendOBUData(wsConn, data); err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(updateInterval)
	}
}
