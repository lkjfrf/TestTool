package content

import (
	"log"
	"net"
	"sync"
	"time"

	"github.com/lkjfrf/TestTool/core"
	"github.com/lkjfrf/TestTool/helper"
)

type TestController struct {
	TestServerIPs         []string
	TestServerConnections []*net.TCPConn
	TestPeopleCount       int
	Packets               interface{}
}

var instance *TestController
var once sync.Once

func GetTestController() *TestController {
	once.Do(func() {
		instance = &TestController{}
	})
	return instance
}

func (tc *TestController) Init() {
	log.Println("INIT_TestController")

	tc.TestServerIPs = []string{"172.30.1.40:8000", "pprk.ddns", ""}
	tc.TestPeopleCount = 100

}

func (tc *TestController) StartTesting() {
	log.Println("Start Testing")
	tc.TestServerConnections = core.GetNetworkCore().Connect(tc.TestServerIPs)
	tc.TestersLogin()
	tc.TesterMove()
}

func (tc *TestController) TestersLogin() {
	packet := FSendPacket_PlayerLogin{PacketName: "FSendPacket_PlayerLogin"}
	for i := 0; i < tc.TestPeopleCount; i++ {
		packet.Id = string(i)
		packet.IsMan = true
		packet.LevelName = "Main"
		packet.UserName = string(i)
		core.GetNetworkCore().SendPacket(tc.TestServerConnections, packet)
	}
}

func (tc *TestController) TesterMove() {
	packet := FSendPacket_PlayerMove{PacketName: "FSendPacket_PlayerMove"}
	packet.MoveSpeed = 800
	packet.RotateSpeed = 300
	FlipFlop := false

	for i := 0; i < tc.TestPeopleCount; i++ {
		go func() {
			packet.Id = string(i)
			for {
				if FlipFlop {
					packet.DestRotation = helper.NewVector3(0, 0, 0)
					packet.Destination = helper.NewVector3(0, 0, 0)
				} else {
					packet.DestRotation = helper.NewVector3(100, 100, 100)
					packet.Destination = helper.NewVector3(100, 100, 100)
				}
				FlipFlop = !FlipFlop

				core.GetNetworkCore().SendPacket(tc.TestServerConnections, packet)

				time.Sleep(time.Millisecond * 3)
			}
		}()
	}

}
