package ziface

//定义服务器接口
type IServer interface {
	//启动服务器方法
	Start()
	//停止服务器方法
	Stop()
	//开启业务服务方法
	Serve()
	//
	AddRouter(msgId uint32, router IRouter)

	GetConnMgr() IConnManager

	SetOnConnStart(func(IConnection))

	SetOnConnStop(func(IConnection))

	CallOnConnStart(conn IConnection)
	CallOnConnStop(conn IConnection)
}
