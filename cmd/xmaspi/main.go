package main

import (
	"context"
	"fmt"
	"github.com/eclipse/paho.golang/autopaho"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("starting xmaspi v3 - mqtt")

	commandChannel := make(chan int)
	defer close(commandChannel)

	statusChannel := make(chan int)
	defer close(statusChannel)

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pahoConfig := autopaho.ClientConfig{}

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
