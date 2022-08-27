package world

import "BurninProject/network/protocol/gen/messageId"

func (mm *MgrMgr) HandlerRegister() {
	mm.Handlers[messageId.MessageId_C2S_CreatePlayer] = mm.CreatePlayer
	mm.Handlers[messageId.MessageId_C2S_Login] = mm.UserLogin
}
