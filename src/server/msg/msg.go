package msg

import (
	//"github.com/name5566/leaf/network"
	"github.com/name5566/leaf/network/json"
)

//var Processor network.Processor
var Processor = json.NewProcessor()

func init() {
	Processor.Register(&Hello{})
	Processor.Register(&Gate{})
}

type Hello struct {
	Name string
}

type Gate struct {
	Host string
}
