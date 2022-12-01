package led

import (
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
)

type Controller interface {
	Init() error
	Close()

	Render() error

	SetLed(led int, color uint32, render bool) error
	Fill(color uint32, render bool) error

	LedCount() int
}

func NewController(options ...Option) (Controller, error) {
	o, err := resolveOptions(options)

	if err != nil {
		return nil, err
	}

	opt := ws2811.DefaultOptions

	opt.Channels[0].Brightness = o.brightness
	opt.Channels[0].LedCount = o.ledCount
	opt.Channels[0].GpioPin = o.dataPin

	dev, _ := ws2811.MakeWS2811(&opt)

	return &controller{
		device:   dev,
		ledCount: o.ledCount,
	}, nil
}

type controller struct {
	device   *ws2811.WS2811
	ledCount int
}

func (c *controller) Init() error {
	return c.device.Init()
}

func (c *controller) Close() {
	c.device.Fini()
}

func (c *controller) Fill(color uint32, render bool) error {
	for i := 0; i < c.ledCount; i++ {
		_ = c.SetLed(i, color, false)
	}

	if render {
		return c.Render()
	}

	return nil
}

func (c *controller) SetLed(led int, color uint32, render bool) error {
	c.device.Leds(0)[led] = color

	if render {
		return c.Render()
	}

	return nil
}

func (c *controller) Render() error {
	return c.device.Render()
}

func (c *controller) LedCount() int {
	return c.ledCount
}
