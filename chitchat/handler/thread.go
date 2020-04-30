package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/chitchat/models"
	"hongbao/chitchat/pkg/utils"
	"net/http"
	"strconv"
)

func ThreadReadHandler(ctx *gin.Context)  {
	ctx.

	ctx.HTML(http.StatusOK,"thread.html",gin.H{})

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
		fmt.Println("Cannot create -- ")
	}
	if _,err := user.CreateThread(topic); err != nil {
		fmt.Println("Cannot create thread")
	}
	ctx.Redirect(http.StatusFound,"/index")

}

















