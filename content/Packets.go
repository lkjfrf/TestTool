package content

import (
	"encoding/json"

	"github.com/lkjfrf/TestTool/helper"
)

const (
	Error = iota
	DBSignin
	PlayerLogout
	ChannelEnter
	NearPlayerUpdate

	PlayerMove //5
	PlayerActionEvent
	OtherPlayerMove
	PlayerLogin
	OtherPlayerSpawnInfo

	OtherPlayerDestroyInfo //10
	OtherInfo
	Voice
	RoomUserList
	RoomListUpdate

	Permission //15
	KickFromRoom
	MicToggle
	NoticeWrite
	NoticeContent

	NoticeList //20
	NoticeDelete
	NoticeModify
	ChannelCreate
	ChannelDelete

	CalenderRequest //25
	ChannelWidgetInfo
	NormalChat
	PrivateChat
	NoticeChat

	Questions // 30
	Invite
	InviteUserList
	CostumeSet
	UpdateCostume

	OtherUpdateCostume // 35
	HeartBeat
	AllFriendList
	SearchAddFriendList
	SearchDeleteFriendList

	AddFriend // 40
	DeleteFriend
	RequestAddFriend
	RequestDeleteFriend
	SpawnAvatar

	SaveFile //45
	CancelQuestion
	ModifyIntroduce
	FileList
	AccpetQuestion

	QuestionList //50
	ESaveShareData
	EPlaySaveShareData
	EEnterAuditorium
	EUploadComplete

	EScreenDataControlling //55
	RecvFileStatus
	ETestPlayerLogin

	Max
)

func JsonStrToStruct[T any](jsonstr string) T {
	var data T
	json.Unmarshal([]byte(jsonstr), &data)
	return data
}

type S_ChannelEnter struct {
	Id          string
	ChannelNum  int32
	ChannelType int32 // 0: Auditorium, 1: Convention, 2: VirtualOffice, 3: VirtualGallery, 4: Plaza
}

type R_ChannelEnter struct {
	Status bool
}

type S_PlayerMove struct {
	Id          string
	Position    helper.Vector3
	Rotation    helper.Vector3
	MoveSpeed   float32
	RotateSpeed float32
}

type R_OtherPlayerMove struct {
	Id           string
	Destination  helper.Vector3
	DestRotation helper.Vector3
	MoveSpeed    float32
	RotateSpeed  float32
}

type S_DBSignin struct {
	Id string
	Pw string
}

type R_DBSignin struct {
	Status bool

	IsCostumed bool // (true 라면 이미 옷고른상태)
	Hair       int32
	Face       int32
	Top        int32
	Bottom     int32
	Shoes      int32
	Name       string
	Team       string
	Grade      string
	Email      string
	Introduce  string
	HairColor  helper.Vector3
	SkinColor  helper.Vector3
	IsMan      bool
	IsGm       bool
	Offline    []FriendTargetInfo
}

type FriendTargetInfo struct {
	Id    string
	Name  string
	Grade string
	Team  string
}

type S_CostumeSet struct {
	Id        string
	Hair      int32
	Face      int32
	Top       int32
	Bottom    int32
	Shoes     int32
	IsMan     bool
	HairColor helper.Vector3
	SkinColor helper.Vector3
}

type R_NearPlayerUpdate struct {
	SpawnList   []FRecvPacket_OtherPlayerSpawnInfo
	DestroyList []FRecvPacket_OtherPlayerDestroyInfo
}

type FRecvPacket_OtherPlayerSpawnInfo struct {
	Id        string
	Position  helper.Vector3
	Rotation  helper.Vector3
	Hair      int32
	Face      int32
	Top       int32
	Bottom    int32
	Shoes     int32
	HairColor helper.Vector3
	SkinColor helper.Vector3
	Name      string
	Team      string
	IsMan     bool
}

type FRecvPacket_OtherPlayerDestroyInfo struct {
	Id string
}

type SR_PlayerActionEvent struct {
	Id       string
	ActionId string
}

// Channel
type S_ChannelCreate struct {
	//ChannelData structs.Channel
	IsUpdate bool
}

type R_ChannelCreate struct {
	Status bool
}

type S_ChannelDelete struct {
	Id          string
	ChannelNum  int32
	Year        int32
	Month       int32
	ChannelType int32
	IsCalender  bool
}

