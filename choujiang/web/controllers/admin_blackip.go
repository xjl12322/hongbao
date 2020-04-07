package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/choujiang/comm"
	"hongbao/choujiang/conf"
	"hongbao/choujiang/models"
	"hongbao/choujiang/services"
	"net/http"
	"strconv"
	"strings"
)
//优惠卷列表
func GetBlackIp(Ctx *gin.Context)  {

	pages := Ctx.DefaultQuery("page","1")
	page ,_:= strconv.Atoi(pages)
	size := 3
	pagePrev := ""
	pageNext := ""
	obj_BlackipService := services.NewBlackipService()
	//数据列表
	datalist := obj_BlackipService.GetAll(page,size)
	total := (page-1)+len(datalist)


	if len(datalist)>=size{
		total = int(obj_BlackipService.CountAll())
		pageNext = fmt.Sprintf("%d",page+1)

	}
	if page >1{
		pagePrev = fmt.Sprintf("%d",page-1)
	}

	Ctx.HTML(http.StatusOK,
		"admin/user.html",
		gin.H{
			"Title":"管理后台",
			"Channel":"user",
			"Datalist":datalist,
			"Total":total,
			"pagePrev":pagePrev,
			"pageNext":pageNext,
		})

	return
}
//设置黑名单
func GetBlack2(Ctx *gin.Context) {
	ids := Ctx.Query("id")
	times := Ctx.DefaultQuery("time","0")
	id ,_:= strconv.Atoi(ids)
	t,_:= strconv.Atoi(times)
	if t>0{
		t = t*86400+comm.NowUnix()
	}
	obj_BlackipService := services.NewBlackipService()
	obj_BlackipService.Update(&models.LtBlackip{Id:id,Blacktime:t,SysUpdated:comm.NowUnix()},[]string{"blacktime"})


	Ctx.Redirect(http.StatusMovedPermanently,"/admin/blackip")



	//TODO：
	gift_ids := Ctx.DefaultQuery("gift_id","0")
	gift_id ,_:= strconv.Atoi(gift_ids)
	obj_CodeService := services.NewCodeService()
	obj_GiftService := services.NewGiftService()
	if gift_id <1{
		Ctx.String(http.StatusOK,"没有指定奖品ID，无法进行导入，<a href='' onclick='history.go(-1);return false;'>返回</a>")
		return
	}
	gift := obj_GiftService.Get(gift_id)

	if gift == nil || gift.Id < 1 || gift.Gtype != conf.GtypeCodeDiff{
		Ctx.String(http.StatusOK,"奖品不存在不是差异化，<a href='' onclick='history.go(-1);return false;'>返回</a>")
		return
	}
	codes := Ctx.PostForm("codes")
	now := comm.NowUnix()
	list := strings.Split(codes,"\n")
	sucNum := 0
	errNum := 0
	for _,code := range list {
		code := strings.TrimSpace(code)
		if code != ""{
			data := &models.LtCode{GiftId:gift_id,Code:code,SysCreated:now}
			err := obj_CodeService.Create(data)
			if err != nil{
				errNum++
			}else {
					sucNum++
					//TODO:导入数据库 需要导入缓存

			}

		}

	}
	html := fmt.Sprintf("成功导入 %d 条，导入失败 %d 条，<a href='/admin/code?gift_id=%d'>返回</a>", sucNum, errNum, gift_id)
	Ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
////删除优惠卷
//func GetCodeDelete(Ctx *gin.Context) {
//	ids := Ctx.DefaultQuery("id","0")
//	id ,_:= strconv.Atoi(ids)
//	obj_CodeService := services.NewCodeService()
//	if id>0{
//		obj_CodeService.Delete(id)
//	}
//	refer := Ctx.GetHeader("Referer")
//	if refer == "" {
//		refer = "/admin/code"
//	}
//	Ctx.Redirect(http.StatusMovedPermanently,refer)
//	//Ctx.Abort()
//	//Ctx.HandlerName()
//	//Ctx.HandlerNames()
//	//Ctx.Writer.
//	//return mvc.Response{
//	//	Path: refer,
//	//}
//
//}
////恢复更新优惠卷
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





