package manager

import (
	"BurninProject/engine/common"
	"BurninProject/engine/player"
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"
	"BurninProject/network/protocol/gen/playerMsg"
	"fmt"
	"google.golang.org/protobuf/proto"
)

//PlayerMgr 维护在线玩家
type PlayerMgr struct {
	players   map[common.EntityID]*player.Player
	addPCh    chan *player.Player
	DelPCh    chan *common.EntityID
	Broadcast chan *network.Message
}

func NewPlayerMgr() *PlayerMgr {
	return &PlayerMgr{
		players:   make(map[common.EntityID]*player.Player),
		addPCh:    make(chan *player.Player, 1),
		DelPCh:    make(chan *common.EntityID, 1),
		Broadcast: make(chan *network.Message),
	}
}

func (pm *PlayerMgr) Add(p *player.Player) {
	if pm.players[p.PlayerInfo.UId] != nil {
		return
	}
	fmt.Println("Add Player")
	pm.players[p.PlayerInfo.UId] = p
	go p.Run()
}

//Del ...
func (pm *PlayerMgr) Del(UID common.EntityID) {
	fmt.Println("Del Player", UID)
	pm.SendPlayerLeaveGame(UID)
}

//Broadcast Message
func (pm *PlayerMgr) BroadcastSend(msg *network.Message) {
	switch msg.ID {
	case uint64(messageId.MessageId_GS2C_EnterSence):
		{
			pm.SendOnLinePlayerList()
		}
	default:
		for _, v := range pm.players {
			v.Session.SendMsg(msg)
		}
	}
}

func (pm *PlayerMgr) SendPlayerLeaveGame(UID common.EntityID) {
	Player := pm.players[UID]
	if Player == nil {
		return
	}
	Msg := &playerMsg.GS2C_PlayerLeave{
		UId: Player.PlayerInfo.UIDDes,
	}
	bytes, err := proto.Marshal(Msg)
	if err != nil {
		return
	}
	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_PlayerLeave),
		Data: bytes,
	}
	delete(pm.players, UID)
	for _, v := range pm.players {
		v.Session.SendMsg(rsp)
	}
}

func (pm *PlayerMgr) SendOnLinePlayerList() {
	Msg := &playerMsg.GS2C_ONLinePlayerList{}
	for _, v := range pm.players {
		posInfo := &playerMsg.PosInfo{
			X: v.PlayerInfo.X,
			Y: v.PlayerInfo.Y,
		}
		info := &playerMsg.ONLinePlayer_Info{
			UId: v.PlayerInfo.UIDDes,
			Pos: posInfo,
			Dir: v.PlayerInfo.Dis,
		}
		Msg.List = append(Msg.List, info)
	}
	bytes, err := proto.Marshal(Msg)
	if err != nil {
		return
	}
	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_ONLinePlayerList),
		Data: bytes,
	}
	for _, v := range pm.players {
		v.Session.SendMsg(rsp)
	}
}

func (pm *PlayerMgr) Run() {
	for {
		select {
		case p := <-pm.addPCh:
			pm.Add(p)
		case p := <-pm.DelPCh:
			pm.Del(*p)
		case broadcastMsg := <-pm.Broadcast:
			pm.BroadcastSend(broadcastMsg)
		}
	}
}

func (pm *PlayerMgr) GetPlayer(uId common.EntityID) *player.Player {
	p, ok := pm.players[uId]
	if ok {
		return p
	}
	return nil
}
