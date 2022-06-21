package main

import (
	"time"

	"github.com/lkjfrf/TestTool/core"
)

type Player struct {
	id       int
	password int
}

func main() {
	core.GetTestController().Init()
	core.GetNetworkCore().Init()

	core.GetTestController().StartTesting()
	//core.GetNetworkCore().RecvPacket()

	time.Sleep(time.Hour * 1)
}
