package gameMap

import (
	"BurninProject/engine/common"
	"BurninProject/engine/consts"
	"BurninProject/engine/player"
	"time"
)

var (
	nilGamMap *GameMap
)

type GameMap struct {
	players map[*player.Player]struct{}
	MapUId  common.EntityID
}

func NewGameMap(entityID common.EntityID) *GameMap {
	if entityID == "" {
		entityID = common.GenEntityID()
	}
	p := &GameMap{
		players: make(map[*player.Player]struct{}, consts.GameMapPlayerMaxNum),
		MapUId:  entityID,
	}
	//p.HandlerRegister()
	return p
}

func (gMap *GameMap) Run() {
	go gMap.Ticker()
}

//地图定时器
func (gMap *GameMap) Ticker() {
	myTimer := time.NewTimer(time.Second) // 启动定时器
	//ForEnd:
	for {
		select {
		case <-myTimer.C:
			//break ForEnd
			myTimer.Reset(time.Second) // 每次使用完后需要人为重置下
		}
	}
	// 不再使用了，结束它
	myTimer.Stop()
}
