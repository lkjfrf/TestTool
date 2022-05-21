package main

import (
	"github.com/lkjfrf/TestTool/content"
	"github.com/lkjfrf/TestTool/core"
)

type Player struct {
	id       int
	password int
}

func main() {
	content.GetTestController().Init()
	core.GetNetworkCore().Init()

	content.GetTestController().StartTesting()
}
