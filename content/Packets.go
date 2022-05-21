package content

import (
	"github.com/lkjfrf/TestTool/helper"
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
	Destination  helper.Vector3
	DestRotation helper.Vector3
	MoveSpeed    float32
	RotateSpeed  float32
}

func New_FSendPacket_PlayerMove() FSendPacket_PlayerMove {
	result := FSendPacket_PlayerMove{}
	result.PacketName = "FSendPacket_PlayerMove"
	return result
}
