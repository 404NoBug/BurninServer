package network

import (
	"BurninProject/aop/logger"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type Session struct {
	UId            uint64
	Conn           net.Conn
	IsClose        bool
	packer         IPacker
	WriteCh        chan *Message
	IsPlayerOnline bool
	MessageHandler func(packet *SessionPacket)
	//
}

func NewSession(conn net.Conn) *Session {
	return &Session{
		Conn:    conn,
		packer:  &NormalPacker{ByteOrder: binary.BigEndian},
		WriteCh: make(chan *Message, 10)}
}

func (s *Session) Run() {
	logger.Logger.InfoF("Session  Run:", s)
	go s.Read()
	go s.Write()

}

func (s *Session) Read() {
	for {
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
