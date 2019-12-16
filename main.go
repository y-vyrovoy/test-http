package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	http_server "main/http-server"
	"os"
	"os/signal"
	"sync"
)

func main() {

	ctx, cancel := context.WithCancel(context.Background())
	wg := &sync.WaitGroup{}

	setupGracefulShutdown(cancel)

	httpServ := http_server.New(`data`, "8080")

	httpServ.Run(ctx, wg)

	wg.Wait()
}

func setupGracefulShutdown(stop func()) {

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	go func() {
		<-ch

		log.Info("Got interrupt signal")
		stop()
	}()
}
