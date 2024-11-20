//go:build linux

package led

import (
	"github.com/EldoranDev/xmaspi/v3/internal/config"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	"strconv"
)

type ws281xController struct {
	device     *ws2811.WS2811
	config     *config.Config
	leds       []Led
	brightness uint8
}

func NewController(cfg *config.Config) Controller {
	brightness, err := strconv.Atoi(cfg.Led.MaxBrightness)
	if err != nil {
		panic(err)
	}

	leds := getLeds(cfg.Led.LedFile)
	opt := ws2811.DefaultOptions

	pin, err := strconv.Atoi(cfg.DataPin)
	if err != nil {
		panic(err)
	}

	opt.Channels[0].Brightness = brightness
	opt.Channels[0].GpioPin = pin
	opt.Channels[0].LedCount = len(leds)

	dev, _ := ws2811.MakeWS2811(&opt)

	return &ws281xController{
		device: dev,
		config: cfg,
		leds:   leds,
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
	return uint32(len(w.leds))
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
