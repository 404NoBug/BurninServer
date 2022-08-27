package player

import (
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"
)

type Player struct {
	UId            uint64
	FriendList     []string //朋友
	Hp             uint32
	X              float32
	Y              float32
	Dis            float32
	HandlerParamCh chan *network.Message
	handlers       map[messageId.MessageId]Handler
	Session        *network.Session
	Broadcast      chan *network.Message
}

func NewPlayer() *Player {
	p := &Player{
		UId:            0,
		Hp:             100,
		FriendList:     make([]string, 100),
		handlers:       make(map[messageId.MessageId]Handler),
		HandlerParamCh: make(chan *network.Message, 10),
	}
	p.HandlerRegister()
	return p
}

func (p *Player) Run() {
	for {
		select {
		case handlerParam := <-p.HandlerParamCh:
			if fn, ok := p.handlers[messageId.MessageId(handlerParam.ID)]; ok {
				fn(handlerParam)
			}
		}
	}
}

func (p *Player) OnLogin() {
	//从db加载数据初始化
	//同步数据给客户端

}

func (p *Player) OnLogout() {
	//存db
}
