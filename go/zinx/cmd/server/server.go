package main

import (
	"fmt"
	"zinx/pkg/ziface"
	"zinx/pkg/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

type HelloZinxRouter struct {
	znet.BaseRouter
}

// func (pr *PingRouter) PreHandle(request ziface.IRequest) {
// 	fmt.Println("Call Router Prehandle")
// 	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before ping...\n"))
// 	if err != nil {
// 		fmt.Println("call back ping error", err)
// 	}
// }
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")

	// _, err := request.GetConnection().GetTCPConnection().Write([]byte("Ping...Ping...Ping\n"))
	// if err != nil {
	// 	fmt.Println("call back ping error", err)
	// }
	fmt.Println("recv from client : msgId =", request.GetMsgID(), "msg = ", string(request.GetData()))
	err := request.GetConnection().SendBuffMsg(0, []byte("ping...ping...ping"))
	if err != nil {
		fmt.Println(err)
	}
}

// func (pr *PingRouter) PostHandle(request ziface.IRequest) {
// 	fmt.Println("Call Router Posthandle")
// 	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping...\n"))
// 	if err != nil {
// 		fmt.Println("call back ping error", err)
// 	}
// }

func (hz *HelloZinxRouter) Handle(request ziface.IRequest) {
	fmt.Println("call HelloZinxRouter handle")
	fmt.Println("recv from client : msgId = ", request.GetMsgID(), ", data=", string(request.GetData()))
	err := request.GetConnection().SendBuffMsg(1, []byte("Hello Zinx Router v0.6"))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("DoConnectionBegin is called ... ")
	err := conn.SendMsg(2, []byte("DoConnection BEGIN..."))
	if err != nil {
		fmt.Println(err)
	}
}

func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("DoConnectionLost is Called...")
}

func main() {
	s := znet.NewServer("[zinx V0.1]")

	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)

	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloZinxRouter{})
	s.Serve()
}
