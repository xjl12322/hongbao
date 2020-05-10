package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type middleWareHandler struct {
	r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router)http.Handler  {
	m := middleWareHandler{}
	m.r = r
	return m
}

func (m middleWareHandler) ServeHTTP(w http.ResponseWriter,r *http.Request)  {
	//check session
	m.r.ServeHTTP(w,r)
}

func RegisterHandlers() *httprouter.Router  {
	router:= httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/:user_name", Login)





	return router
}



func main()  {
	fmt.Println("sta")
	r:= RegisterHandlers()
	mh := NewMiddleWareHandler(r)
	http.ListenAndServe(
		":8080",mh)


}




