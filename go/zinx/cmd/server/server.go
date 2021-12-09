package main

import (
	"fmt"
	"zinx/pkg/ziface"
	"zinx/pkg/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (pr *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("Call Router Prehandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Before ping...\n"))
	if err != nil {
		fmt.Println("call back ping error", err)
	}
}
func (pr *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("Call PingRouter Handle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("Ping...Ping...Ping\n"))
	if err != nil {
		fmt.Println("call back ping error", err)
	}
}
func (pr *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("Call Router Posthandle")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("After ping...\n"))
	if err != nil {
		fmt.Println("call back ping error", err)
	}
}
func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
