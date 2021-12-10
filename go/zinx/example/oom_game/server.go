package main

import (
	"fmt"
	"zinx/example/oom_game/api"
	"zinx/example/oom_game/core"
	"zinx/pkg/ziface"
	"zinx/pkg/znet"
)

func OnConnectionAdd(conn ziface.IConnection) {
	player := core.NewPlayer(conn)
	player.SyncPid()
	player.BroadCaseStartPositionew()

	core.WorldMgrObj.AddPlayer(player)
	conn.SetProperty("pid", player.Pid)

	player.SyncSurrounding()
	fmt.Println("===> Player idId = ", player.Pid, " arrived ====")
}

func OnConnectionLost(conn ziface.IConnection) {
	pid, _ := conn.GetProperty("pid")
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	if pid != nil {
		player.LostConnection()
	}
	fmt.Println("====> Player ", pid, " left =====")
}

func main() {
	s := znet.NewServer("zinx game")
	s.SetOnConnStart(OnConnectionAdd)
	s.SetOnConnStop(OnConnectionLost)

	s.AddRouter(2, &api.WorldChatApi{})
	s.AddRouter(3, &api.MoveApi{})
	s.Serve()
}
