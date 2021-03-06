package znet

import (
	"fmt"
	"hongbao/zinx/ziface"
	"net"
)

//iserver 的接口实现，定义一个server的服务器模块
type Server struct {
	//服务器名称
	Name string
	//服务器绑定ip
	IPVersion string
	//服务监听ip
	IP string
	//服务监听端口
	Port int
	//当前的Server添加一个router，server注册的链接对应的处理业务
	Router ziface.IRouter
}

//定义当前客户端所绑定的handler API  （目前handler是写死的好，后期优化用户自定义handle方法）
//func CallBackToClient(conn *net.TCPConn,data []byte,cnt int) error{
//	//回显业务
//	fmt.Println("[Conn handlerl] CallbackToclient...")
//	if _,err := conn.Write(data[:cnt]);err != nil{
//		fmt.Println("write back buf err=",err)
//		return errors.New("CallBackToClient error")
//	}
//	return nil
//}

//路由功能：给当前的服务注册一个路由方法，供客户端的链接处理使用
func (s *Server)AddRouter(router ziface.IRouter)()  {
	s.Router = router
	fmt.Println("Add Router Succ!")

}


//初始化server
func NewServer(name string)ziface.IServer {
	s := &Server{
		Name:name,
		IPVersion:"tcp4",
		IP:"0.0.0.0",
		Port:8999,
	}
	return s
}


func (s *Server)Start()  {
	fmt.Printf("[start] Server Listenner at IP :%s prot:%d\n",s.IP,s.Port)
	go func() {
		// 1 获取一个tcp的 addr
		addr,err := net.ResolveTCPAddr(s.IPVersion,fmt.Sprintf("%s:%d",s.IP,s.Port))

		if err != nil{
			fmt.Println("resolve tcp adds error\n:",err)
			return
		}
		// 2 监听服务器的地址
		listenner,err := net.ListenTCP(s.IPVersion,addr)
		if err != nil{
			fmt.Printf("listen :%s err=%s\n",s.IPVersion,err)
			return
		}
		fmt.Printf("start zinx server succuss\n")

		//3 阻塞的等待客户端连接，处理客户端的连接业务（读写）
		for {
			//如果客户端连接过来，阻塞并返回
			conn,err := listenner.AcceptTCP()
			if err != nil{
				fmt.Printf("Accept err :%s\n",err)
				continue
			}
			var cid uint32
			cid = 0

			dealConn := NewConnection(conn,cid,s.Router)
			cid ++
			//启动当前业务处理
			go dealConn.Start()

		}


	}()


}

func (s *Server)Stop()  {
    //TODO  将一些服务器资源，状态或一些已经开辟的链接信息 进行回收
}

//运行服务器
func (s *Server)Serve()  {
	//启动
	s.Start()
	//TODO：在serve启动start可以在这做额外的扩展
	//阻塞状态
	select {

	}

}



