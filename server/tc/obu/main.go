package main

import (
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
)

const wsEndpoint = "ws://127.0.0.1:30000/ws"

var updateInterval = time.Second * 1

type OBUData struct {
	OBUID int     `json:"obu_id"`
	Lat   float64 `json:"lat"`
	Long  float64 `json:"long"`
}

func getCoord() float64 {
	n := float64(rand.Intn(100) + 1)
	f := rand.Float64()

	return n + f
}

func sendOBUData(wsConn *websocket.Conn, data OBUData) error {
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

	wsConn, _, err := websocket.DefaultDialer.Dial(wsEndpoint, nil)

	if err != nil {
		log.Fatal(err)
	}

	obuIDs := generateOBUIDs(25)

	for {

		for i := 0; i < len(obuIDs); i++ {

			lat, long := generateLatLong()

			data := OBUData{
				OBUID: obuIDs[i],
				Lat:   lat,
				Long:  long,
			}

			if err := sendOBUData(wsConn, data); err != nil {
				log.Fatal(err)
			}
		}

		time.Sleep(updateInterval)
	}
}

// func init() {
// 	r := rand.New(rand.NewSource(seed))
// }
