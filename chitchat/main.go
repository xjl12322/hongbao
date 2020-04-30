package main

import (
	"github.com/gin-gonic/gin"
	"hongbao/chitchat/handler"
	"hongbao/chitchat/middlewares"
	"log"
)

func main()  {
	startWebServer("8080")
}


// 通过指定端口启动 Web 服务器
func startWebServer(port string)  {
	r := gin.Default()


	r.Static("/static/","./public")
	//b.Application.Static("/public/","./web/public")
	//b.Application.SetFuncMap(template.FuncMap{
	//	"FromUnixtime":comm.FromUnixtime,
	//	"FromUnixtimeShort":comm.FromUnixtimeShort,
	//})
	r.LoadHTMLGlob("./views/*")
	r.GET("/index",middlewares.Indexrequited(),handler.IndexHandle)   //首页
	r.GET("/login",middlewares.Loginrequits(),handler.LoginHandle)//登录页
	r.GET("/signup",handler.SignupHandle)//注册页
	r.GET("/login_outs",handler.LoginOutHandle)//登出

	r.POST("/login_account",handler.LoginAccountHandle)//登录认证
	r.POST("/signup_account",handler.SignupAccountHandle)//注册认证
	r.GET("/thread/read",handler.ThreadReadHandler)//
	r.GET("/thread/new",middlewares.Loginrequits(),handler.ThreadNewHandler)// 创建主题页面
	r.POST("/thread/create",middlewares.Loginrequits(),handler.CreateThreadNewHandler)// 创建主题



	log.Println("Starting HTTP service at " + port)
	err := r.Run(":" + port) // 启动协程监听请求
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + port)
		log.Println("Error: " + err.Error())
	}
}





