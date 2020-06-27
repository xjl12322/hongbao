package main

import (
	"fmt"
	"github.com/gin-gonic/gin"

)

type middleWareHandler struct {
	r *gin.RouterGroup
}

//func NewMiddleWareHandler(r *gin.RouterGroup)http.Handler  {
//	m := middleWareHandler{}
//	m.r = r
//	return m
//}

//func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
//	// 检查身份的合法性
//	validateUserSession(r)
//	m.r.ServeHTTP(w,r)
//}

func RegisterHandlers() *gin.Engine  {
	router := gin.Default()

	//router:= httprouter.New()
	router.POST("/user" ,CreateUser)
	router.POST("/user/:user_name",validateUserSession(), Login)
	return router
}



func main()  {
	fmt.Println("sta")
	r:= RegisterHandlers()
	r.Run(":8080")
	//mh := NewMiddleWareHandler(r.RouterGroup)  //中间件处理
	//http.ListenAndServe(
	//	":8080",mh)


}




