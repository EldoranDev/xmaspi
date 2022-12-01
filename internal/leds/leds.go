package leds

import (
	"encoding/json"
	"os"
)

type Led struct {
	Index int
	Pos   Coordinate
}

type Leds = []Led

type Coordinate struct {
	X, Y, Z int
}

func LoadLeds(file string) (Leds, error) {
	data, err := os.ReadFile(file)

	if err != nil {
		return nil, err
	}

	var leds Leds

	err = json.Unmarshal(data, &leds)

	return leds, err
}

func SaveLeds(file string, leds Leds) error {
	data, err := json.Marshal(leds)

	if err != nil {
		return err
	}

	err = os.WriteFile(file, data, 0644)

	return err
}
