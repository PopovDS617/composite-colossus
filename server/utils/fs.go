package utils

import (
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadAndUnmarshal[T interface{}](filepath string, dataToBind []T, resChan chan []T) {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(dir + filepath)

	if err != nil {
		log.Fatalf("cannot open %v", err)
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("cannot read %v", err)
	}

	err = json.Unmarshal(content, &dataToBind)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON: %v", err)
	}

	resChan <- dataToBind

	close(resChan)

}
