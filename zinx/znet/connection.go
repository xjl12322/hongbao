package znet

import (
	"hongbao/zinx/ziface"
	"net"
)

/*
链接模块
 */
type Connecttion struct {
	//当前链接的socket tcp套接字
	Conn *net.TCPConn
	//链接的ID
	ConnID uint32
	//当前的链接状态
	isClosed bool
	//当前链接诶所绑定的处理业务方法API
	handleAPI ziface.HandleFunc
	//告知当前链接已经退出停止 channel
	ExitChan chan bool
}
//初始化链接方法
func NewConnection(conn *net.TCPConn,connID uint32,callback_api ziface.HandleFunc)  {
	c := &Connecttion{
		Conn:conn,
		ConnID:connID,
		handleAPI:callback_api,
		isClosed:false,
		ExitChan:make(chan bool),
	}



}
