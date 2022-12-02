package controller

import (
	"context"
	"github.com/EldoranDev/xmaspi/v2/internal/leds"
	ws2811 "github.com/rpi-ws281x/rpi-ws281x-go"
	"time"
)

type Controller interface {
	Init() error
	Close()

	Render() error

	StartAnimation(animaton Animation)
	StopAnimation()

	SetLed(led int, color uint32)
	Fill(color uint32)

	LedCount() int
	Leds() leds.Leds
}

func NewController(leds leds.Leds, options ...Option) (Controller, error) {
	o, err := resolveOptions(options)

	if err != nil {
		return nil, err
	}

	opt := ws2811.DefaultOptions

	opt.Channels[0].Brightness = o.brightness
	opt.Channels[0].LedCount = len(leds)
	opt.Channels[0].GpioPin = o.dataPin

	dev, _ := ws2811.MakeWS2811(&opt)

	return &controller{
		leds:     leds,
		device:   dev,
		ledCount: len(leds),
	}, nil
}

type controller struct {
	device           *ws2811.WS2811
	ledCount         int
	animationRunning bool

	leds leds.Leds

	cancelFunc context.CancelFunc
}

func (c *controller) Init() error {
	return c.device.Init()
}

func (c *controller) Leds() leds.Leds {
	return c.leds
}

func (c *controller) Close() {
	c.device.Fini()
}

func (c *controller) Fill(color uint32) {
	for i := 0; i < c.ledCount; i++ {
		c.SetLed(i, color)
	}
}

func (c *controller) SetLed(led int, color uint32) {
	c.device.Leds(0)[led] = color
}

func (c *controller) Render() error {
	return c.device.Render()
}

func (c *controller) LedCount() int {
	return c.ledCount
}

func (c *controller) StartAnimation(anim Animation) {
	// Check if animation is running
	ctx, cancelFunc := context.WithCancel(context.Background())

	c.cancelFunc = cancelFunc
	c.animationRunning = true

	go func() {
		active := true

		// Initialize Animation if initializer implemented
		if init, ok := anim.(AnimationWithInitializer); ok {
			init.Init(c)
		}

		for active {
			select {
			case <-time.After(anim.FrameDuration()):
				anim.ApplyFrame(c)
				_ = c.Render()
				break
			case <-ctx.Done():
				active = false
			}
		}
	}()
}

func (c *controller) StopAnimation() {
	if !c.animationRunning {
		return
	}

	c.cancelFunc()
	c.cancelFunc = nil
}
