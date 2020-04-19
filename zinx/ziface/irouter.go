package ziface


/*
路由的抽象接口，
路由里的度数据都是IRequest
 */

type IRouter interface {
	//在处理conn业务之前的钩子方法Hook
	PreHandle(request IRequest)
	//在处理conn业务的主方法hook
	Handle(request IRequest)
	//在处理conn业务之后的钩子方法Hook
	PostHandle(request IRequest)

}





