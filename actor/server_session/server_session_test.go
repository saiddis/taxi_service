package serversession_test

import (
	"flag"
	"testing"
	"time"

	"github.com/anthdm/hollywood/actor"
	serversession "github.com/saiddis/taxi_service/actor/server_session"
	"github.com/saiddis/taxi_service/types"
)

var mongoConnStr = flag.String("db", "", "mongodb connection string")

type testFunc func(tb testing.TB, e *actor.Engine, pid *actor.PID)

func TestServerSessionStart(t *testing.T) {
	if *mongoConnStr == "" {
		t.Fatal("connection string is not set")
	}
	producer, err := serversession.New(*mongoConnStr)
	if err != nil {
		t.Fatal(err)
	}
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		t.Fatal(err)
	}
	pid := e.Spawn(producer, "serversession")

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
	if *mongoConnStr == "" {
		tb.Skip("connection string is not set; skipping test")
	}

	producer, err := serversession.New(*mongoConnStr)
	if err != nil {
		tb.Fatal(err)
	}
	e, err := actor.NewEngine(actor.NewEngineConfig())
	if err != nil {
		tb.Fatal(err)
	}
	pid := e.Spawn(producer, "serversession")

	defer func() {
		<-e.Poison(pid).Done()
	}()

	f(tb, e, pid)
}
