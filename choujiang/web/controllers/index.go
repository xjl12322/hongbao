package controllers

import (
	"fmt"
	"hongbao/choujiang/comm"
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
	Ctx.HTML(http.StatusOK,"public/contentx.html",nil)
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
// 将登录的用户信息设置到cookie中
func GetLogin(Ctx *gin.Context) {
	uid := comm.Random(100000)
	loginuser := models.ObjLoginuser{
		Uid:uid,
		Username:fmt.Sprintf("admin-%d",uid),
		Now:comm.NowUnix(),
		Ip:comm.ClientIP(Ctx.Request),

	}
	comm.SetLoginuser(Ctx.Writer,&loginuser)
	comm.Redirect(Ctx.Writer,"public/contentx.html?from=login")
}

func GetLogout(Ctx *gin.Context) {
	comm.SetLoginuser(Ctx.Writer,nil)
	comm.Redirect(Ctx.Writer,"public/contentx.html?from=logout")   // 跳转重定向URL

}













