package gameConf

// csv配置表
var G_RoomList map[string]interface{} // 房间列表
var RoomListData map[string]*RoomList // 房间列表

func init() {
	G_RoomList = make(map[string]interface{})
	RoomListData = make(map[string]*RoomList)
}

//------------------------------------------------------------------------------

// 房间列表
type RoomList struct {
	RoomID    string // 房间列表
	NeedPiece string // 进场花费的金币
	NeedLev   string // 进场需要的等级
	Desc      string // 描述
	SysPiece  string // 系统抽水
	WinReward string // 每局获得
	IsTop     string // 是否置顶
	TypeICON  string // 活动的ICON
}
