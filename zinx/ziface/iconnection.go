package ziface

import "net"

//定义一个服务器接口
type IConnection interface {
	//启动链接让当前的链接准备开始工作
	Start()
	//停止链接 结束当前的链接工作
	Stop()
	//获取当前链接绑定的socket conn
	GetTCPConnection()*net.TCPConn
	//获取当前链接模块的链接ID
	GetConnID() uint32
	//获取客户端的tcp状态 IP  port
	RemoteAddr()net.Addr
	//发送数据，将数据发送给远程客户端
	Send(data []byte)error
}

//定义一个处理链接业务的方法
type HandleFunc func(*net.TCPConn,[]byte,int) error