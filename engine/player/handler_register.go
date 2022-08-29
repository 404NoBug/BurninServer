package player

import (
	"BurninProject/network/protocol/gen/messageId"
)

func (p *Player) HandlerRegister() {
	p.handlers[messageId.MessageId_C2GS_AddFriend] = p.AddFriend
	p.handlers[messageId.MessageId_C2GS_DelFriend] = p.DelFriend
	p.handlers[messageId.MessageId_C2GS_SendChatMsg] = p.ResolveChatMsg
	p.handlers[messageId.MessageId_C2GS_EnterSence] = p.PlayerEnter
	p.handlers[messageId.MessageId_C2GS_PlayerMove] = p.PlayerMove
	p.handlers[messageId.MessageId_C2GS_PlayerStopMove] = p.PlayerStopMove
}
