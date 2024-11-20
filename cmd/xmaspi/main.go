package main

import (
	"context"
	"fmt"
	"github.com/EldoranDev/xmaspi/v3/internal"
	"github.com/EldoranDev/xmaspi/v3/internal/cli"
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

	config := cli.GetConfig()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	manager := internal.NewManager(config)
	defer manager.Close()

	u, err := url.Parse(config.Mqtt.Broker)
	if err != nil {
		panic(err)
	}

	handler := mqtt.NewHandler(manager, ctx)

	pahoConfig := autopaho.ClientConfig{
		ServerUrls:                    []*url.URL{u},
		KeepAlive:                     20,
		CleanStartOnInitialConnection: false,
		SessionExpiryInterval:         60,
		OnConnectionUp:                handler.OnConnectionUp,
		DisconnectPacketBuilder: func() *paho.Disconnect {
			// Until I figure out how to send something before closing
			// I'll simply disconnect without a proper message
			// This will trigger the Will message to be send
			return nil
		},
		OnConnectError: func(err error) {
			panic(err)
		},
		ConnectUsername: config.Mqtt.User,
		ConnectPassword: []byte(config.Mqtt.Password),
		WillMessage:     handler.GetWillMessage(),
		ClientConfig: paho.ClientConfig{
			ClientID: clientID,
			OnPublishReceived: []func(paho.PublishReceived) (bool, error){
				handler.Handle,
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

	go handler.SendStateUpdates(c)

	// Wait for interrupt
	<-ctx.Done()

	// Await clean shutdown of mqtt client
	<-c.Done()
}
