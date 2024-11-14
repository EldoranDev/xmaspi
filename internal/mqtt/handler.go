package mqtt

import (
	"context"
	"encoding/json"
	"github.com/EldoranDev/xmaspi/v3/internal"
	"github.com/EldoranDev/xmaspi/v3/internal/rendering"
)

type Handler interface {
	Handle(topic string, payload []byte) (bool, error)
	GetState() State
	GetConfig() Config
}

type handler struct {
	manager      internal.Manager
	ctx          context.Context
	stateChannel chan State
}

func NewHandler(
	manager internal.Manager,
	stateChannel chan State,
	ctx context.Context,
) Handler {
	return &handler{
		manager:      manager,
		stateChannel: stateChannel,
		ctx:          ctx,
	}
}

func (h *handler) Handle(topic string, payload []byte) (bool, error) {
	switch topic {
	case TopicSet:
		var set State
		if err := json.Unmarshal(payload, &set); err != nil {
			return false, err
		}

		return h.handleSet(set)
	}

	return true, nil
}

func (h *handler) GetState() State {
	return State{}
}

func (h *handler) GetConfig() Config {
	return Config{
		Dev: struct {
			Model     string `json:"model"`
			Name      string `json:"name"`
			SwVersion string `json:"sw_version"`
		}{
			Model:     "XMAS-PI5",
			Name:      "xmaspi",
			SwVersion: "3.0.0",
		},
		Schema:       "json",
		StateTopic:   TopicState,
		CommandTopic: TopicSet,
		Platform:     "light",
		Qos:          "2",
		// TODO: Add generation for unique id
		UniqueId:            "xmaspi-5",
		Brightness:          true,
		Effect:              true,
		EffectList:          rendering.GetAllRendererNames(),
		SupportedColorModes: []string{"rgb"},
	}
}

func (h *handler) handleSet(set State) (bool, error) {
	if set.State == "OFF" {
		h.manager.ClearRenderer()
	} else {
		renderer, err := rendering.GetRenderer(set.Effect)
		if err != nil {
			return false, err
		}

		h.manager.Render(h.ctx, renderer)
	}

	h.stateChannel <- h.GetState()

	return true, nil
}
