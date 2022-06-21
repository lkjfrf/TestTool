package content

import (
	"github.com/lkjfrf/TestTool/helper"
)

type FSendPacket_TestPlayerLogin struct {
	PacketName string
	Id         string
	IsMan      bool
	LevelName  string
	UserName   string
}

func New_FSendPacket_TestPlayerLogin() FSendPacket_TestPlayerLogin {
	result := FSendPacket_TestPlayerLogin{}
	result.PacketName = "FSendPacket_TestPlayerLogin"
	return result
}

type FSendPacket_PlayerMove struct {
	PacketName  string
	Id          string
	Position    helper.Vector3
	Rotation    helper.Vector3
	MoveSpeed   float32
	RotateSpeed float32
}

func New_FSendPacket_PlayerMove() FSendPacket_PlayerMove {
	result := FSendPacket_PlayerMove{}
	result.PacketName = "FSendPacket_PlayerMove"
	return result
}

type FSendPacket_HeartBeat struct {
	PacketName string
	Id         string
}

func New_FSendPacket_HeartBeat() FSendPacket_HeartBeat {
	result := FSendPacket_HeartBeat{}
	result.PacketName = "FSendPacket_HeartBeat"
	return result
}

type FSendPacket_NormalChat struct {
	PacketName string
	Id         string
	UserName   string
	Message    string
}

func New_FSendPacket_NormalChat() FSendPacket_NormalChat {
	result := FSendPacket_NormalChat{}
	result.PacketName = "FSendPacket_NormalChat"
	return result
}
