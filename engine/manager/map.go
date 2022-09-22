package manager

import (
	"BurninProject/engine/common"
	"BurninProject/engine/gameMap"
	"github.com/xiaonanln/goworld/engine/gwlog"
)

var (
	mapManager = newMapMgr()
)

//MapMgr 维护地图
type _MapMgr struct {
	gameMaps map[common.EntityID]*gameMap.GameMap
}

func newMapMgr() *_MapMgr {
	return &_MapMgr{
		gameMaps: make(map[common.EntityID]*gameMap.GameMap),
	}
}

func (mapmgr *_MapMgr) putSpace(space *gameMap.GameMap) {
	mapmgr.gameMaps[space.MapUId] = space
}

func (mapmgr *_MapMgr) delSpace(id common.EntityID) {
	delete(mapmgr.gameMaps, id)
}

func (mapmgr *_MapMgr) getSpace(id common.EntityID) *gameMap.GameMap {
	gwlog.Infof("spmgr.gameMaps %s", mapmgr.gameMaps)
	return mapmgr.gameMaps[id]
}
