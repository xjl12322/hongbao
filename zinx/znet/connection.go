package znet

import (
	"fmt"
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

	//告知当前链接已经退出停止 channel
	ExitChan chan bool
	//该链接处理的方法Router
	Router ziface.IRouter
}
//初始化链接方法
func NewConnection(conn *net.TCPConn,connID uint32,router ziface.IRouter) ziface.IConnection  {
	c := &Connecttion{
		Conn:conn,
		ConnID:connID,
		Router:router,
		isClosed:false,
		ExitChan:make(chan bool),
	}

	return c

}
func (c *Connecttion)StartReader()  {
	fmt.Println("Reader Goruoutine is runing....")
	defer fmt.Printf("connID=%d Reader is exit, remote add is ",c.ConnID,c.RemoteAddr().String())
	defer c.Stop()

	for {
		//读取客户端的数据到buf中 最大512字节
		buf := make([]byte,512)
		_,err := c.Conn.Read(buf)
		if err != nil{
			fmt.Println("recv buf err ",err)
			continue
		}
		////调用当前链接回显的handleAPIv.2
		//if err := c.handleAPI(c.Conn,buf,cnt); err != nil{
		//	fmt.Printf("ConnID = %d handle is error = %s",c.ConnID,err)
		//	break
		//}

		//得到当前conn数据的Request请求数v.3
		req := Request {
			conn:c,
			data:buf,
		}
		//执行注册的路由方法
		go func(request ziface.IRequest) {
			c.Router.PreHandle(request)
			c.Router.Handle(request)
			c.Router.PostHandle(request)
		}(&req)
		//从路由中，找到注册绑定的Conn对应的router调用


	}

}
//启动链接让当前链接准备开始工作
func (c *Connecttion)Start()  {
	fmt.Println("Conn Start() ... ConnID=",c.ConnID)
	//启动当前链接的读取数据业务
	go c.StartReader()

	//TODO 启动当前链接写数据的业务
}

//停止链接 结束当前的链接工作
func (c *Connecttion)Stop() {
	fmt.Printf("Conn Stop.. ConnID= %d",c.ConnID)
	//如果当前链接已经关闭
	if c.isClosed == true{
		return
	}

	c.isClosed = true
	//关闭链接
	c.Conn.Close()

	close(c.ExitChan)


}
//获取当前链接绑定的socket conn
func (c *Connecttion)GetTCPConnection()*net.TCPConn{
	return c.Conn

}
//获取当前链接模块的链接ID
func (c *Connecttion)GetConnID() uint32{
	return c.ConnID

}
//获取客户端的tcp状态 IP  port
func (c *Connecttion)RemoteAddr()net.Addr{
	return c.Conn.RemoteAddr()
}
//发送数据，将数据发送给远程客户端
func (c *Connecttion)Send(data []byte)error{
	return nil

}



