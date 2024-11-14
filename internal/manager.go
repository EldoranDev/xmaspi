package internal

import (
	"context"
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"github.com/EldoranDev/xmaspi/v3/internal/rendering"
	"time"
)

type Manager interface {
	Close()
	Render(ctx context.Context, renderer rendering.Renderer)
	ClearRenderer()
}

type manager struct {
	controller led.Controller
	renderer   rendering.Renderer

	cancelFunc context.CancelFunc
}

func NewManager() Manager {
	controller := led.NewController()
	controller.Init()

	return &manager{
		controller: controller,
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

	go func() {
		// Handle Init

		if init, ok := renderer.(rendering.RendererWithInitializer); ok {
			init.Init(m.controller)
		}

		for active {
			select {
			case <-time.After(renderer.FrameDuration()):
				renderer.ApplyFrame(m.controller)
				m.controller.Apply()
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

func (m *manager) Close() {
	m.controller.Close()
}
