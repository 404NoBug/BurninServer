package main

import (
	"BurninProject/network/protocol/gen/player"
	"fmt"
	"google.golang.org/protobuf/proto"
)

func main() {
	cc := &player.CSCreateUser{
		UserName: "An",
		Password: "1234789",
	}
	marshal, err := proto.Marshal(cc)
	if err != nil {
		fmt.Println("proto.Marshal err = ", err)
		return
	}
	fmt.Println("msg = ", marshal)
	msg := &player.CSCreateUser{}
	if err := proto.Unmarshal(marshal, msg); err == nil {
		fmt.Println("msg = ", msg)
	}

}
