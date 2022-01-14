package packet

import (
	"./transform"
)

type FSendPacket_PlayerLogin struct {
	PacketName string
	Id         string
	IsMan      bool
	LevelName  string
	UserName   string
}

type FSendPacket_PlayerMove struct {
	PacketName   string
	Id           string
	Destination  transform.Vector3
	DestRotation transform.Vector3
	MoveSpeed    float32
	RotateSpeed  float32
}
