package mqtt

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EldoranDev/xmaspi/v3/internal"
	"github.com/EldoranDev/xmaspi/v3/internal/rendering"
	"github.com/EldoranDev/xmaspi/v3/internal/to"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
)

type Handler interface {
	OnConnectionUp(cm *autopaho.ConnectionManager, _ *paho.Connack)
	Handle(pr paho.PublishReceived) (bool, error)
	GetWillMessage() *paho.WillMessage
	SendStateUpdates(cm *autopaho.ConnectionManager)
	GetState() State
	GetConfig() Config
}

type handler struct {
	manager      internal.Manager
	ctx          context.Context
	stateUpdates chan State
}

func NewHandler(
	manager internal.Manager,
	ctx context.Context,
) Handler {
	return &handler{
		manager:      manager,
		ctx:          ctx,
		stateUpdates: make(chan State),
	}
}

func (h *handler) Handle(pr paho.PublishReceived) (bool, error) {
	switch pr.Packet.Topic {
	case CommandTopic:
		var set State
		if err := json.Unmarshal(pr.Packet.Payload, &set); err != nil {
			return false, err
		}

		return h.handleSet(set)
	}

	return true, nil
}

func (h *handler) GetState() State {
	var state string

	effect := h.manager.GetRendererName()

	if effect != nil {
		state = "ON"
	} else {
		state = "OFF"
	}

	return State{
		Brightness: to.Uint8Ptr(h.manager.GetBrightness()),
		Color:      h.manager.GetColor(),
		Effect:     effect,
		State:      state,
	}
}

func (h *handler) GetWillMessage() *paho.WillMessage {
	return &paho.WillMessage{
		Topic:   AvailabilityTopic,
		Payload: []byte("offline"),
		Retain:  true,
	}
}

func (h *handler) OnConnectionUp(cm *autopaho.ConnectionManager, _ *paho.Connack) {
	fmt.Println("mqtt connection started")

	if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
		Subscriptions: []paho.SubscribeOptions{
			{Topic: "xmaspi/set", QoS: 1},
		},
	}); err != nil {
		fmt.Printf("failed to subscribe (%s). No messages will be received.", err)
	}

	h.sendBirthMessages(h.ctx, cm)
}

func (h *handler) GetConfig() Config {
	return Config{
		Device: Device{
			Model:        "XMAS-PI5",
			Name:         "xmaspi",
			Manufacturer: "EldoranDev",
			SwVersion:    "3.0.0",

			Identifiers: []string{
				"xmaspi",
			},
		},
		Origin: Origin{
			Name:            "xmaspi",
			SoftwareVersion: "3.0.0",
			Url:             "https://github.com/EldoranDev/xmaspi",
		},
		Components: map[string]Component{
			"leds": {
				Schema:       "json",
				StateTopic:   StateTopic,
				CommandTopic: CommandTopic,
				Platform:     "light",
				Qos:          "2",
				Availability: Availability{
					Topic: AvailabilityTopic,
				},
				// TODO: Add generation for unique id
				UniqueId:            "xmaspi-5",
				Brightness:          true,
				Effect:              true,
				EffectList:          rendering.GetAllRendererNames(),
				SupportedColorModes: []string{"rgb"},
			},
		},
	}
}

func (h *handler) sendBirthMessages(ctx context.Context, cm *autopaho.ConnectionManager) {
	config, err := json.Marshal(h.GetConfig())
	if err != nil {
		panic(err)
	}

	if _, err = cm.Publish(ctx, &paho.Publish{
		Topic:   DiscoveryTopic,
		Payload: config,
	}); err != nil {
		panic(err)
	}

	state, err := json.Marshal(h.GetState())
	if err != nil {
		panic(err)
	}

	if _, err = cm.Publish(ctx, &paho.Publish{
		Topic:   StateTopic,
		Payload: state,
	}); err != nil {
		panic(err)
	}

	if _, err = cm.Publish(ctx, &paho.Publish{
		Topic:   AvailabilityTopic,
		Payload: []byte("online"),
		Retain:  true,
	}); err != nil {
		panic(err)
	}
}

func (h *handler) SendStateUpdates(cm *autopaho.ConnectionManager) {
	for {
		select {
		case state := <-h.stateUpdates:
			payload, err := json.Marshal(state)
			if err != nil {
				fmt.Printf("could not marshal message (%s)", err)
			}

			if _, err = cm.Publish(h.ctx, &paho.Publish{
				Topic:   StateTopic,
				Payload: payload,
				Retain:  true,
			}); err != nil {
				fmt.Printf("could not publish state message (%s)\n", err)
			}
			break
		case <-h.ctx.Done():
			return
		}
	}
}
