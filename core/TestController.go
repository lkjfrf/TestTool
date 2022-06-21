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
	TestServerIPs         []string
	TestServerConnections []*net.TCPConn
	TestPeopleCount       int
	Packets               interface{}
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

	tc.TestServerIPs = []string{"150.109.254.215:8000"}
	//	tc.TestServerIPs = []string{"121.162.7.67:7002"}
	tc.TestPeopleCount = 600

}

func (tc *TestController) StartTesting() {
	log.Println("Start Testing")
	tc.TestServerConnections = GetNetworkCore().Connect(tc.TestServerIPs)
	tc.TestersLogin()
	tc.TesterMove()
}

func (tc *TestController) TestersLogin() {
	packet := content.New_FSendPacket_TestPlayerLogin()
	for i := 0; i < tc.TestPeopleCount; i++ {
		packet.Id = strconv.Itoa(i)
		packet.IsMan = true
		packet.LevelName = "MusicCastle"
		packet.UserName = strconv.Itoa(i)
		GetNetworkCore().SendPacket(tc.TestServerConnections, packet)
		time.Sleep(time.Microsecond * 3)
	}
}

func (tc *TestController) TesterMove() {
	packet := content.New_FSendPacket_PlayerMove()
	packet.MoveSpeed = 800
	packet.RotateSpeed = 300
	FlipFlop := 0

	go func() {
		for {
			for i := 0; i < tc.TestPeopleCount; i++ {
				packet.Id = strconv.Itoa(i)
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

				GetNetworkCore().SendPacket(tc.TestServerConnections, packet)
				//log.Println("Id : ", packet.Id, " = ", packet.Position)
				time.Sleep(time.Microsecond * 3)
			}
		}
	}()

	go func() {
		packet := content.New_FSendPacket_HeartBeat()
		for {
			for i := 0; i < tc.TestPeopleCount; i++ {
				packet.Id = strconv.Itoa(i)
				GetNetworkCore().SendPacket(tc.TestServerConnections, packet)
				//log.Println("Id : ", packet.Id, " = ", packet.PacketName)
			}
			time.Sleep(time.Second * 10)
		}
	}()

	go func() {
		packet := content.New_FSendPacket_NormalChat()
		for {
			for i := 0; i < tc.TestPeopleCount; i++ {
				packet.Id = strconv.Itoa(i)
				packet.UserName = strconv.Itoa(i) + "User"
				packet.Message = "ioawef;ejfeajiofji23ijj;ia3wjawj389j9a3w889aw8j9wj38;aw8;48jawe8f8ja wj8fj qawjfaw3f8q8jw4f;jqwaifja;wio3jf;oiawj;ojㅈㄷㅁ;ㅐㅑㅓㅁㅈㄷ;ㅐ랴ㅓㅁㅈㄷ;ㅐㅑ럼;ㅈ대ㅑ러;머8ㅈ3ㅁㅈ;ㅐ38ㅓㅁ;ㅈ3ㅐㅑㅓ;"
				GetNetworkCore().SendPacket(tc.TestServerConnections, packet)
				log.Println("Id : ", packet.Id, " = ", packet.PacketName)
				time.Sleep(time.Microsecond * 3)
			}
		}
	}()

}
