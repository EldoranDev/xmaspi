package client

import (
	"encoding/json"
	"image"
	"log"
	"os"
)

type ScanResult = [][]image.Point

func LoadScan(file string, ledCount int) (ScanResult, error) {
	scan := make([][]image.Point, 4)

	for s := range scan {
		scan[s] = make([]image.Point, ledCount)
	}

	content, err := os.ReadFile(file)

	if err != nil {
		log.Println("Could not load scan, raw scan created")
		return nil, err
	}

	err = json.Unmarshal(content, &scan)

	if err != nil {
		log.Fatalf("Could not load given scan")
	}

	return scan, err
}
