package world

import (
	logicPlayer "BurninProject/engine/player"
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"
	"BurninProject/network/protocol/gen/playerMsg"
	"google.golang.org/protobuf/proto"
	"time"
)

func (mm MgrMgr) MsgPing(message *network.SessionPacket) {
	msg := &playerMsg.GS2C_MsgPong{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	message.Sess.LastPingTime = uint64(time.Now().Unix())
	mm.SendMsg(uint64(messageId.MessageId_C2GS_MsgPong), &playerMsg.GS2C_MsgPong{}, message.Sess)
}

func (mm MgrMgr) CreateAccount(message *network.SessionPacket) {
	req := &playerMsg.C2S_Register_Accoount{}
	err := proto.Unmarshal(message.Msg.Data, req)
	if err != nil {
		return
	}
	if mm.MongoDb.FindOne(req.UserAccoount, req.Password) != nil {

	} else {
		mm.SendMsg(uint64(messageId.MessageId_S2C_Register_Accoount), &playerMsg.S2C_Register_Accoount{RetCode: 1}, message.Sess)
	}
}

func (mm *MgrMgr) CreatePlayer(message *network.SessionPacket) {
	msg := &playerMsg.S2C_CreatePlayer{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	mm.SendMsg(uint64(messageId.MessageId_S2C_CreatePlayer), &playerMsg.S2C_CreatePlayer{}, message.Sess)
}

func (mm *MgrMgr) UserLogin(message *network.SessionPacket) {
	msg := &playerMsg.C2S_Login{}
	err := proto.Unmarshal(message.Msg.Data, msg)
	if err != nil {
		return
	}
	newPlayer := logicPlayer.NewPlayer()
	newPlayer.PlayerInfo.UId = uint64(time.Now().Unix())
	message.Sess.IsPlayerOnline = true
	message.Sess.UId = newPlayer.PlayerInfo.UId
	newPlayer.Session = message.Sess
	newPlayer.Broadcast = mm.Pm.Broadcast
	mm.Pm.Add(newPlayer)

}

func (mm *MgrMgr) SendMsg(id uint64, message proto.Message, session *network.Session) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	rsp := &network.Message{
		ID:   id,
		Data: bytes,
	}
	session.SendMsg(rsp)
}
