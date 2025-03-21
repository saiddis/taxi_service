package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/anthdm/hollywood/actor"
	riderengine "github.com/saiddis/taxi_service/actor/riderEngine"
)

func main() {

	var mongodbURIFile string
	if mongodbURIFile = os.Getenv("MONGODB_URI_FILE"); mongodbURIFile == "" {
		log.Fatal("error getting db connection string")
	}

	mongodbURI, err := os.ReadFile(mongodbURIFile)
	if err != nil {
		log.Fatal("error reading connection string")
	}

	ctx, cancel := context.WithCancel(context.Background())
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() { <-c; cancel() }()

	log.Println("started")

	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		log.Fatal(err)
	}

	riderEngineProducer, err := riderengine.New(string(mongodbURI))
	if err != nil {
		log.Fatalf("error creating riderengine actor: %v", err)
	}

	pid := e.Spawn(riderEngineProducer, "riderengine")

	<-ctx.Done()
	<-e.Poison(pid).Done()

	log.Println("rider_service: done")
}
