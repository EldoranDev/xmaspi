package led

import (
	"encoding/json"
	"os"
)

func getLeds(file string) []Led {
	data, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}

	var leds []Led
	if err = json.Unmarshal(data, &leds); err != nil {
		panic(err)
	}

	return leds
}
