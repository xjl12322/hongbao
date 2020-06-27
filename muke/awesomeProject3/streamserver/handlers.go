package main

import (
	"github.com/julienschmidt/httprouter"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"time"
)



//客户端获取播放视频
func streamHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {

	vid := p.ByName("vid-id")
	vl := VIDEO_DIR +vid

	video,err := os.Open(vl) //打开文件二进制
	if err != nil{
		log.Printf("Error when try to open file: %v", err)
		sendErrorResponse(w,http.StatusInternalServerError,"internal erro")
		return
	}
	w.Header().Set("Content-Type","video/mp4") //定义播放协议头浏览器自动解析mp4

	http.ServeContent(w,r,"",time.Now(),video) //浏览器播放
	defer video.Close()


}

//客户端更新
func uploadHandler(w http.ResponseWriter,r *http.Request, p httprouter.Params)  {
	r.Body = http.MaxBytesReader(w,r.Body,MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err!=nil{
		sendErrorResponse(w,http.StatusBadRequest,"文件太大了")
		return
	}

	file,_,err := r.FormFile("file") //获取name<from name=
	if err!=nil{
		sendErrorResponse(w,http.StatusInternalServerError,"服务器错误")
		return
	}
	data,err := ioutil.ReadAll(file)
	if err != nil{
		log.Printf("Read file error: %v",err)
	}

	fn := p.ByName("vid-id")
	err = ioutil.WriteFile(VIDEO_DIR+fn,data,0666)
	if err != nil{
		log.Printf("write file error：%v",err)
		sendErrorResponse(w,http.StatusInternalServerError,"文件写入错误")
		return

	}
	w.WriteHeader(http.StatusCreated)
	io.WriteString(w,"上次文件视频成功")
}


