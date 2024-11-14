package main

import (
	"context"
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
const baseTopic = "xmaspi"

func main() {
	fmt.Println("starting xmaspi v3 - mqtt")

	commandChannel := make(chan int)
	defer close(commandChannel)

	statusChannel := make(chan int)
	defer close(statusChannel)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	manager := internal.NewManager()
	defer manager.Close()

	u, err := url.Parse("mqtt://mqtt.eclipseprojects.io:1883")
	if err != nil {
		panic(err)
	}

	handler := mqtt.NewHandler(manager, ctx)

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

	// Wait for interrupt
	<-ctx.Done()

	// Await clean shutdown of mqtt client
	<-c.Done()
}
