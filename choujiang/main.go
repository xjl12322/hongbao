package main

//import "github.com/kataras/iris/_examples/structuring/bootstrap/bootstrap"
import (
	"fmt"
	"hongbao/choujiang/bootrap"
	"hongbao/choujiang/web/routes"
)


func newApp() *bootrap.Bootstrapper {
	app:= bootrap.New("111","222")
	app.Bootstrap()
	//app.Configure()
	return app

}



func main()  {
	app := newApp()
	routes.InitRegister(app)
	app.Application.Run(fmt.Sprintf(":%d",8001))

}










