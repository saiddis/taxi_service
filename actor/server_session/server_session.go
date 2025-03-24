package serversession

import (
	"context"
	"log"

	"github.com/anthdm/hollywood/actor"
	"github.com/saiddis/taxi_service/types"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type ServerSession struct {
	db *mongo.Client
}

func (r *ServerSession) Receive(c *actor.Context) {
	switch c.Message().(type) {
	case actor.Started:
		log.Println("RiderEngine started")
	case actor.Stopped:
		if err := r.db.Disconnect(c.Context()); err != nil {
			log.Println("error disconnecting with the database")
		}
		log.Println("RiderEngine stopped")
	case types.Ping:
		var result bson.M
		if err := r.db.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Decode(&result); err != nil {
			c.Respond(err)
			return
		}
		c.Respond("pong")
	}
}

func New(mongoConnStr string) (actor.Producer, error) {
	db, err := mongo.Connect(options.Client().ApplyURI(mongoConnStr))
	if err != nil {
		return nil, err
	}
	return func() actor.Receiver {
		return &ServerSession{db: db}
	}, nil
}
