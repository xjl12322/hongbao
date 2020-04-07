package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/choujiang/models"
	"hongbao/choujiang/services"
	"net/http"
	"strconv"
)
//中奖列表
func GetResult(Ctx *gin.Context)  {
	giftIds := Ctx.DefaultQuery("gift_id","0")
	pages := Ctx.DefaultQuery("page","1")
	uids := Ctx.DefaultQuery("uid","0")
	giftId ,_:= strconv.Atoi(giftIds)
	page ,_:= strconv.Atoi(pages)
	uid ,_:= strconv.Atoi(uids)
	size := 3
	pagePrev := ""
	pageNext := ""
	//数据列表
	var datalist []models.LtResult
	//TODO:
	obj_ResultService := services.NewResultService()
	if giftId > 0{
		datalist = obj_ResultService.SearchByGift(giftId,page,size)
	}else if uid >0 {
		datalist = obj_ResultService.SearchByUser(uid,page,size)
	}else {
		datalist = obj_ResultService.GetAll(page,size)
	}
	total := (page-1)+len(datalist)


	if  len(datalist) >=size{
		//TODO:
		if giftId>0{
			total = int(obj_ResultService.CountByGift(giftId))
		}else if uid>0{
			total = int(obj_ResultService.CountByUser(uid))
		}else {
			total = int(obj_ResultService.CountAll())
		}
		pageNext = fmt.Sprintf("%d",page+1)

	}
	if page >1 {
		pagePrev = fmt.Sprintf("%d",page-1)
	}

	Ctx.HTML(http.StatusOK,
		"admin/result.html",
		gin.H{
			"Title":"管理后台",
			"Channel":"result",
			"GiftId":giftId,
			"Datalist":datalist,
			"Total":total,
			"pagePrev":pagePrev,
			"pageNext":pageNext,
		})

	return
}
////导入优惠卷
//func PostImport(Ctx *gin.Context) {
//	//TODO：
//	gift_ids := Ctx.DefaultQuery("gift_id","0")
//	gift_id ,_:= strconv.Atoi(gift_ids)
//	obj_CodeService := services.NewCodeService()
//	obj_GiftService := services.NewGiftService()
//	if gift_id <1{
//		Ctx.String(http.StatusOK,"没有指定奖品ID，无法进行导入，<a href='' onclick='history.go(-1);return false;'>返回</a>")
//		return
//	}
//	gift := obj_GiftService.Get(gift_id)
//
//	if gift == nil || gift.Id < 1 || gift.Gtype != conf.GtypeCodeDiff{
//		Ctx.String(http.StatusOK,"奖品不存在不是差异化，<a href='' onclick='history.go(-1);return false;'>返回</a>")
//		return
//	}
//	codes := Ctx.PostForm("codes")
//	now := comm.NowUnix()
//	list := strings.Split(codes,"\n")
//	sucNum := 0
//	errNum := 0
//	for _,code := range list {
//		code := strings.TrimSpace(code)
//		if code != ""{
//			data := &models.LtCode{GiftId:gift_id,Code:code,SysCreated:now}
//			err := obj_CodeService.Create(data)
//			if err != nil{
//				errNum++
//			}else {
//				sucNum++
//				//TODO:导入数据库 需要导入缓存
//
//			}
//
//		}
//
//	}
//	html := fmt.Sprintf("成功导入 %d 条，导入失败 %d 条，<a href='/admin/code?gift_id=%d'>返回</a>", sucNum, errNum, gift_id)
//	Ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
//}
//删除优惠卷
func GetResultDelete(Ctx *gin.Context) {
	ids := Ctx.DefaultQuery("id","0")
	id ,_:= strconv.Atoi(ids)
	obj_ResultService := services.NewResultService()
	if id>0{
		obj_ResultService.Delete(id)
	}
	refer := Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/result"
	}
	Ctx.Redirect(http.StatusMovedPermanently,refer)

}

func GetResultCheat(Ctx *gin.Context) {
	ids := Ctx.DefaultQuery("id","0")
	id ,_:= strconv.Atoi(ids)
	obj_ResultService := services.NewResultService()
	if id >0{
		obj_ResultService.Update(&models.LtResult{Id:id,SysStatus:2},[]string{"sys_status"})
	}

	refer := Ctx.GetHeader("Referer")
	if refer == "" {
		refer = "/admin/result"
	}
	Ctx.Redirect(http.StatusMovedPermanently,refer)

}



////更新优惠卷
//func GetCodeReset(Ctx *gin.Context) {
//	ids := Ctx.DefaultQuery("id","0")
//	id ,_:= strconv.Atoi(ids)
//	obj_GiftService := services.NewCodeService()
//	if id>0{
//		obj_GiftService.Update(&models.LtCode{Id:id,SysStatus:0},[]string{"sys_status"})
//	}
//	refer := Ctx.GetHeader("Referer")
//	if refer == "" {
//		refer = "/admin/code"
//	}
//	Ctx.Redirect(http.StatusMovedPermanently,refer)
//
//}
//
//
//


