package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandler() *httprouter.Router {
	router := httprouter.New()
	router.GET("/", homeHandler) //首页
	router.POST("/", homeHandler)

	router.GET("/userhome", userHomeHandler)//用户页
	router.POST("/userhome", userHomeHandler)
	//
	router.POST("/api", apiHandler)
	router.POST("/upload/:vid-id", proxyHandler) //域名的转换 web请求转发到streamserver 上传文件接口
	//123.0.0.1/statics 指向templates文件
	router.ServeFiles("/statics/*filepath", http.Dir("./templates"))
	return router
}

func main() {
	r := RegisterHandler()
	http.ListenAndServe(":8080", r)



}

