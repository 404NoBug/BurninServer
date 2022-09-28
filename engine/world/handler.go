package world

import (
	"BurninProject/engine/kvdb"
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
	username := req.UserAccoount
	passWord := req.Password
	GetOrPutKVDB("password$"+username, passWord, func(oldVal string, err error) {
		if err != nil {
			mm.SendMsg(uint64(messageId.MessageId_S2C_Register_Accoount), &playerMsg.S2C_Register_Accoount{RetCode: 3}, message.Sess) // 服务器错误
			return
		}
		if oldVal == "" {
			newPlayer := logicPlayer.NewPlayer("")
			//player := goworld.CreateEntityLocally("Player") // 创建一个Player对象然后立刻销毁，产生一次存盘
			//player.Attrs.SetStr("name", username)
			//player.Destroy()
			kvdb.Put("playerID$"+username, string(newPlayer.PlayerInfo.UId), func(err error) {
				mm.SendMsg(uint64(messageId.MessageId_S2C_Register_Accoount), &playerMsg.S2C_Register_Accoount{RetCode: 0, PlayerId: string(newPlayer.PlayerInfo.UId)}, message.Sess) // 注册成功，请点击登录
			})
		} else {
			mm.SendMsg(uint64(messageId.MessageId_S2C_Register_Accoount), &playerMsg.S2C_Register_Accoount{RetCode: 1}, message.Sess) // 抱歉，这个账号已经存在
		}
	})
}

func GetOrPutKVDB(key string, val string, callback kvdb.KVDBGetOrPutCallback) {
	kvdb.GetOrPut(key, val, callback)
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
	newPlayer := logicPlayer.NewPlayer("")
	playerPid := newPlayer.CreatPlayerPID()
	logicPlayer.PlayerSendMsg(playerPid, logicPlayer.Hello{"gaoyang"})
	newPlayer.Pid = playerPid
	message.Sess.IsPlayerOnline = true
	message.Sess.UId = newPlayer.PlayerInfo.UId
	newPlayer.Session = message.Sess
	newPlayer.Broadcast = mm.Pm.Broadcast
	mm.Pm.Add(newPlayer)
	mm.SendMsg(uint64(messageId.MessageId_S2C_Login), &playerMsg.S2C_Login{PlayerId: string(newPlayer.PlayerInfo.UId), Ok: 0}, message.Sess)
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
