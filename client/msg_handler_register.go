package main

import "BurninProject/network/protocol/gen/messageId"

func (c *Client) MessageHandlerRegister() {
	c.messageHandlers[messageId.MessageId_S2C_CreatePlayer] = c.OnCreatePlayerRsp
	c.messageHandlers[messageId.MessageId_S2C_Login] = c.OnLoginRsp
	c.messageHandlers[messageId.MessageId_GS2C_AddFriend] = c.OnAddFriendRsp
	c.messageHandlers[messageId.MessageId_GS2C_DelFriend] = c.OnDelFriendRsp
	c.messageHandlers[messageId.MessageId_GS2C_SendChatMsg] = c.OnSendChatMsgRsp
}
