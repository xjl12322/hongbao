package routes

import (
	"github.com/gin-gonic/gin"
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
	app.Application.GET("/login",controllers.GetLogin)
	app.Application.GET("/logout",controllers.GetLogout)


	admin := app.Application.Group("/admin")
	admin.Use(gin.BasicAuth(gin.Accounts{
		"admin":"123456",
	}))
	//奖品管理
	admin.GET("/i",controllers.Admin)
	admin.GET("/index",controllers.GetAdmin)
	//app.Application.HandleContext()
	admin.GET("/gift",controllers.GetGift())

	admin.GET("/gift/edit",controllers.GetEdit)
	admin.GET("/sest",controllers.Sest())
	admin.POST("/gift/save",controllers.PostSave(app.Application))
	admin.GET("/gift/delete",controllers.GetDelete)
	admin.GET("/gift/reset",controllers.GetReset)
	//优惠卷管理
	code := admin.Group("/code")
	code.GET("/",controllers.GetCode)
	code.POST("/import",controllers.PostImport)
	code.GET("/delete",controllers.GetCodeDelete)
	code.GET("/reset",controllers.GetCodeReset)

	result := admin.Group("/result")
	result.GET("/",controllers.GetResult)
	result.POST("/delete",controllers.GetResultDelete)
	result.POST("/cheat",controllers.GetResultCheat)
	//result.GET("/gift/reset",controllers.GetReset)



	user := admin.Group("/user")
	user.GET("/",controllers.GetUser)
	user.GET("/black",controllers.GetBlack)
}







////设定请求url不存在的返回值
//router.NoRoute(NoResponse)
//
//
//}
//
////NoResponse 请求的url不存在，返回404
//func NoResponse(c *gin.Context) {
//	//返回404状态码
//	c.JSON(http.StatusNotFound, gin.H{
//		"status": 404,
//		"error":  "404, page not exists!",
//	})
//}