type S_ChannelWidgetInfo struct {
	Id          string
	ChannelType int32 // 0: AuditoriumSpeacker, 1: Convention, 2: VirtualOffice, 3: VirtualGallery, 4: Plaza
}

type R_ChannelWidgetInfo struct {
	//ChannelData []structs.Channel
}

// Calender
type CalenderData struct {
	ChannelNum  int32 `gorm:"primaryKey"`
	RoomTitle   string
	Country     string
	ChannelType int32 // 0: AuditoriumSpeacker, 1: Convention, 2: VirtualOffice, 3: VirtualGallery, 4: Plaza
	Name        string
	Id          string
	UserCount   int32
	MaxUsers    int32
	Password    string
	Content     string

	StartTime string
	EndTime   string

	StartYear  int32
	StartMonth int32
	StartDay   int32

	EndYear  int32
	EndMonth int32
	EndDay   int32

	StartDate int32
}

type S_CalenderRequest struct {
	Id    string
	Year  int32
	Month int32
}
type R_CalenderRequest struct {
	Year         int32
	Month        int32
	CalenderData []CalenderData
}

// Costume

type SR_UpdateCostume struct {
	Id    string
	Type  int32 // 2: Top, 3: Bottom, 4: Shoes
	Index int32
}

// Chat
type S_NormalChat struct {
	Id      string
	Message string
}

type R_NormalChat struct {
	Name    string
	Team    string
	Grade   string
	Message string
}

type S_PrivateChat struct {
	SendId  string
	RecvId  string
	Message string
	Voice   bool
}

type R_PrivateChat struct {
	SendId  string
	RecvId  string
	Name    string
	Team    string
	Message string
	Voice   bool
}

type SR_NoticeChat struct {
	Message string
}

// Voice

type SR_Voice struct {
	Id          string
	VoiceData   []uint16
	Numchannels int32
	SampleRate  int32
	PCMSize     int32
	Volume      float32
}

type S_PlayerLogout struct {
	Id string
}

type S_HeartBeat struct {
	Id string
}

// Notice

type S_NoticeWrite struct {
	Id      string
	Title   string
	Content string
}

type S_NoticeList struct {
	Id         string
	PageNum    int32
	SearchType int32 // 0 = Title, 1 = Content, 2 = Writer
	SearchText string
}

type R_NoticeList struct {
	NoticeList []Notice
	TotalCount int32
}

type Notice struct {
	Index int32
	Title string
	Id    string
	Date  string
	View  int32
}

type S_NoticeContent struct {
	Id    string
	Index int32
}

type R_NoticeContent struct {
	Content string
}

type S_NoticeDelete struct {
	Id    string
	Index int32
}

type S_NoticeModify struct {
	Id      string
	Index   int32
	Title   string
	Content string
}

// Friend
type Friend struct {
	Id        string
	Name      string
	Team      string
	Grade     string
	Email     string
	Introduce string
	Online    bool
}

type S_AllFriendList struct {
	Id string
}

type R_AllFriendList struct {
	AllFriendList []Friend
	FriendRequest []Friend
}

type S_SearchFriendList struct { // S_SearchAddFriendList, S_SearchDeleteFriendList
	Id     string
	Search string
}

type S_IdOnly struct {
	Id string
}

type R_SearchAddFriendList struct {
	SearchAddFriendList []Friend
}

type R_SearchDeleteFriendList struct {
	SearchDeleteFriendList []Friend
}

type S_RequestAddFriend struct { // 리스트에 친구요청
	Id       string
	TargetId string
}

type R_RequestAddFriend struct { // 친추 받았을떄 친추보낸사람의 정보
	Id    string
	Name  string
	Team  string
	Grade string
}

type R_AddFriend struct { // 친추 받은 사람이 수락했냐 거절했냐고 친추보낸 사람이 받음
	Status bool
	Name   string
	Team   string
	Grade  string
}

type S_AddFriend struct { // 친추대상의 수락여부 서버에 보냄
	Id       string
	Status   bool
	TargetId string
}

type S_DeleteFriend struct { // 친구삭제
	Id       string
	TargetId string
}

// Room
type R_RoomListUpdate struct {
	ChangedId   string // RoomList에 추가 되거나 삭제 or 마이크 변경 인원ID
	IsIn        int32  // 변경 없으면 0, 들어오면 1, 나가면 2
	IsMicToggle bool   // 기본값 false
	UserData    RoomUserData
}

