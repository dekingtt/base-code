package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"zinx/pkg/ziface"
)

func init() {
	GlobalObject = &GlobalObj{
		Name:             "ZinxServerApp",
		Version:          "V0.4",
		TcpPort:          7777,
		Host:             "0.0.0.0",
		MaxConn:          12000,
		MaxPacketSize:    4096,
		ConfFilePaht:     "conf/zinx.json",
		WorkerPoolSize:   10,
		MaxWorkerTaskLen: 1024,
		MaxMsgChanLen:    512,
	}
	GlobalObject.Reload()
}

type GlobalObj struct {
	// server
	TcpServer ziface.IServer
	Host      string
	TcpPort   int
	Name      string

	// zinx
	Version          string
	MaxPacketSize    uint32
	MaxConn          int
	WorkerPoolSize   uint32
	MaxWorkerTaskLen uint32
	MaxMsgChanLen    uint32

	ConfFilePaht string
}

var GlobalObject *GlobalObj

func (g *GlobalObj) Reload() {
	if _, err := os.Stat("conf/zinx.json"); errors.Is(err, os.ErrNotExist) {
		fmt.Println("conf file not exist")
		return
	}
	data, err := ioutil.ReadFile("conf/zinx.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &GlobalObject)
	if err != nil {
		panic(err)
	}
}
