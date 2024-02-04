package main

import (
	"math"
	"math/rand"
	"obu/types"
	"os"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

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
	var wsEndpoint = os.Getenv("WS_RECEIVER_ADDRESS")

	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors: true,
	})

	wsConn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		logrus.Fatal(err)
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
				logrus.Fatal(err)
			}
		}

		time.Sleep(updateInterval)
	}
}