type RoomUserData struct {
	Id         string
	Team       string
	Grade      string
	Name       string
	IsMicOn    bool
	Permission int32 // 0 유저 1 발표자 2 매니저 3 관리자 4 질문자
}

type S_RoomUserList struct {
	Id         string
	ChannelNum int32
}
type R_RoomUserList struct {
	UserList []RoomUserData
}

type S_Permission struct {
	TargetId   string
	Permission int32 // 0 유저 1 발표자 2 매니저 3 관리자 4 질문자
}

type R_Permission struct {
	Permission int32 // 0 유저 1 발표자 2 매니저 3 관리자 4 질문자
}

type S_OtherInfo struct {
	Id       string
	TargetId string
}

type R_OtherInfo struct {
	Name      string
	Team      string
	Grade     string
	Email     string
	Introduce string
}

type S_KickFromRoom struct {
	Id         string
	ChannelNum int32
}
type R_KickFromRoom struct {
	KickType int32 // 0 이면 kick, 1 이면 방폭파
}

type S_MicToggle struct {
	TargetId string
	IsMicOn  bool
}

type R_MicToggle struct {
	IsMicOn bool
}

type S_InviteUserList struct {
	Id     string
	Search string
}

type R_InviteUserList struct {
	UserList []RoomUserData
}

type S_Invite struct {
	TargetId   string
	Name       string
	ChannelNum int32
}

type R_Invite struct {
	Name        string
	ChannelNum  int32
	ChannelType int32
}

type R_File struct {
	FileName string
	Data     []byte
}

type S_SaveFile struct {
	FileName   string
	IsLastFile bool
}

type R_SaveFile struct {
	IsFinished bool
	Size       int32
	Data       []byte
}

type Question struct {
	Id       string
	Index    int32
	Name     string
	Team     string
	Grade    string
	Question string
}

type S_Question struct { // 새로운 Question 등록
	SendQuestion Question
	ChannelNum   int32
}

type R_Question struct {
	RecvQuestion Question
}

type S_QuestionList struct { // List 요청
	Id         string
	ChannelNum int32
}

type R_QuestionList struct { // QuestionList Send하면 리스트줌
	QuestionList      []Question
	Comp_QuestionList []Question
	Questioning       Question
}

type S_CancelQuestion struct { // 질문 취소함
	Index      int32
	ChannelNum int32
}

type R_CancelQuestion struct { // 질문취소 받음
	Index         int32
	IsQuestioning bool // 질문자를 지운건지, 질문대기 중인걸 지운건지
}

type S_AcceptQuestion struct { // 수락한 질문 보냄
	Id         string
	Index      int32
	ChannelNum int32
}

type R_AcceptQuestion struct { // 질문중으로 표시할 질문 받음
	QuestionData Question
}

type S_ModifyIntroduce struct {
	Id        string
	Introduce string
}

type R_FileList struct {
	FileList []string
	IsEnter  bool
}

type S_FileList struct {
	Id         string
	ChannelNum int32
	IsEnter    bool
}

type R_Error struct {
	Status int32 //  1 : 발표자가 4명이상일때
}

// 화면송출
type SaveShareData struct {
	LeftFileName   string
	CenterFileName string
	RightFileName  string
}

type S_SaveShareData struct {
	ChannelNum       int32
	ArrSaveShareData []SaveShareData
}

type R_SaveShareData struct {
	ArrSaveShareData []SaveShareData
}

type S_PlaySaveShareData struct {
	ChannelNum      int32
	CurPlaySaveData PlaySaveShareData
}

type R_PlaySaveShareData struct {
	CurPlaySaveData PlaySaveShareData
}

type S_EnterAuditorium struct {
	Id string
}

type R_UploadComplete struct {
	Data bool
}

type PlaySaveShareData struct {
	PlayIndex  int32
	ArrPdfPage []int32 // left, center, right 의 페이지
	IsPDF      []bool  // true = PDF, false = MP4
	Seconds    []int32
}

type SR_ScreenDataControlling struct {
	ChannelNum     int32
	PlayIndex      int32 // 모니터 설정번호
	Position       int32 // 0,1,2 왼,가,오
	IsPDF          bool
	ControlCommand int32
	Seconds        int32 //controlcommand flase 일떄 시간 저장후 true일떄 시간 주기
}

type R_RecvFileStatus struct {
	TotalSize int32
	IsEnd     bool
}
