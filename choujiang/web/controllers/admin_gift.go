package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/choujiang/comm"
	"hongbao/choujiang/models"
	"hongbao/choujiang/services"
	"hongbao/choujiang/viewmodels"
	"net/http"
	"strconv"
	"time"
)

func GetGift()gin.HandlerFunc  {
	return func(Ctx *gin.Context) {
		fmt.Println("GetGift.GetGift.GetGift.GetGift")
		// 数据列表
		//routes.InitRegister()
		obj_GiftService := services.NewGiftService()
		datalist := obj_GiftService.GetAll()
		//奖品发放计划数据 时间类型转换
		for i,giftInfo := range datalist{
			//datalist[i].TimeBegin = comm.FormatFromUnixTime(int64(giftInfo.TimeBegin))
			//datalist[i].TimeEnd= comm.FormatFromUnixTime(int64(giftInfo.TimeEnd))

			prizedata := make([][2]int,0)
			err := json.Unmarshal([]byte(giftInfo.PrizeData),&prizedata)
			if err != nil || len(prizedata)<1{
				datalist[i].PrizeData = "[]"
			}else {
				newpd := make([]string,len(prizedata))
				for index,pd := range prizedata{
					ct := comm.FormatFromUnixTime(int64(pd[0]))
					newpd[index] = fmt.Sprintf("[%s]:%d",ct,pd[1])
				}
				str,err := json.Marshal(newpd)
				if err != nil&&len(str)>0{
					datalist[i].PrizeData = string(str)
				}else {
					datalist[i].PrizeData = "[]"
				}
			}

		}
		total := len(datalist)
		Ctx.HTML(http.StatusOK,
			"admin/gift.html",
			gin.H{
				"Title":"管理后台",
				"Channel":"gift",
				"Datalist":datalist,
				"Total":total,
			})
		return

	}


}
//func GetGift(Ctx *gin.Context) {
//
//
//
//
//}

func GetEdit(Ctx *gin.Context) {
	ids := Ctx.DefaultQuery("id","0")
	id ,_:= strconv.Atoi(ids)
	giftInfo := viewmodels.ViewGift{}
	if id>0{
		obj_GiftService := services.NewGiftService()
		data := obj_GiftService.Get(id)
		giftInfo.Id = data.Id
		giftInfo.Title = data.Title
		giftInfo.PrizeNum = data.PrizeNum
		giftInfo.PrizeCode = data.PrizeCode
		giftInfo.PrizeTime = data.PrizeTime
		giftInfo.Img = data.Img
		giftInfo.Displayorder = data.Displayorder
		giftInfo.Gtype = data.Gtype
		giftInfo.Gdata = data.Gdata
		giftInfo.TimeBegin = comm.FormatFromUnixTime(int64(data.TimeBegin))
		giftInfo.TimeEnd = comm.FormatFromUnixTime(int64(data.TimeEnd))
	}

	Ctx.HTML(http.StatusOK,
		"admin/giftEdit.html",
		gin.H{
			"Title":"管理后台",
			"Channel": "gift",
			"info":    giftInfo,
		})
}
//更新奖品信息提交
func PostSave(r *gin.Engine)gin.HandlerFunc {
	return func(Ctx *gin.Context) {
		data := viewmodels.ViewGift{}
		err := Ctx.ShouldBind(&data)
		if err != nil{
			fmt.Println("admin_gift.postsave readform err",err)
			Ctx.String(http.StatusOK,fmt.Sprintf("readform 转换异常err!=%s",err))
			return
		}
		giftInfo := models.LtGift{}
		giftInfo.Id = data.Id
		giftInfo.Title = data.Title
		giftInfo.PrizeNum = data.PrizeNum
		giftInfo.PrizeCode = data.PrizeCode
		giftInfo.PrizeTime = data.PrizeTime
		giftInfo.Img = data.Img
		giftInfo.Displayorder = data.Displayorder
		giftInfo.Gtype = data.Gtype
		giftInfo.Gdata = data.Gdata
		t1, err1 := comm.ParseTime(data.TimeBegin)
		t2, err2 := comm.ParseTime(data.TimeEnd)
		if err1 != nil || err2 != nil {
			Ctx.String(http.StatusOK,fmt.Sprintf("开始时间、结束时间的格式不正确, err1=%s, err2=%s", err1, err2))
		}
		giftInfo.TimeBegin = int(t1.Unix())
		giftInfo.TimeEnd = int(t2.Unix())
		obj_GiftService := services.NewGiftService()
		if giftInfo.Id >0{
			//数据更新   剩余数量做更行处理
			datainfo := obj_GiftService.Get(giftInfo.Id)
			if datainfo != nil&&datainfo.Id >0{
				if datainfo.PrizeNum != giftInfo.PrizeNum {
					//奖品数量发生变化
					giftInfo.LeftNum = datainfo.LeftNum-datainfo.PrizeNum-giftInfo.PrizeNum
					if giftInfo.LeftNum<0 || giftInfo.PrizeNum <=0{
						giftInfo.LeftNum = 0
					}
					//TODO: 后期处理
				}
				if datainfo.PrizeTime != giftInfo.PrizeTime{
					//TODO:发奖的周期发生变化
				}
				giftInfo.SysUpdated =int(time.Now().Unix())
				obj_GiftService.Update(&giftInfo,[]string{""})
			}else {
				giftInfo.Id = 0
			}
		}else {
			giftInfo.LeftNum = giftInfo.PrizeNum
			giftInfo.SysIp = comm.ClientIP(Ctx.Request)
			giftInfo.SysCreated = int(time.Now().Unix())
			obj_GiftService.Create(&giftInfo)
		}

		Ctx.Redirect(http.StatusMovedPermanently,"/admin/gift")


	}
}
//func PostSave(Ctx *gin.Context) {
//
//}
func Sest()gin.HandlerFunc {
	return func(Ctx *gin.Context) {
		Ctx.String(http.StatusOK,"okkkkkk")
		return


	}
}

func GetDelete(Ctx *gin.Context) {
	ids:= Ctx.Query("id")
	id,err := strconv.Atoi(ids)
	if err == nil{
		obj_GiftService := services.NewGiftService()
		obj_GiftService.Delete(id)
	}
	Ctx.Redirect(http.StatusMovedPermanently,"/admin/gift")
	return


}
func GetReset(Ctx *gin.Context) {
	ids:= Ctx.Query("id")
id,err := strconv.Atoi(ids)
	if err == nil{
		obj_GiftService := services.NewGiftService()
		obj_GiftService.Update(&models.LtGift{Id:id,SysStatus:0},[]string{"sys_status"})
	}
	Ctx.Redirect(http.StatusMovedPermanently,"/admin/gift")
	return


}