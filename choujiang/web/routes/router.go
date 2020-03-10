package routes

import (

	"hongbao/choujiang/bootrap"
	"hongbao/choujiang/web/controllers"
)
//func Configure()  {
//	giftservice := services.NewGiftService()
//
//
//}



func InitRegister(app *bootrap.Bootstrapper){
	app.Application.GET("/",controllers.Get)
	app.Application.GET("/gifts",controllers.GetGifts)



}














