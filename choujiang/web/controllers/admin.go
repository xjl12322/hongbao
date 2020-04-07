package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/choujiang/models"
	"hongbao/choujiang/services"
	"strconv"

	"net/http"
)

func Admin(Ctx *gin.Context) {
	Ctx.Header("Content-Type", "text/html")
	Ctx.HTML(http.StatusOK,"admin/index.html",gin.H{
		"Title":"管理后台",
		"Channel":"",
	})

}

//type AdminResultController struct {
//	Ctx            iris.Context
//	ServiceGift    services.GiftService
//	ServiceCode    services.CodeService
//	ServiceResult  services.ResultService
//	ServiceUserday services.UserdayService
//	ServiceBlackip services.BlackipService
//}

func GetAdmin(Ctx *gin.Context) {
	giftId := Ctx.DefaultQuery("gift_id", "0")
	uid := Ctx.DefaultQuery("uid", "0")
	page := Ctx.DefaultQuery("page", "1")
	size := 10
	pagePrev := ""
	pageNext := ""
	// 数据列表
	var datalist []models.LtResult
	obj_GiftService := services.NewResultService()

	pages,_:= strconv.Atoi(page)
	if giftId,_:=strconv.Atoi(giftId);giftId>0{
		datalist = obj_GiftService.SearchByGift(giftId,pages,size)
	} else if uid,_:=strconv.Atoi(uid);uid > 0 {
		datalist =obj_GiftService.SearchByUser(uid,pages,size)
	} else {
		datalist =obj_GiftService.GetAll(pages, size)
	}
	total := (pages-1)+len(datalist)
	// 数据总数
	if len(datalist) >= size {
		if giftId,_:=strconv.Atoi(giftId);giftId>0{
			total = int(obj_GiftService.CountByGift(giftId))
		}else if uid,_:=strconv.Atoi(uid);uid > 0{
			total = int(obj_GiftService.CountByUser(uid))
		}else {
			total = int(obj_GiftService.CountAll())
	}
		pageNext = fmt.Sprintf("%d", pages+1)

	}
	if pages>1{
		pagePrev = fmt.Sprintf("%d", pages-1)
	}

	Ctx.HTML(http.StatusOK,"admin/result.html",gin.H{
		"Title":    "管理后台",
		"Channel":  "result",
		"GiftId":   giftId,
		"Uid":      uid,
		"Datalist": datalist,
		"Total":    total,
		"PagePrev": pagePrev,
		"PageNext": pageNext,
	})
}

//func GetAdmin(Ctx *gin.Context) {
//
//
//
//}



