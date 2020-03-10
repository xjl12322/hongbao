package controllers

import (
	"fmt"
	"hongbao/choujiang/models"

	"github.com/gin-gonic/gin"
	"hongbao/choujiang/services"
	"net/http"
)





type IndexController struct {
//	hander func(*gin.Context)
//	ServiceUser services.UserService
//	ss services.GiftService
	
	//ServiceCode services.CodeService
	//ServiceResult services.ResultService
	//ServiceUserday services.UserdayService
	//ServiceBlackip services.BlackipService
}





// http://localhost:8080/
func Get(Ctx *gin.Context) {
	Ctx.Header("Content-Type", "text/html")
	Ctx.String(http.StatusOK,fmt.Sprintf("welcome to Go抽奖系统，<a href='/public/index.html'>开始抽奖</a>"))
	return
	
}


// http://localhost:8080/gifts
func GetGifts(Ctx *gin.Context) {
	obj_GiftService := services.NewGiftService()
	datalist := obj_GiftService.GetAll()
	list := make([]models.LtGift,0)
	for _,data := range datalist{
		if data.SysStatus == 0{
			list = append(list,data)
		}
	}
	Ctx.JSON(http.StatusOK,gin.H{
		"code":0,
		"msg":"",
		"gifts":list,
	})
	return
}

func GetMyprize(Ctx *gin.Context) {


}















