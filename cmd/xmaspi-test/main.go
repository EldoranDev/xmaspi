package main

import (
	"github.com/EldoranDev/xmaspi/v3/internal/config"
	"github.com/EldoranDev/xmaspi/v3/internal/led"
)

func main() {
	cfg := &config.Config{
		DataPin: "18",
		Led: config.LedConfig{
			MaxBrightness: "100",
			LedFile:       "./leds.json",
		},
	}

	controller := led.NewController(cfg)

	if err := controller.Init(); err != nil {
		panic(err)
	}

	controller.Fill(&led.Color{100, 100, 100})
	if err := controller.Apply(); err != nil {
		panic(err)
	}
}
