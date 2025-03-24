package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/anthdm/hollywood/actor"
	serversession "github.com/saiddis/taxi_service/actor/server_session"
)

func main() {

	var mongodbURI string
	if mongodbURI = os.Getenv("MONGODB_URI"); mongodbURI == "" {
		log.Fatal("error getting db connection string")
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)

	go func() { <-c; cancel() }()

	log.Println("started")

	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	producer, err := serversession.New(mongodbURI)
	if err != nil {
		log.Fatalf("error creating server_session actor: %v", err)
	}

	pid := e.Spawn(producer, "server_session")

	<-ctx.Done()
	<-e.Poison(pid).Done()

	log.Println("server_session: done")
}
