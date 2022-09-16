package csv

import (
	"BurninProject/engine/consts"
	"BurninProject/engine/gameConf"
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadCsv_ConfigFile_RoomListST_Fun() bool {
	// 获取数据，按照文件
	fileName := "roomlist.csv"
	fileName = consts.GameConfigPath + fileName
	cntb, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic("读取配置文件出错!")
		return false
	}
	// 读取文件数据
	r2 := csv.NewReader(strings.NewReader(string(cntb)))
	ss, _ := r2.ReadAll()
	sz := len(ss)
	roomidtemp := ""
	// 循环取数据
	for i := 1; i < sz; i++ {
		Infotmp := new(gameConf.RoomList)
		Infotmp.RoomID = ss[i][0]
		Infotmp.NeedPiece = ss[i][1]
		Infotmp.NeedLev = ss[i][2]
		Infotmp.Desc = ss[i][3]
		Infotmp.SysPiece = ss[i][4]
		Infotmp.WinReward = ss[i][5]
		Infotmp.IsTop = ss[i][6]
		Infotmp.TypeICON = ss[i][7]

		s := string([]byte(Infotmp.RoomID)[:5])
		if len(roomidtemp) == 0 {
			roomidtemp = s
			gameConf.RoomListData[Infotmp.RoomID] = Infotmp

		} else {
			if roomidtemp == s {
				gameConf.RoomListData[Infotmp.RoomID] = Infotmp
				fmt.Println("+++++++++", gameConf.RoomListData)
				// 仅仅有一个游戏的时候
				if i == sz-1 {
					gameConf.G_RoomList[roomidtemp] = gameConf.RoomListData
				}
			} else {
				// 保存数据
				gameConf.G_RoomList[roomidtemp] = gameConf.RoomListData
				roomidtemp = s
				gameConf.RoomListData = make(map[string]*gameConf.RoomList)
			}
		}
	}
	return true
}
