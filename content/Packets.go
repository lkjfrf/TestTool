package content

import (
	"encoding/json"

	"github.com/lkjfrf/TestTool/helper"
)

const (
	Error = iota

	Max
)

func JsonStrToStruct[T any](jsonstr string) T {
	var data T
	json.Unmarshal([]byte(jsonstr), &data)
	return data
}

type S_MatchingDone struct {
	SessionId int32
	Players   []Player
}

type Player struct {
	Playerid string
	Nickname string
	Teamid   int32
}

type C_EnterGame struct { // 재연결 때 이걸보내줌
	Playerid string
	Nickname string
	GameMode int32

	Type int32

	//GStar
	Phonenum string

	SessionId int32
}

type C_ThrowAttack struct {
	Playerid string
	Roomid   int32
	Type     int32
	CCtype   int32

	Charactertype int32
	Soundtype     int32
	Timer         float32

	Startpos helper.Vector3
	Rotation helper.Vector3
	Velocity helper.Vector3

	IsSkill bool
}

type C_GiveDamages struct {
	Playerid  string
	Targetids []string
	Damage    int32
}
