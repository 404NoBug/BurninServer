package manager

import (
	"BurninProject/engine/player"
	"BurninProject/network"
	"fmt"
)

//PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players   map[uint64]*player.Player
	addPCh    chan *player.Player
	Broadcast chan *network.Message
}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		players:   make(map[uint64]*player.Player),
		addPCh:    make(chan *player.Player, 1),
		Broadcast: make(chan *network.Message),
	}
}

func (pm *PlayerMgr) Add(p *player.Player) {
	if pm.players[p.UId] != nil {
		return
	}
	fmt.Println("AddAddAddAddAddAddAdd")
	pm.players[p.UId] = p
	go p.Run()
}

//Del ...
func (pm *PlayerMgr) Del(p player.Player) {
	delete(pm.players, p.UId)
}

//Broadcast Message
func (pm *PlayerMgr) BroadcastSend(msg *network.Message) {
	fmt.Println("BroadcastSend = ", msg)
	for _, v := range pm.players {
		v.Session.SendMsg(msg)
	}
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		case broadcastMsg := <-pm.Broadcast:
			pm.BroadcastSend(broadcastMsg)
		}
	}
}

func (pm *PlayerMgr) GetPlayer(uId uint64) *player.Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
