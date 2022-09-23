package network

import (
	"BurninProject/aop/logger"
	"BurninProject/engine/common"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

const pingInterval = 2

type Session struct {
	UId            common.EntityID
	Conn           net.Conn
	IsClose        bool
	packer         IPacker
	WriteCh        chan *Message
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
	LastPingTime   uint64
	//
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		Conn:         conn,
		packer:       &NormalPacker{ByteOrder: binary.BigEndian},
		WriteCh:      make(chan *Message, 10),
		LastPingTime: uint64(time.Now().Unix()),
	}
}

func (s *Session) Run() {
	logger.Logger.InfoF("Session  Run:", s)
	go s.Read()
	go s.Write()
	go s.Ticker()
}

func (s *Session) Ticker() {
	myTimer := time.NewTimer(time.Second) // 启动定时器
	//defer func(Conn net.Conn) {
	//	err := Conn.Close()
	//	logger.Logger.InfoF("Conn.Close = ", err)
	//	if err != nil {
	//		logger.Logger.InfoF("CheckPing Conn.Close err = ", err)
	//	}
	//}(s.Conn)
ForEnd:
	for {
		select {
		case <-myTimer.C:
			if s.CheckPing() {
				break ForEnd
			}
			myTimer.Reset(time.Second) // 每次使用完后需要人为重置下
		}
	}
	// 不再使用了，结束它
	myTimer.Stop()
}

//Ping检查
func (s *Session) CheckPing() bool {
	//现在的时间戳
	timeNow := uint64(time.Now().Unix())
	if timeNow-s.LastPingTime > pingInterval*4 {
		return true
	} else {
		return false
	}
}

func (s *Session) Read() {
	for {
		if s.Conn == nil {
			return
		}
		err := s.Conn.SetReadDeadline(time.Now().Add(time.Second))
		if err != nil {
			fmt.Println(err)
			continue
		}
		message, err := s.packer.Unpack(s.Conn)
		if _, ok := err.(net.Error); ok {
			continue
		}
		s.MessageHandler(&SessionPacket{
			Msg:  message,
			Sess: s,
		})
		if err == io.EOF {
			defer func(Conn net.Conn) {
				err := Conn.Close()
				fmt.Println("Conn.Close = ", err)
				if err != nil {
					fmt.Println("Conn.Close Err = ", err)
				}
			}(s.Conn)
			fmt.Println("Read Over")
			return
		}
	}
}

func (s *Session) Write() {
	for {
		select {
		case resp := <-s.WriteCh:
			s.send(resp)
		}
	}
}

func (s *Session) send(message *Message) {
	err := s.Conn.SetWriteDeadline(time.Now().Add(time.Second))
	if err != nil {
		fmt.Println(err)
		return
	}
	bytes, err := s.packer.Pack(message)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = s.Conn.Write(bytes)
	if err != nil {
		fmt.Println(err)
	}

}

func (s *Session) SendMsg(msg *Message) {
	s.WriteCh <- msg
}
