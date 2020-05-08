package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-delve/delve/pkg/config"
	"hongbao/chitchat/handler"
	"hongbao/chitchat/middlewares"
	"hongbao/chitchat/models"
	"html/template"
	"log"
	. "hongbao/chitchat/config"
	"net/http"

)


func main()  {

	startWebServer("8080")
}


// 通过指定端口启动 Web 服务器
func startWebServer(port string)  {
	r := gin.Default()
	// 在入口位置初始化全局配置
	config := LoadConfig()
	//assets := http.FileServer(http.Dir(config.App.Static))
	//r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", assets))
	r.Static("/static/","./public")
	r.Static("/public/","./web/public")
	r.SetFuncMap(template.FuncMap{
		"CreatedAtDate":models.CreatedAtDate,
	})
	r.LoadHTMLGlob("./views/*")
	r.GET("/index",middlewares.Indexrequited(),handler.IndexHandle)   //首页
	r.GET("/login",middlewares.Loginrequits(),handler.LoginHandle)//登录页
	r.GET("/signup",handler.SignupHandle)//注册页
	r.GET("/login_outs",handler.LoginOutHandle)//登出

	r.POST("/login_account",handler.LoginAccountHandle)//登录认证
	r.POST("/signup_account",handler.SignupAccountHandle)//注册认证
	r.GET("/thread/read",middlewares.Indexrequited(),handler.ThreadReadHandler)//通过ID渲染指定群组页面
	r.GET("/thread/new",middlewares.Loginrequits(),handler.ThreadNewHandler)// 创建主题页面
	r.POST("/thread/create",middlewares.Loginrequits(),handler.CreateThreadNewHandler)// 创建主题
	r.POST("/thread/post",middlewares.Loginrequits(),handler.CreateThreadPostHandler)// 创建主题


	log.Println("Starting HTTP service at " + config.App.Address)
	err := r.Run(config.App.Address) // 启动协程监听请求
	if err != nil {
		log.Println("An error occured starting HTTP listener at port " + config.App.Address)
		log.Println("Error: " + err.Error())
	}
}





