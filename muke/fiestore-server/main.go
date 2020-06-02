package main

import (
	"hongbao/muke/fiestore-server/handler"
	"fmt"
	"net/http"
)

func main()  {
	//文件的crud
	http.HandleFunc("/file/upload",handler.UploadHandler) //文件上传
	http.HandleFunc("/file/upload/suc",handler.UploadSucHandler)  //返回上传成功跳转信息
	http.HandleFunc("/file/meta",handler.GetFileMetaHandler) //获取文件信息
	http.HandleFunc("/file/download",handler.DownloadHandler)//客户端下载文件
	http.HandleFunc("/file/delete",handler.FileDeleteHandler)//修改文件名称
	http.HandleFunc("/file/update",handler.FileMetaUpdateHandler)//删除文件
	http.HandleFunc("/file/fastupload",handler.HTTPInterceptor(handler.TryFastUploadHandler)) //秒传文件接口

	// 分块上传接口
	http.HandleFunc("/file/mpupload/init",
		handler.HTTPInterceptor(handler.InitialMultipartUploadHandler))
	http.HandleFunc("/file/mpupload/uppart",
		handler.HTTPInterceptor(handler.UploadPartHandler))
	http.HandleFunc("/file/mpupload/complete",
		handler.HTTPInterceptor(handler.CompleteUploadHandler))


	//用户crud
	http.HandleFunc("/user/signup",handler.SignupHandler)//注册用户
	http.HandleFunc("/user/signin",handler.SignInHandler)//登录
	http.HandleFunc("/user/info",handler.HTTPInterceptor(handler.UserInfoHandler))// 查询用户信息



	h := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", h))
	err := http.ListenAndServe(
		":8080",nil,
		)
	fmt.Println("11111111111")
	if err != nil{
		fmt.Println("Failed to start server,err:%s",err.Error())
	}
}




