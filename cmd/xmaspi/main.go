package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/EldoranDev/xmaspi/v3/internal"
	"github.com/EldoranDev/xmaspi/v3/internal/mqtt"
	"github.com/eclipse/paho.golang/autopaho"
	"github.com/eclipse/paho.golang/paho"
	"net/url"
	"os"
	"os/signal"
	"syscall"
)

const clientID = "654838354938035789297247726751"

func main() {
	fmt.Println("starting xmaspi v3 - mqtt")

	stateUpdateChan := make(chan mqtt.State)
	defer close(stateUpdateChan)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	manager := internal.NewManager()
	defer manager.Close()

	u, err := url.Parse("mqtt://mqtt.eclipseprojects.io:1883")
	if err != nil {
		panic(err)
	}

	handler := mqtt.NewHandler(manager, stateUpdateChan, ctx)

	pahoConfig := autopaho.ClientConfig{
		ServerUrls:                    []*url.URL{u},
		KeepAlive:                     20,
		CleanStartOnInitialConnection: false,
		SessionExpiryInterval:         60,
		OnConnectionUp: func(cm *autopaho.ConnectionManager, _ *paho.Connack) {
			fmt.Println("mqtt connection started")

			if _, err := cm.Subscribe(context.Background(), &paho.Subscribe{
				Subscriptions: []paho.SubscribeOptions{
					{Topic: "xmaspi/set", QoS: 1},
				},
			}); err != nil {
				fmt.Printf("failed to subscribe (%s). No messages will be received.", err)
			}

			fmt.Println("mqtt subscriptions are made")
		},
		ClientConfig: paho.ClientConfig{
			ClientID: clientID,
			OnPublishReceived: []func(paho.PublishReceived) (bool, error){
				func(pr paho.PublishReceived) (bool, error) {
					return handler.Handle(pr.Packet.Topic, pr.Packet.Payload)
				},
			},
		},
	}

	c, err := autopaho.NewConnection(ctx, pahoConfig)
	if err != nil {
		panic(err)
	}

	if err = c.AwaitConnection(ctx); err != nil {
		panic(err)
	}

	go func() {
		for {
			select {
			case state := <-stateUpdateChan:
				payload, err := json.Marshal(state)
				if err != nil {
					fmt.Printf("could not marshal message (%s)", err)
				}

				if _, err = c.Publish(ctx, &paho.Publish{
					Topic:   mqtt.TopicSet,
					Payload: payload,
				}); err != nil {
					fmt.Printf("could not publish state message (%s)\n", err)
				}
				break
			case <-ctx.Done():
				return
			}
		}
	}()

	config, err := json.Marshal(handler.GetConfig())
	if err != nil {
		panic(err)
	}

	fmt.Println(string(config))

	c.Publish(ctx, &paho.Publish{
		Topic:   "homeassistant/devices",
		Payload: config,
	})

	// Wait for interrupt
	<-ctx.Done()

	// Await clean shutdown of mqtt client
	<-c.Done()
}
