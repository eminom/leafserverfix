package internal

import (
	"github.com/name5566/leaf/module"
	"server/base"
	"math/rand"
	"time"
)

var (
	skeleton = base.NewSkeleton()
	ChanRPC  = skeleton.ChanRPCServer
	randNo   = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type Module struct {
	*module.Skeleton
}

func (m *Module) OnInit() {
	m.Skeleton = skeleton
}

func (m *Module) OnDestroy() {

}
