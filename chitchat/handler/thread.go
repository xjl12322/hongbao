package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/chitchat/models"
	"hongbao/chitchat/pkg/utils"
	"net/http"
	"strconv"
)
// GET /thread/read
// 通过 ID 渲染指定群组页面
func ThreadReadHandler(ctx *gin.Context)  {
	uuid := ctx.Query("id")
	thread,err := models.ThreadByUUID(uuid)

	if err != nil {
		ctx.HTML(http.StatusServiceUnavailable,"error.html",gin.H{"error":err})
	} else {
		posts,err := thread.Posts()

		if err != nil{
			ctx.HTML(http.StatusServiceUnavailable,"error.html",gin.H{"error":err})
		}

		claims := ctx.MustGet("claims")
		var user models.User
		if claims != nil{
			claims := claims.(*utils.CustomClaims)
			ID, _ := strconv.Atoi(claims.ID)
			user = models.User{Id:ID,Name:claims.Name,Email:claims.Email,Password:claims.Password}

			ctx.HTML(http.StatusOK,"auth.thread.html",gin.H{"thread":thread,"user":user,"post":posts})
		}else {

			ctx.HTML(http.StatusOK,"thread.html",gin.H{"thread":thread,"post":posts})
		}
	}
	//ctx.HTML(http.StatusOK,"thread.html",gin.H{})

}

func ThreadNewHandler(ctx *gin.Context)  {
	ctx.HTML(http.StatusOK,"newthread.html",gin.H{})

}


func CreateThreadNewHandler(ctx *gin.Context)  {
	topic := ctx.DefaultPostForm("topic","")
	var user models.User
	claims := ctx.MustGet("claims").(*utils.CustomClaims)
	ID, _ := strconv.Atoi(claims.ID)
	user = models.User{Id:ID,Name:claims.Name,Email:claims.Email,Password:claims.Password}
	if topic == "" {
		ctx.HTML(http.StatusServiceUnavailable,"error.html",gin.H{"error":topic})
	}
	if _,err := user.CreateThread(topic); err != nil {
		ctx.HTML(http.StatusServiceUnavailable,"error.html",gin.H{"error":err})
	}
	ctx.Redirect(http.StatusFound,"/index")

}

// POST /thread/post
// 在指定群组下创建新主题
func CreateThreadPostHandler(ctx *gin.Context)  {
	body := ctx.PostForm("body")
	uuid := ctx.PostForm("uuid")
	thread, err := models.ThreadByUUID(uuid)
	if err != nil {
		fmt.Println("Cannot read thread")
	}
	var user models.User
	claims := ctx.MustGet("claims").(*utils.CustomClaims)
	ID, _ := strconv.Atoi(claims.ID)
	user = models.User{Id:ID,Name:claims.Name,Email:claims.Email,Password:claims.Password}
	if _,err := user.CreatePost(*thread,body); err != nil {
		ctx.HTML(http.StatusServiceUnavailable,"error.html",gin.H{"error":err})
	}
	url := fmt.Sprint("/thread/read?id=", uuid)
	ctx.Redirect(http.StatusFound,url)

}














