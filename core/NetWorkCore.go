package core

import (
	"encoding/json"
	"fmt"
	"io"
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
	len := len(IPs)
	Arr := make([]*net.TCPConn, len)

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
		log.Fatal("Parse Error")
		return
	}

	for _, c := range ConnArr {
		if c != nil {
			_, err := (*c).Write(e)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}

func (nc *NetworkCore) RecvPacket() {
	go func() {
		header := make([]byte, 4096)
		for {
			for i, t := range GetTestController().TestServerConnections {
				n, err := t.Read(header)
				if err != nil {
					if err == io.EOF {
						fmt.Println("EOF err : ", header)
					}
					break
				}

				if n > 0 {
					fmt.Println(header)
					GetTestController().TestServerConnections[i].Read(header)
					log.Println(header)
				}
			}
		}
	}()
}
