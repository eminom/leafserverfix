package gate

import "server/msg"
import "server/game"

func init() {
	msg.Processor.SetRouter(&msg.Hello{}, game.ChanRPC)
}
