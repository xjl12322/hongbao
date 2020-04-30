package handler

import (
	"github.com/gin-gonic/gin"
	"hongbao/chitchat/models"
	"hongbao/chitchat/pkg/utils"
	"net/http"
	"strconv"
)

func IndexHandle(ctx *gin.Context)  {
	threads, err := models.Threads();
	if err != nil{
		ctx.JSON(http.StatusOK,gin.H{"cuowu":"err"})
	}
	claims := ctx.MustGet("claims")
	var user models.User
	if claims != nil{
		claims := claims.(*utils.CustomClaims)
		ID, _ := strconv.Atoi(claims.ID)
		user = models.User{Id:ID,Name:claims.Name,Email:claims.Email,Password:claims.Password}
		ctx.HTML(http.StatusOK,"index.html",gin.H{"threads":threads,"user":user})
	}else {
		ctx.HTML(http.StatusOK,"index.html",gin.H{"threads":threads})
	}


}


