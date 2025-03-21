package riderengine

import (
	"log"

	"github.com/anthdm/hollywood/actor"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type RiderEngine struct {
	db *mongo.Client
}

func (r *RiderEngine) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
		log.Println("RiderEngine started")

	case actor.Stopped:
		log.Println("RiderEngine stopped")
		if err := r.db.Disconnect(c.Context()); err != nil {
			log.Println("error disconnecting with the database")
		}
	}
}

func New(mongoConnStr string) (actor.Producer, error) {
	db, err := mongo.Connect(options.Client().ApplyURI(mongoConnStr))
	if err != nil {
		return nil, err
	}
	return func() actor.Receiver {
		return &RiderEngine{db: db}
	}, nil
}
