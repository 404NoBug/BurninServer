package main

import (
	"BurninProject/engine/kvdb"
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
	kvdb.GetOrPut("password$"+"gy001", "123456", func(oldVal string, err error) {
		if err != nil {
			fmt.Println("11")
			return
		}
		if oldVal == "" {
			fmt.Println("22")
			//player := goworld.CreateEntityLocally("Player") // 创建一个Player对象然后立刻销毁，产生一次存盘
			//player.Attrs.SetStr("name", username)
			//player.Destroy()
		} else {
			fmt.Println("33")
		}
	})
	console.ReadLine()
}
