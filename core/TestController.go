package core

import (
	"log"
	"math/rand"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/lkjfrf/TestTool/content"
	"github.com/lkjfrf/TestTool/helper"
)

type TestController struct {
	TestServerIP             string
	TestServerUDP            string
	TestServerConnections    []*net.TCPConn
	TestUDPServerConnections []*net.UDPConn
	TestPeopleCount          int
	Packets                  interface{}

	wg     sync.WaitGroup
	moveWg sync.WaitGroup
}

var Testinstance *TestController
var Testonce sync.Once

func GetTestController() *TestController {
	once.Do(func() {
		Testinstance = &TestController{}
	})
	return Testinstance
}

func (tc *TestController) Init() {
	log.Println("INIT_TestController")

	//tc.TestServerIP = "121.162.7.67:8000"
	tc.TestServerUDP = "192.168.0.6:8005"
	tc.TestPeopleCount = 900

	tc.moveWg = sync.WaitGroup{}
}

func (tc *TestController) StartTesting() {
	log.Println("Start Testing")
	//tc.TestServerConnections = GetNetworkCore().Connect(tc.TestServerIP, tc.TestPeopleCount)
	tc.TestUDPServerConnections = GetNetworkCore().ConnectUDP(tc.TestServerUDP, tc.TestPeopleCount)
	//if tc.TestersLogin() {
	//tc.TesterMove()
	tc.TestersChannelEnter()
	tc.WatchToggle()
	//go tc.HeartBeat()
	//go tc.ChatTest()
	//}
}

func (tc *TestController) TestersLogin() bool {
	packet := content.S_DBSignin{}
	for i := 0; i < tc.TestPeopleCount; i++ {
		packet.Id = "tester" + strconv.Itoa(i)
		GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, content.ETestPlayerLogin)
		time.Sleep(time.Microsecond * 100)
	}
	return true
}
func (tc *TestController) TestersChannelEnter() bool {
	packet2 := content.S_ChannelEnter{}
	for i := 0; i < tc.TestPeopleCount; i++ {
		packet2.Id = "tester" + strconv.Itoa(i)
		packet2.ChannelNum = 11
		packet2.ChannelType = 2
		//GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet2, content.ChannelEnter)
		GetNetworkCore().SendUDPPacket(tc.TestUDPServerConnections[i], packet2, content.ChannelEnter)
		log.Println("Send ChannelEnter ", packet2.Id)
		time.Sleep(time.Microsecond * 100)
	}
	return true
}

func (tc *TestController) TesterMove() {
	packet := content.S_PlayerMove{}
	packet.MoveSpeed = 800
	packet.RotateSpeed = 300
	FlipFlop := 0

	go func() {
		for {
			tc.moveWg.Add(tc.TestPeopleCount)
			go func() {
				for i := 0; i < tc.TestPeopleCount; i++ {

					packet.Id = "tester" + strconv.Itoa(int(i))
					if FlipFlop == 1 {
						packet.Rotation = helper.NewVector3(0, 0, 0)
						packet.Position = helper.NewVector3(0, 0, 200)
					} else if FlipFlop == 2 {
						packet.Rotation = helper.NewVector3(0, 0, 0)
						packet.Position = helper.NewVector3(400, 0, 200)
					} else {
						packet.Rotation = helper.NewVector3(0, 0, 0)
						packet.Position = helper.NewVector3(400, 400, 200)
					}

					FlipFlop = rand.Intn(3)

					//log.Println("Id : ", packet.Id, " = ", packet.Position)
					GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, content.PlayerMove)
				}
				tc.moveWg.Done()
			}()
			time.Sleep(time.Microsecond * 1000)
		}
		tc.moveWg.Wait()
	}()
}

func (tc *TestController) HeartBeat() {
	packet := content.S_HeartBeat{}
	for {
		for i := 0; i < tc.TestPeopleCount; i++ {
			packet.Id = "tester" + strconv.Itoa(int(i))
			GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, content.HeartBeat)
			time.Sleep(time.Microsecond * 10)
			//log.Println("Id : ", packet.Id, " = ", packet.PacketName)
		}
		time.Sleep(time.Second * 10)
	}
}

func (tc *TestController) ChatTest() {
	packet := content.S_NormalChat{}
	j := 0
	for {
		for i := 0; i < tc.TestPeopleCount; i++ {
			j++
			packet.Id = "tester" + strconv.Itoa(int(i))
			packet.Message = strconv.Itoa(j)
			GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, content.NormalChat)
			time.Sleep(time.Microsecond * 10)
		}
	}
}

func (tc *TestController) LogOut() {
	packet := content.S_PlayerLogout{}
	j := 0
	for {
		for i := 0; i < tc.TestPeopleCount; i++ {
			j++
			packet.Id = "tester" + strconv.Itoa(int(i))
			GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, content.NormalChat)
			time.Sleep(time.Microsecond * 10)
		}
	}
}

func (tc *TestController) WatchToggle() {
	packet := content.S_ScreenWatchToggle{}
	for i := 0; i < tc.TestPeopleCount; i++ {
		packet.Id = "tester" + strconv.Itoa(int(i))
		packet.IsOn = true
		packet.ChannelNum = 11
		GetNetworkCore().SendUDPPacket(tc.TestUDPServerConnections[i], packet, content.EScreenWatchToggle)
		time.Sleep(time.Microsecond * 10)
	}
}
