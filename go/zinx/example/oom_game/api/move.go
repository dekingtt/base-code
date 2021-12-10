package api

import (
	"fmt"
	"zinx/example/oom_game/core"
	"zinx/example/oom_game/pb"
	"zinx/pkg/ziface"
	"zinx/pkg/znet"

	"google.golang.org/protobuf/proto"
)

type MoveApi struct {
	znet.BaseRouter
}

func (mv *MoveApi) Handle(request ziface.IRequest) {
	msg := &pb.Position{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Move: position unmarshal error ", err)
		return
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("GetProperty pid error", err)
		request.GetConnection().Stop()
		return
	}

	fmt.Printf("user pid = %d , move(%f,%f,%f,%f)", pid, msg.X, msg.Y, msg.Z, msg.V)
	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))

	player.UpdatePos(msg.X, msg.Y, msg.Z, msg.V)
}
