package main

import (
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"

	"google.golang.org/protobuf/proto"
)

func (c *Client) InputHandlerRegister() {
	c.inputHandlers[messageId.MessageId_C2S_CreatePlayer.String()] = c.CreatePlayer
	c.inputHandlers[messageId.MessageId_C2S_Login.String()] = c.Login
	c.inputHandlers[messageId.MessageId_C2GS_AddFriend.String()] = c.AddFriend
	c.inputHandlers[messageId.MessageId_C2GS_DelFriend.String()] = c.DelFriend
	c.inputHandlers[messageId.MessageId_C2GS_SendChatMsg.String()] = c.SendChatMsg
}

func (c *Client) GetMessageIdByCmd(cmd string) messageId.MessageId {
	mid, ok := messageId.MessageId_value[cmd]
	if ok {
		return messageId.MessageId(mid)
	}
	return messageId.MessageId_None
}

func (c *Client) Transport(id messageId.MessageId, message proto.Message) {
	bytes, err := proto.Marshal(message)
	if err != nil {
		return
	}
	c.cli.ChMsg <- &network.Message{
		ID:   uint64(id),
		Data: bytes,
	}
}
