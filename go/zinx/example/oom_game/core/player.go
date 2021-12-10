package core

import (
	"fmt"
	"math/rand"
	"sync"
	"zinx/example/oom_game/pb"
	"zinx/pkg/ziface"

	"google.golang.org/protobuf/proto"
)

type Player struct {
	Pid  int32
	Conn ziface.IConnection
	Y    float32
	X    float32
	Z    float32
	V    float32
}

var PidGen int32 = 1
var IdLock sync.Mutex

func NewPlayer(conn ziface.IConnection) *Player {
	IdLock.Lock()
	id := PidGen
	PidGen++
	IdLock.Unlock()

	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)),
		Y:    0,
		Z:    float32(134 + rand.Intn(17)),
		V:    0,
	}
	return p
}

func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	fmt.Printf("before marshal data = %+v\n", data)
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("marshal msg err:", err)
		return
	}
	fmt.Printf("after marshal data = %+v\n", msg)
	if p.Conn == nil {
		fmt.Println("connection in player is nil")
		return
	}
	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("player sendmsg error!")
		return
	}
	return
}

func (p *Player) SyncPid() {
	data := &pb.SyncPid{
		Pid: p.Pid,
	}

	p.SendMsg(1, data)
}

func (p *Player) BroadCaseStartPositionew() {
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //TP2 代表广播坐标
		Data: &pb.BroadCast_P{
			&pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	p.SendMsg(200, msg)
}

func (p *Player) Talk(content string) {
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  1,
		Data: &pb.BroadCast_Content{
			Content: content,
		},
	}

	players := WorldMgrObj.GetAllPlayers()
	for _, player := range players {
		player.SendMsg(200, msg)
	}
}

func (p *Player) SyncSurrounding() {
	pids := WorldMgrObj.AoiMgr.GetPIDsByPos(p.X, p.Z)
	players := make([]*Player, 0, len(pids))

	for _, pid := range pids {
		players = append(players, WorldMgrObj.GetPlayerByPid(int32(pid)))
	}

	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  2, //TP2 代表广播坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	for _, player := range players {
		player.SendMsg(200, msg)
	}

	playersData := make([]*pb.Player, 0, len(players))
	for _, player := range players {
		p := &pb.Player{
			Pid: player.Pid,
			P: &pb.Position{
				X: player.X,
				Y: player.Y,
				Z: player.Z,
				V: player.V,
			},
		}
		playersData = append(playersData, p)
	}

	SyncPlayersMsg := &pb.SyncPlayers{
		Ps: playersData[:],
	}

	p.SendMsg(202, SyncPlayersMsg)
}

func (p *Player) UpdatePos(x float32, y float32, z float32, v float32) {
	//更新玩家的位置信息
	p.X = x
	p.Y = y
	p.Z = z
	p.V = v

	//组装protobuf协议，发送位置给周围玩家
	msg := &pb.BroadCast{
		Pid: p.Pid,
		Tp:  4, //4 - 移动之后的坐标信息
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	//获取当前玩家周边全部玩家
	players := p.GetSurroundingPlayers()
	//向周边的每个玩家发送MsgID:200消息，移动位置更新消息
	for _, player := range players {
		player.SendMsg(200, msg)
	}
}

//获得当前玩家的AOI周边玩家信息
func (p *Player) GetSurroundingPlayers() []*Player {
	//得到当前AOI区域的所有pid
	pids := WorldMgrObj.AoiMgr.GetPIDsByPos(p.X, p.Z)

	//将所有pid对应的Player放到Player切片中
	players := make([]*Player, 0, len(pids))
	for _, pid := range pids {
		players = append(players, WorldMgrObj.GetPlayerByPid(int32(pid)))
	}

	return players
}

func (p *Player) LostConnection() {
	//1 获取周围AOI九宫格内的玩家
	players := p.GetSurroundingPlayers()

	//2 封装MsgID:201消息
	msg := &pb.SyncPid{
		Pid: p.Pid,
	}

	//3 向周围玩家发送消息
	for _, player := range players {
		player.SendMsg(201, msg)
	}

	//4 世界管理器将当前玩家从AOI中摘除
	WorldMgrObj.AoiMgr.RemoveFromGridByPos(int(p.Pid), p.X, p.Z)
	WorldMgrObj.RemovePlayerByPid(p.Pid)
}
