package core

import (
	"fmt"
	"log"
	"net"
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

	tc.TestServerIP = "43.155.170.138:8001"
	tc.TestPeopleCount = 100

	tc.moveWg = sync.WaitGroup{}
}

func (tc *TestController) StartTesting() {
	log.Println("Start Testing")
	tc.TestServerConnections = GetNetworkCore().Connect(tc.TestServerIP, tc.TestPeopleCount)
	tc.MatchingDone()
	time.Sleep(time.Second * 3)
	tc.TestersLogin()
	time.Sleep(time.Second * 3)
	tc.TesterAttack()
}

func (tc *TestController) MatchingDone() {
	var players []content.Player
	for i := 0; i < tc.TestPeopleCount-1; i++ {
		iString := fmt.Sprintf("%d", i)
		player := content.Player{Playerid: iString, Teamid: 1, Nickname: iString}
		players = append(players, player)
	}
	player := content.Player{Playerid: "99999", Teamid: 1, Nickname: "99999"}
	players = append(players, player)

	packet := content.S_MatchingDone{SessionId: 100, Players: players}

	GetNetworkCore().SendPacket(tc.TestServerConnections[0], packet, "MatchingDone")
}

func (tc *TestController) TestersLogin() {
	for i := 0; i < tc.TestPeopleCount-1; i++ {
		iString := fmt.Sprintf("%d", i)
		packet := content.C_EnterGame{SessionId: 100, Playerid: iString, Nickname: iString, GameMode: 50, Type: 0, Phonenum: iString}
		GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, "EnterGame")
	}
}

func (tc *TestController) TesterAttack() {
	go func() {
		for {
			for i := 0; i < tc.TestPeopleCount-1; i++ {
				iString := fmt.Sprintf("%d", i)
				packet := content.C_ThrowAttack{Playerid: iString, Roomid: 1, Type: 1, CCtype: 9999, Charactertype: 1, Soundtype: 1, Timer: 1, Startpos: helper.Vector3{999, 999, 999},
					Rotation: helper.Vector3{999, 999, 999}, Velocity: helper.Vector3{999, 999, 999}, IsSkill: true}
				GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, "ThrowAttack")
			}
			time.Sleep(time.Second * 1)
		}
	}()
}

func (tc *TestController) TesterDamage() {
	go func() {
		for {
			for i := 0; i < tc.TestPeopleCount-1; i++ {
				iString := fmt.Sprintf("%d", i)
				packet := content.C_GiveDamages{Playerid: iString, Targetids: []string{iString}, Damage: 1}
				GetNetworkCore().SendPacket(tc.TestServerConnections[i], packet, "GiveDamages")
			}
			time.Sleep(time.Second * 1)
		}
	}()
}
