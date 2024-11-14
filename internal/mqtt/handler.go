package mqtt

import (
	"context"
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v3/internal"
	"github.com/EldoranDev/xmaspi/v3/internal/rendering"
)

type Handler interface {
	Handle(topic string, payload []byte) (bool, error)
}

type handler struct {
	manager internal.Manager
	ctx     context.Context
}

func NewHandler(manager internal.Manager, ctx context.Context) Handler {
	return &handler{
		manager: manager,
		ctx:     ctx,
	}
}

func (h *handler) Handle(topic string, payload []byte) (bool, error) {

	switch topic {
	case TopicSet:
		var set Set
		if err := json.Unmarshal(payload, &set); err != nil {
			return false, err
		}

		return h.handleSet(set)
	}

	return true, nil
}

func (h *handler) handleSet(set Set) (bool, error) {
	if set.State == "OFF" {
		h.manager.ClearRenderer()
	} else {
		renderer, err := rendering.GetRenderer(set.Effect)
		if err != nil {
			return false, err
		}

		h.manager.Render(h.ctx, renderer)
	}

	return true, nil
}
