package player

import (
	"BurninProject/engine/common"
	"BurninProject/engine/post"
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"
	"fmt"
	"github.com/asynkron/protoactor-go/actor"
)

type Actor = actor.Actor

type Player struct {
	PlayerInfo     *Player_Info
	HandlerParamCh chan *network.Message
	handlers       map[messageId.MessageId]Handler
	Session        *network.Session
	Broadcast      chan *network.Message
	Pid            *actor.PID
}
type Player_Info struct {
	UId        common.EntityID
	FriendList []string //朋友
	Hp         uint32
	X          float32
	Y          float32
	Dis        float32
	UIDDes     string
}

type Hello struct{ Who string }

func NewPlayer(entityID common.EntityID) *Player {
	if entityID == "" {
		entityID = common.GenEntityID()
	}
	pi := &Player_Info{
		UId:        entityID,
		Hp:         100,
		FriendList: make([]string, 100),
	}
	p := &Player{
		PlayerInfo:     pi,
		handlers:       make(map[messageId.MessageId]Handler),
		HandlerParamCh: make(chan *network.Message, 10),
	}
	p.HandlerRegister()
	return p
}

func (player *Player) NewPlayerActor() Actor {
	return player
}

//接收消息
func (player *Player) Receive(context actor.Context) {
	switch msg := context.Message().(type) {
	case Hello:
		fmt.Printf("Hello %v\n", msg.Who)
	}
}

//生成Player的ActorPID
func (player *Player) CreatPlayerPID() *actor.PID {
	system := actor.NewActorSystem()
	props := actor.PropsFromProducer(player.NewPlayerActor)
	return system.Root.Spawn(props)
}

func (p *Player) Run() {
	p.loop()
}

func (p *Player) loop() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[messageId.MessageId(handlerParam.ID)]; ok {
				fn(handlerParam)
			}
		}
		post.Tick()
	}
}

func (p *Player) OnLogin() {
	//从db加载数据初始化
	//同步数据给客户端

}

func (p *Player) OnLogout() {
	//存db
}
