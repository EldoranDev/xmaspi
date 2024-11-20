package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/EldoranDev/xmaspi/v3/internal"
	"github.com/EldoranDev/xmaspi/v3/internal/led"
	"github.com/EldoranDev/xmaspi/v3/internal/mqtt"
	"github.com/EldoranDev/xmaspi/v3/internal/rendering"
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
var settingsFile string

func main() {
	fmt.Println("starting xmaspi v3 - mqtt")

	setupFlags()
	flag.Parse()

	fmt.Printf("connecting to mqtt broker at %s\n", broker)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	manager := internal.NewManager(settingsFile)
	defer manager.Close()

	u, err := url.Parse(broker)
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
		ConnectUsername: mqttUser,
		ConnectPassword: []byte(mqttPassword),
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

	renderer, _ := rendering.GetRenderer("static")

	manager.SetColor(led.Color{100, 100, 100})

	manager.Render(ctx, renderer)

	fmt.Println("Finished booting")

	// Wait for interrupt
	<-ctx.Done()

	// Await clean shutdown of mqtt client
	<-c.Done()
}

func setupFlags() {
	flag.StringVar(&broker, "broker", "mqtt://localhost:1883", "Url of the mqtt broker to connect to")
	flag.StringVar(&mqttUser, "user", "mqtt", "Username for mqtt authentication")
	flag.StringVar(&mqttPassword, "password", "mqtt", "Password for mqtt authentication")
	flag.StringVar(&settingsFile, "settings", "./settings.json", "Location of the json file that includes the led configuration")
}
