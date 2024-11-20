package internal

import (
	"context"
	"fmt"
	"github.com/EldoranDev/xmaspi/v3/internal/config"
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"github.com/EldoranDev/xmaspi/v3/internal/rendering"
	"github.com/EldoranDev/xmaspi/v3/internal/to"
	"time"
)

type Manager interface {
	Close()
	Render(ctx context.Context, renderer rendering.Renderer)
	ClearRenderer()
	Clear()
	GetRendererName() *string
	GetColor() *led.Color
	SetColor(color led.Color)
	GetBrightness() uint8
	SetBrightness(uint8)
}

type manager struct {
	controller led.Controller

	color led.Color

	renderer   rendering.Renderer
	cancelFunc context.CancelFunc
}

func NewManager(config *config.Config) Manager {
	controller := led.NewController(config)
	if err := controller.Init(); err != nil {
		panic(err)
	}

	return &manager{
		controller: controller,
		color:      led.Color{R: 255, G: 255, B: 255},
	}
}

func (m *manager) Render(ctx context.Context, renderer rendering.Renderer) {
	if m.renderer != nil {
		m.ClearRenderer()
	}

	cancelCtx, cancelFunc := context.WithCancel(context.Background())

	m.cancelFunc = cancelFunc
	m.renderer = renderer

	active := true

	fmt.Println("Setting new Renderer")

	go func() {
		if init, ok := renderer.(rendering.RendererWithInitializer); ok {
			init.Init(m.controller)
		}

		fmt.Println("Finished Initializing new Renderer")

		renderer.ApplyFrame(m.controller, &m.color)
		_ = m.controller.Apply()

		for active {
			select {
			case <-time.After(renderer.FrameDuration()):
				renderer.ApplyFrame(m.controller, &m.color)
				_ = m.controller.Apply()
				break
			case <-cancelCtx.Done():
			case <-ctx.Done():
				active = false
				break
			}
		}
	}()
}

func (m *manager) ClearRenderer() {
	if m.renderer == nil {
		// no active renderer
		return
	}

	m.cancelFunc()

	m.cancelFunc = nil
	m.renderer = nil
}

func (m *manager) Clear() {
	m.controller.FillRaw(0x000000)
	_ = m.controller.Apply()
}

func (m *manager) GetRendererName() *string {
	if m.renderer == nil {
		return nil
	}

	return to.StringPtr(m.renderer.GetIdentifier())
}

func (m *manager) GetBrightness() byte {
	return m.controller.GetBrightness()
}

func (m *manager) SetBrightness(brightness uint8) {
	m.controller.SetBrightness(brightness)
}

func (m *manager) GetColor() *led.Color {
	return &m.color
}

func (m *manager) SetColor(color led.Color) {
	m.color = color
}

func (m *manager) Close() {
	m.controller.Close()
}
