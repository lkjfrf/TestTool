package core

import (
	"encoding/json"
	"log"
	"net"
	"reflect"
	"sync"
)

type NetworkCore struct {
}

var instance *NetworkCore
var once sync.Once

func GetNetworkCore() *NetworkCore {
	once.Do(func() {
		instance = &NetworkCore{}
	})
	return instance
}

func (nc *NetworkCore) Init() {
	log.Println("INIT_NetworkCore")
}

func (nc *NetworkCore) Connect(IPs []string) []*net.TCPConn {
	Arr := make([]*net.TCPConn, 0)

	for i, _ := range Arr {
		Arr = append(Arr, nc.ConnectToServer(IPs[i]))
	}

	return Arr
}

func (nc *NetworkCore) ConnectToServer(serverAddr string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Println(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err)
	}

	return conn
}

func (nc *NetworkCore) SendPacket(ConnArr []*net.TCPConn, p interface{}) {

	packet := reflect.ValueOf(p).Interface()

	e, err := json.Marshal(packet)
	if err != nil {
		log.Println(err)
	}

	for _, c := range ConnArr {
		_, err := c.Write(e)
		if err != nil {
			log.Println("Packet Send Err : ", err)
		}
	}
}
