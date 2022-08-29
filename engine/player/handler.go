package player

import (
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"
	"BurninProject/network/protocol/gen/playerMsg"
	"github.com/phuhao00/sugar"

	"google.golang.org/protobuf/proto"
)

type Handler func(*network.Message)

func (p *Player) AddFriend(packet *network.Message) {
	req := &playerMsg.C2GS_AddFriend{}
	type Handler func(packet *network.Message)

	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}

	if !sugar.CheckInSlice(req.UId, p.FriendList) {
		p.FriendList = append(p.FriendList, req.UId)
	}

	bytes, err := proto.Marshal(&playerMsg.GS2C_SendChatMsg{})
	if err != nil {
		return
	}

	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_AddFriend),
		Data: bytes,
	}

	p.Session.SendMsg(rsp)
}

func (p *Player) DelFriend(packet *network.Message) {
	req := &playerMsg.C2GS_DelFriend{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	p.FriendList = sugar.DelOneInSlice(req.UId, p.FriendList)

	bytes, err := proto.Marshal(&playerMsg.GS2C_DelFriend{})
	if err != nil {
		return
	}

	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_DelFriend),
		Data: bytes,
	}

	p.Session.SendMsg(rsp)
}

func (p *Player) ResolveChatMsg(packet *network.Message) {

	req := &playerMsg.C2GS_SendChatMsg{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}

	bytes, err := proto.Marshal(&playerMsg.GS2C_SendChatMsg{
		Msg: req.Msg,
	})
	if err != nil {
		return
	}

	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_SendChatMsg),
		Data: bytes,
	}
	//p.Broadcast <- rsp
	p.Session.SendMsg(rsp)
}

func (p *Player) PlayerEnter(packet *network.Message) {
	req := &playerMsg.C2GS_EnterSence{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	p.X = req.Pos.X
	p.Y = req.Pos.Y
	p.Dis = req.Dir
	p.UIDDes = req.UId
	posInfo := &playerMsg.PosInfo{
		X: req.Pos.X,
		Y: req.Pos.Y,
	}
	bytes, err := proto.Marshal(&playerMsg.GS2C_EnterSence{
		UId: req.UId,
		Pos: posInfo,
		Dir: req.Dir,
	})
	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_EnterSence),
		Data: bytes,
	}
	p.Broadcast <- rsp
}

func (p *Player) PlayerMove(packet *network.Message) {
	req := &playerMsg.C2GS_PlayerMove{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	posInfo := &playerMsg.PosInfo{
		X: req.Pos.X,
		Y: req.Pos.Y,
	}
	p.X = req.Pos.X
	p.Y = req.Pos.Y
	//p.Dis = req.Dir
	bytes, err := proto.Marshal(&playerMsg.GS2C_PlayerMove{
		UId: p.UIDDes,
		Pos: posInfo,
	})
	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_PlayerMove),
		Data: bytes,
	}
	p.Broadcast <- rsp
	//p.Session.SendMsg(rsp)
}
func (p *Player) PlayerStopMove(packet *network.Message) {
	req := &playerMsg.C2GS_PlayerStopMove{}
	err := proto.Unmarshal(packet.Data, req)
	if err != nil {
		return
	}
	p.Dis = req.Dir
	bytes, err := proto.Marshal(&playerMsg.GS2C_PlayerStopMove{
		UId: p.UIDDes,
		Dir: req.Dir,
	})
	rsp := &network.Message{
		ID:   uint64(messageId.MessageId_GS2C_PlayerStopMove),
		Data: bytes,
	}
	p.Broadcast <- rsp
}
