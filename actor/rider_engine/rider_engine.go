package riderengine

import (
	"log"

	"github.com/anthdm/hollywood/actor"
)

type RiderEngine struct {
}

func (r *RiderEngine) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
		log.Println("RiderEngine started")
	case actor.Stopped:
		log.Println("RiderEngine stopped")
	}
}

func New() actor.Producer {
	return func() actor.Receiver {
		return &RiderEngine{}
	}
}
