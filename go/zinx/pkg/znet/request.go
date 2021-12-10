package znet

import (
	"zinx/pkg/ziface"
)

type Request struct {
	conn ziface.IConnection
	//data []byte
	msg ziface.IMessage
}

func (r *Request) GetConnection() ziface.IConnection {
	return r.conn
}

func (r *Request) GetData() []byte {
	return r.msg.GetData()
}

func (r *Request) GetMsgID() uint32 {
	return r.msg.GetMsgId()
}
