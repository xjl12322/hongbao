package main

import (
	"awesomeProject/scheduler/taskrunner"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func RegisterHandlers()*httprouter.Router  {

	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router

}
func main()  {
	fmt.Println("sta")
	go taskrunner.Start()
	r := RegisterHandlers()
	http.ListenAndServe(":9001", r)


}




