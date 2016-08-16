package internal

import (
	"github.com/name5566/leaf/gate"
	"server/conf"
	"server/game"
	"server/msg"
)

type Module struct {
	*gate.Gate
}

func (m *Module) OnInit() {
	m.Gate = &gate.Gate{
		MaxConnNum:      conf.Server.MaxConnNum,
		PendingWriteNum: conf.PendingWriteNum,
		MaxMsgLen:       conf.MaxMsgLen,
		WSAddr:          conf.Server.WSAddr,
		HTTPTimeout:     conf.HTTPTimeout,
		TCPAddr:         conf.Server.TCPAddr,
		LenMsgLen:       conf.LenMsgLen,
		LittleEndian:    conf.LittleEndian,
		Processor:       msg.Processor,
		AgentChanRPC:    game.ChanRPC,
	}
}

//在这里实现Module.OnDestroy并不稀奇...
//但是在父类中实现?? (从语意来讲, 的确是继承了!) 如果把此处Module视为gate.Gate的指针, 那么调用方法.OnDestroy也是成立的. 
