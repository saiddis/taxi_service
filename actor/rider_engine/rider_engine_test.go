package riderengine_test

import (
	"testing"
	"time"

	"github.com/anthdm/hollywood/actor"
	riderengine "github.com/saiddis/taxi_service/actor/rider_engine"
	"github.com/saiddis/taxi_service/types"
)

type testFunc func(tb testing.TB, e *actor.Engine, pid *actor.PID)

func TestRiderEngineStart(t *testing.T) {
	producer := riderengine.New()
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		t.Fatal(err)
	}
	pid := e.Spawn(producer, "riderengine")

	defer func() {
		<-e.Poison(pid).Done()
	}()

	respond := e.Request(pid, types.Ping{}, time.Second)
	if result, err := respond.Result(); err != nil {
		t.Fatal(err)
	} else if msg, ok := result.(string); !ok {
		t.Fatal("extected respose of type string")
	} else if got, want := msg, "pong"; got != want {
		t.Fatalf("got ping response %s, want %s", got, want)
	}
}

func WithTestFunc(tb testing.TB, f testFunc) {
	tb.Helper()

	producer := riderengine.New()
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		tb.Fatal(err)
	}
	pid := e.Spawn(producer, "riderengine")

	defer func() {
		<-e.Poison(pid).Done()
	}()

	f(tb, e, pid)
}
