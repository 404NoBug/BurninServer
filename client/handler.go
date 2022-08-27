package main

import (
	"BurninProject/network"
	"BurninProject/network/protocol/gen/playerMsg"
	"fmt"
	"google.golang.org/protobuf/proto"
	"strconv"
)

type MessageHandler func(packet *network.ClientPacket)

type InputHandler func(param *InputParam)

//CreatePlayer 创建角色
func (c *Client) CreatePlayer(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 2 {
		return
	}

	msg := &playerMsg.C2S_CreatePlayer{
		UserName: param.Param[0],
		Password: param.Param[1],
	}

	c.Transport(id, msg)
}

func (c *Client) OnCreatePlayerRsp(packet *network.ClientPacket) {
	fmt.Println("恭喜你创建角色成功")
}

func (c *Client) Login(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 2 {
		return
	}

	msg := &playerMsg.C2S_Login{
		UserName: param.Param[0],
		Password: param.Param[1],
	}

	c.Transport(id, msg)

}

func (c *Client) OnLoginRsp(packet *network.ClientPacket) {
	rsp := &playerMsg.S2C_Login{}

	err := proto.Unmarshal(packet.Msg.Data, rsp)
	if err != nil {
		return
	}

	fmt.Println("登陆成功")
}

func (c *Client) AddFriend(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 { //""
		return
	}

	//parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	//if err != nil {
	//	return
	//}

	msg := &playerMsg.C2GS_AddFriend{
		UId: param.Param[0],
	}
	c.Transport(id, msg)
}

func (c *Client) OnAddFriendRsp(packet *network.ClientPacket) {
	fmt.Println("add friend success !!")
}

func (c *Client) DelFriend(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 1 || len(param.Param[0]) == 0 { //""
		return
	}

	//parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	//if err != nil {
	//	return
	//}

	msg := &playerMsg.C2GS_DelFriend{
		UId: param.Param[0],
	}

	c.Transport(id, msg)
}

func (c *Client) OnDelFriendRsp(packet *network.ClientPacket) {
	fmt.Println("you have del friend success")

}

func (c *Client) SendChatMsg(param *InputParam) {
	id := c.GetMessageIdByCmd(param.Command)

	if len(param.Param) != 3 { //""
		return
	}

	//parseUint, err := strconv.ParseUint(param.Param[0], 10, 64)
	//if err != nil {
	//	return
	//}
	parseInt32, err := strconv.ParseInt(param.Param[2], 10, 32)
	if err != nil {
		return
	}

	msg := &playerMsg.C2GS_SendChatMsg{
		UId: param.Param[0],
		Msg: &playerMsg.ChatMessage{
			Content: param.Param[1],
			Extra:   nil,
		},
		Category: int32(parseInt32),
	}

	c.Transport(id, msg)
}

func (c *Client) OnSendChatMsgRsp(packet *network.ClientPacket) {
	fmt.Println("send  chat message success")

}
