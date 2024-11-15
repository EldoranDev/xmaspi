package main

import (
	"context"
	"encoding/json"
	"flag"
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

var broker string
var mqttUser string
var mqttPassword string

func main() {
	fmt.Println("starting xmaspi v3 - mqtt")

	setupFlags()
	flag.Parse()

	fmt.Printf("connecting to mqtt broker at %s\n", broker)

	stateUpdateChan := make(chan mqtt.State)
	defer close(stateUpdateChan)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	manager := internal.NewManager()
	defer manager.Close()

	u, err := url.Parse(broker)
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

			sendBirthMessages(ctx, cm, handler)
		},
		DisconnectPacketBuilder: func() *paho.Disconnect {
			// Until I figure out how to send something before closing
			// I'll simply disconnect without a proper message
			// This will trigger the Will message to be send
			return nil
		},

		OnConnectError: func(err error) {
			panic(err)
		},
		ConnectUsername: mqttUser,
		ConnectPassword: []byte(mqttPassword),
		WillMessage: &paho.WillMessage{
			Topic:   mqtt.AvailabilityTopic,
			Payload: []byte("offline"),
			Retain:  true,
		},
		ClientConfig: paho.ClientConfig{
			ClientID: clientID,
			OnPublishReceived: []func(paho.PublishReceived) (bool, error){
				func(pr paho.PublishReceived) (bool, error) {
					fmt.Printf("Got message: %s - %s", pr.Packet.Topic, pr.Packet.Payload)

					return true, nil
				},
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
					Topic:   mqtt.StateTopic,
					Payload: payload,
					Retain:  true,
				}); err != nil {
					fmt.Printf("could not publish state message (%s)\n", err)
				}
				break
			case <-ctx.Done():
				return
			}
		}
	}()

	// Wait for interrupt
	<-ctx.Done()

	// Await clean shutdown of mqtt client
	<-c.Done()
}

func setupFlags() {
	flag.StringVar(&broker, "broker", "mqtt://localhost:1883", "Url of the mqtt broker to connect to")
	flag.StringVar(&mqttUser, "user", "mqtt", "Username for mqtt authentication")
	flag.StringVar(&mqttPassword, "password", "mqtt", "Password for mqtt authentication")
}

func sendBirthMessages(ctx context.Context, cm *autopaho.ConnectionManager, handler mqtt.Handler) {
	config, err := json.Marshal(handler.GetConfig())
	if err != nil {
		panic(err)
	}

	if _, err = cm.Publish(ctx, &paho.Publish{
		Topic:   mqtt.DiscoveryTopic,
		Payload: config,
	}); err != nil {
		panic(err)
	}

	state, err := json.Marshal(handler.GetState())
	if err != nil {
		panic(err)
	}

	if _, err = cm.Publish(ctx, &paho.Publish{
		Topic:   mqtt.StateTopic,
		Payload: state,
	}); err != nil {
		panic(err)
	}

	if _, err = cm.Publish(ctx, &paho.Publish{
		Topic:   mqtt.AvailabilityTopic,
		Payload: []byte("online"),
		Retain:  true,
	}); err != nil {
		panic(err)
	}
}
