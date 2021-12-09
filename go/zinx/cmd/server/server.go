package main

import (
	"zinx/pkg/znet"
)

func main() {
	s := znet.NewServer("[zinx V0.1]")
	s.Serve()
}
