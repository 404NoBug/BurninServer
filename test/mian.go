package main

import (
	"fmt"
	console "github.com/AsynkronIT/goconsole"
	"github.com/asynkron/protoactor-go/actor"
)

type Hello struct{ Who string }

type Actor = actor.Actor

type HelloActor struct {
	actor.Actor
}

func NewHelloActor() Actor {
	return &HelloActor{}
}

func (state *HelloActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

type WorldActor struct {
	actor.Actor
}

func NewWorldActor() Actor {
	return &WorldActor{}
}

func (state *WorldActor) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("World %v\n", msg.Who)
	}
}

func main() {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(NewHelloActor)
	pid := system.Root.Spawn(props)
	system.Root.Send(pid, &Hello{Who: "Roger"})
	_, _ = console.ReadLine()
}
