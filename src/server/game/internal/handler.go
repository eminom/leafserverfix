
package internal


import (
	"github.com/name5566/leaf/log"
	"github.com/name5566/leaf/gate"
	"server/msg"  // Local
	"reflect"
)

//~ This is what the framework calls.
func init(){
	skeleton.RegisterChanRPC(reflect.TypeOf(&msg.Hello{}), handleHello)
}


//~ Process and reply(to client)
func handleHello(args[]interface{}){
	m := args[0].(*msg.Hello)
	a := args[1].(gate.Agent)
	log.Debug("Received <%v>", m.Name) //~ Print all.  (Not just m.Name)
	switch randNo.Intn(2){
	case 0:
		log.Debug("msg.Hello sent")
		a.WriteMsg(&msg.Hello{
			Name:"God",
		})
	default:
		log.Debug("msg.Gate sent")
		a.WriteMsg(&msg.Gate{
			Host:"www.blizzard.com",
		})
	}
}


