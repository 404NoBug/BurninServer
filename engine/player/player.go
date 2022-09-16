package player

import (
	"BurninProject/engine/common"
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"
)

type Player struct {
	PlayerInfo     *Player_Info
	HandlerParamCh chan *network.Message
	handlers       map[messageId.MessageId]Handler
	Session        *network.Session
	Broadcast      chan *network.Message
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
