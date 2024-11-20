package main

import "github.com/EldoranDev/xmaspi/v3/internal/led"

func main() {
	settings := led.Settings{
		Leds: []led.Led{
			{},
			{},
			{},
			{},
		},
		DataPin:       18,
		MaxBrightness: 255,
	}

	controller := led.NewController(&settings)

	if err := controller.Init(); err != nil {
		panic(err)
	}

	controller.Fill(&led.Color{100, 100, 100})
	if err := controller.Apply(); err != nil {
		panic(err)
	}
}
