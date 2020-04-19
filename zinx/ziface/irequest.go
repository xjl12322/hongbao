package ziface


/*
实际上是把客户端链接的信息 和请求的数据包装到一个request
 */

type IRequest interface {
	//得到当前的链接
	GetConnection() IConnection
	//得到请求的消息数据
	GetData() []byte
}



