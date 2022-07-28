package core

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"sync"
)

type NetworkCore struct {
	wg sync.WaitGroup
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

func (nc *NetworkCore) Connect(IP string, len int) []*net.TCPConn {
	Arr := []*net.TCPConn{}

	for i := 0; i < len; i++ {
		Arr = append(Arr, nc.ConnectToServer(IP, string(i)))
	}

	return Arr
}

func (nc *NetworkCore) ConnectToServer(serverAddr string, id string) *net.TCPConn {
	tcpAddr, err := net.ResolveTCPAddr("tcp", serverAddr)
	if err != nil {
		log.Println(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Println(err)
	}

	go nc.Recv(conn, id)
	return conn
}

func (nc *NetworkCore) SendPacket(c *net.TCPConn, recvpkt any, pkttype uint16) {
	sendBuffer := MakeSendBuffer(pkttype, recvpkt)

	if c != nil {
		sent, err := c.Write(sendBuffer)
		if err != nil {
			log.Println("SendPacket ERROR :", err)
		} else {
			if sent != len(sendBuffer) {
				log.Println("[Sent diffrent size] : SENT =", sent, "BufferSize =", len(sendBuffer))
			}
			log.Println("c:", c, "-", pkttype)
		}

	}
}

func (nc *NetworkCore) Recv(conn *net.TCPConn, id string) {
	header := make([]byte, 4)
	for {
		n, err := conn.Read(header)
		if err != nil {
			break
		}

		if n > 0 {
			pktsize, pktid := nc.ParseHeader(header)
			datasize := pktsize - 4
			if datasize > 0 {
				recv := make([]byte, datasize)
				_n, _ := conn.Read(recv)
				fmt.Println(_n, pktid)
				// if pktid == 12 {
				// 	recvpkt := JsonStrToStruct[content.SR_Voice](string(recv[:_n]))
				// 	recvpkt.Id = "tester" + id

				// 	sendBuffer := MakeSendBuffer(12, recvpkt)
				// 	conn.Write(sendBuffer)
				// }
			}

		}
	}
}

func JsonStrToStruct[T any](jsonstr string) T {
	var data T
	json.Unmarshal([]byte(jsonstr), &data)
	return data
}

func MakeSendBuffer[T any](pktid uint16, data T) []byte {
	sendData, err := json.Marshal(&data)
	if err != nil {
		log.Println("MakeSendBuffer : Marshal Error", err)
	}
	sendBuffer := make([]byte, 4)

	pktsize := len(sendData) + 4

	binary.LittleEndian.PutUint16(sendBuffer, uint16(pktsize))
	binary.LittleEndian.PutUint16(sendBuffer[2:], pktid)

	sendBuffer = append(sendBuffer, sendData...)

	return sendBuffer
}

func (nc *NetworkCore) ParseHeader(header []byte) (int, int) {
	pktsize := binary.LittleEndian.Uint16(header[:2])
	pktid := binary.LittleEndian.Uint16(header[2:4])

	return int(pktsize), int(pktid)
}
