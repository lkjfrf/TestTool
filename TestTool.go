package main

import (
	"encoding/json"
	"log"
	"net"
	"reflect"

	"./packet"
)

type Player struct {
	id       int
	password int
}

func main() {
	ConnArr := MakeConnections()
	log.Println(ConnArr)

	packet := packet.FSendPacket_PlayerLogin{}
	packet.PacketName = "FSendPacket_PlayerLogin"

	for i := 0; i < 100; i++ {
		packet.Id = string(i)
		packet.IsMan = true
		packet.LevelName = "Hyundai_Main"
		packet.UserName = string(i)
		BroadCastToAllServer(ConnArr, packet)
	}
}

func MakeConnections() []*net.TCPConn {
	Arr := make([]*net.TCPConn, 0)

	Arr = append(Arr, ConnectToServer("172.30.1.73:8000"))
	//Arr = append(Arr, network.ConnectToServer("13.125.193.33:8000"))
	//Arr = append(Arr, network.ConnectToServer("13.124.41.135:8000"))
	//Arr = append(Arr, network.ConnectToServer("3.37.118.146:8000"))

	return Arr
}

func ConnectToServer(serverAddr string) *net.TCPConn {
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

func SendPacket(c *net.TCPConn, p interface{}) {

	packet := reflect.ValueOf(p).Interface()

	e, err := json.Marshal(packet)
	if err != nil {
		log.Println(err)
	}

	done, err := c.Write(e)
	if err != nil {
		log.Println(err)
	}

	log.Println("Sent Byte -", done)
}

func BroadCastToAllServer(ConnArr []*net.TCPConn, p interface{}) {

	for _, c := range ConnArr {
		SendPacket(c, p)
	}
}
