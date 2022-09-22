package world

import (
	"BurninProject/aop/logger"
	"BurninProject/engine/config"
	"BurninProject/engine/manager"
	"BurninProject/network"
	"BurninProject/network/protocol/gen/messageId"
	"os"
	"syscall"
)

type MgrMgr struct {
	Pm     *manager.PlayerMgr
	Server *network.Server
	//MongoDb         *MongoDB.MongoConn
	Handlers        map[messageId.MessageId]func(message *network.SessionPacket)
	chSessionPacket chan *network.SessionPacket
}

func NewMgrMgr() *MgrMgr {
	cfg := config.GetGate(1)
	m := &MgrMgr{Pm: manager.NewPlayerMgr()}
	m.Server = network.NewServer(cfg.ListenAddr)
	m.Server.OnSessionPacket = m.OnSessionPacket
	m.Handlers = make(map[messageId.MessageId]func(message *network.SessionPacket))
	//m.MongoDb = MongoDB.InitMongoConn("127.0.0.1:27017", "locahost", "123456", "Burnin")
	return m
}

var MM *MgrMgr

func (mm *MgrMgr) Run() {
	mm.HandlerRegister()
	go mm.Server.Run()
	go mm.Pm.Run()
}

func (mm *MgrMgr) OnSessionPacket(packet *network.SessionPacket) {
	if packet.Msg == nil {
		if packet.Sess.UId != "" {
			mm.Pm.DelPCh <- &packet.Sess.UId
		}
		return
	}
	if handler, ok := mm.Handlers[messageId.MessageId(packet.Msg.ID)]; ok {
		handler(packet)
		return
	}
	if p := mm.Pm.GetPlayer(packet.Sess.UId); p != nil {
		p.HandlerParamCh <- packet.Msg
	}
}

func (mm *MgrMgr) OnSystemSignal(signal os.Signal) bool {
	logger.Logger.DebugF("[MgrMgr] 收到信号 %v \n", signal)
	tag := true
	switch signal {
	case syscall.SIGHUP:
		//todo
	case syscall.SIGPIPE:
	default:
		logger.Logger.DebugF("[MgrMgr] 收到信号准备退出...")
		tag = false

	}
	return tag
}
