package core

import "sync"

type WorldManager struct {
	AoiMgr  *AOIManager
	Players map[int32]*Player
	pLock   sync.RWMutex
}

var WorldMgrObj *WorldManager

func init() {
	WorldMgrObj = &WorldManager{
		Players: make(map[int32]*Player),
		AoiMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
	}
}

func (wm *WorldManager) AddPlayer(player *Player) {
	wm.pLock.Lock()
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()

	wm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	wm.pLock.Lock()
	delete(wm.Players, pid)
	wm.pLock.Unlock()
}

func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	return wm.Players[pid]
}

func (wm *WorldManager) GetAllPlayers() []*Player {
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()
	players := make([]*Player, 0)

	for _, v := range wm.Players {
		players = append(players, v)
	}

	return players
}
