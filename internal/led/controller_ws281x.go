//go:build linux

package led

import (
	"fmt"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type ws281xController struct {
	device     *ws2811.WS2811
	settings   *Settings
	brightness uint8
}

func NewController(settings *Settings) Controller {
	opt := ws2811.DefaultOptions

	fmt.Printf("Starting with %d LEDs\n", len(settings.Leds))
	fmt.Println("Starting with Hardware Controller")

	opt.Channels[0].Brightness = int(settings.MaxBrightness)
	opt.Channels[0].GpioPin = int(settings.DataPin)
	opt.Channels[0].LedCount = len(settings.Leds)

	dev, _ := ws2811.MakeWS2811(&opt)

	return &ws281xController{
		device:   dev,
		settings: settings,
	}
}

func (w *ws281xController) Init() error {
	return w.device.Init()
}

func (w *ws281xController) Close() {
	w.device.Fini()
}

func (w *ws281xController) Apply() error {
	return w.device.Render()
}

func (w *ws281xController) GetLedCount() uint32 {
	return uint32(len(w.settings.Leds))
}

func (w *ws281xController) SetLed(led int, color *Color) {
	w.SetLedRaw(led, color.Int())
}

func (w *ws281xController) SetLedRaw(led int, color uint32) {
	w.device.Leds(0)[led] = color
}

func (w *ws281xController) Fill(color *Color) {
	w.FillRaw(color.Int())
}

func (w *ws281xController) FillRaw(color uint32) {
	for led := range w.device.Leds(0) {
		w.device.Leds(0)[led] = color
	}
}

func (w *ws281xController) SetBrightness(u uint8) {
	w.device.SetBrightness(0, int(u))
	w.brightness = u
}

func (w *ws281xController) GetBrightness() uint8 {
	return w.brightness
}
