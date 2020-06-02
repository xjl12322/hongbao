package handler

import (
	"encoding/json"
	"hongbao/muke/fiestore-server/meta"
	"hongbao/muke/fiestore-server/util"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"
	dblayer "hongbao/muke/fiestore-server/db"
)

//文件上传
func UploadHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("文件上传")
	if r.Method == "GET"{
		//返回上传的html页面
		data,err := ioutil.ReadFile("./static/view/index.html")
		if err!=nil{
			io.WriteString(w,"internet server error")
			return
		}
		io.WriteString(w,string(data))
	}else if r.Method == "POST"{
		file,head,err := r.FormFile("file")
		if err != nil{
			fmt.Print("Failed to get data,err:%s\n",err.Error())
			return
		}
		defer file.Close()

		fileMeta := meta.FileMeta{
			FileName:head.Filename,
			Location:"/tmp/"+head.Filename,
			UploadAt:time.Now().Format("2006-01-02 15:05:05"),
		}

		newFile,err := os.Create("."+fileMeta.Location)
		if err != nil{
			fmt.Print("Failed to create filed,err:%s\n",err.Error())
			return
		}
		defer newFile.Close()
		fileMeta.FileSize,err = io.Copy(newFile,file)
		if err != nil{
			fmt.Print("Failed to save data into file,err:%s\n",err.Error())
			return
		}

		newFile.Seek(0,0)
		fileMeta.FileSha1 = util.FileSha1(newFile)
		//meta.UpdateFileMeta(fileMeta)   本地
		_ = meta.UpdateFileMetaDB(fileMeta)  //文件上传到mysql
	http.Redirect(w,r,"/file/upload/suc",http.StatusFound)
	}

}
//上传已经完成
func UploadSucHandler(w http.ResponseWriter,r *http.Request)  {
	io.WriteString(w,"upload finished!")


}
//c8649a0c68985786fc596747e4b62cbeb00b0c4e
//获取上传文件信息  //传入文件sha1 加密值返回
func GetFileMetaHandler(w http.ResponseWriter,r *http.Request)  {
	fmt.Println("获取上传文件信息")
	r.ParseForm()
	filehash := r.Form["filehash"][0]
	//fMeta := meta.GetFileMeta(filehash)
	fMeta,err:= meta.GetFileMeteDB(filehash)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data,err := json.Marshal(fMeta)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)

}

//文件下载接口
func DownloadHandler(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	fsha1 := r.Form.Get("filehash")
	fm := meta.GetFileMeta(fsha1)
	f,err := os.Open("."+fm.Location)
	if err != nil{

		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer f.Close()

	data,err := ioutil.ReadAll(f)
	if err != nil{

		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type","application/octect-stream")
	w.Header().Set("content-descrption","attachment;filename=\""+fm.FileName+"\"")
	w.Write(data)


}

////更新（重命名）文件的信息
func FileMetaUpdateHandler(w http.ResponseWriter,r *http.Request)  {
	r.ParseForm()
	opType := r.Form.Get("op")  //模式为0 默认0
	fileSha1 := r.Form.Get("filehash")  //要修改的文件hash值
	newfileName := r.Form.Get("filename")//要修改文件名称

	if opType !="0"{
		w.WriteHeader(http.StatusForbidden)
		return
	}
	if r.Method !="POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	curFileMeta := meta.GetFileMeta(fileSha1)
	curFileMeta.FileName = newfileName
	meta.UpdateFileMeta(curFileMeta)

	data,err := json.Marshal(curFileMeta)
	if err != nil{
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
// FileDeleteHandler : 删除文件及元信息
func FileDeleteHandler(w http.ResponseWriter,r *http.Request) {
	r.ParseForm()
	fileSha1 := r.Form.Get("filehash")

	fMeta := meta.GetFileMeta(fileSha1)
	os.Remove(fMeta.Location)

	meta.RemoveFileMeta(fileSha1)

	w.WriteHeader(http.StatusOK)

}


// TryFastUploadHandler : 尝试秒传接口
func TryFastUploadHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// 1. 解析请求参数
	username := r.Form.Get("username")
	filehash := r.Form.Get("filehash")
	filename := r.Form.Get("filename")
	filesize, _ := strconv.Atoi(r.Form.Get("filesize"))

	// 2. 从文件表中查询相同hash的文件记录
	fileMeta, err := meta.GetFileMeteDB(filehash)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. 查不到记录则返回秒传失败
	if fileMeta == nil {
		resp := util.RespMsg{
			Code: -1,
			Msg:  "秒传失败，请访问普通上传接口",
		}
		w.Write(resp.JSONBytes())
		return
	}

	// 4. 上传过则将文件信息写入用户文件表， 返回成功
	suc := dblayer.OnUserFileUploadFinished(
		username, filehash, filename, int64(filesize))
	if suc {
		resp := util.RespMsg{
			Code: 0,
			Msg:  "秒传成功",
		}
		w.Write(resp.JSONBytes())
		return
	}
	resp := util.RespMsg{
		Code: -2,
		Msg:  "秒传失败，请稍后重试",
	}
	w.Write(resp.JSONBytes())
	return
}



