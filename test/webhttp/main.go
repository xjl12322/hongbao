package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter,r *http.Request)  {
	//获取那日容长度
	lenght := r.ContentLength
	//创建一个字节切片
	body := make([]byte,lenght)
	//读取请求体
	r.Body.Read(body)
	fmt.Fprintln(w,"内容是",string(body))


}

func main() {
	max := http.NewServeMux()
	max.HandleFunc("/ss",handler)
	http.ListenAndServe(":8081",max)
}
