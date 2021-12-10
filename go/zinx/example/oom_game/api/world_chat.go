package api

import (
	"fmt"
	"zinx/example/oom_game/core"
	"zinx/example/oom_game/pb"
	"zinx/pkg/ziface"
	"zinx/pkg/znet"

	"google.golang.org/protobuf/proto"
)

type WorldChatApi struct {
	znet.BaseRouter
}

func (wc *WorldChatApi) Handle(request ziface.IRequest) {
	msg := &pb.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("talk unmarshal error ", err)
		return
	}

	pid, err := request.GetConnection().GetProperty("pid")
	if err != nil {
		fmt.Println("GetProperty pid error", err)
		request.GetConnection().Stop()
		return
	}

	player := core.WorldMgrObj.GetPlayerByPid(pid.(int32))
	player.Talk(msg.Content)
}
