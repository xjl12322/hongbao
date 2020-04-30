package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/chitchat/models"
	"hongbao/chitchat/pkg/common"
	"hongbao/chitchat/pkg/utils"
	"net/http"
	"strconv"

)

func LoginHandle(ctx *gin.Context)  {
	threads, err := models.Threads();
	if err != nil{
		ctx.JSON(http.StatusOK,gin.H{"cuowu":"err"})
	}
	claims := ctx.MustGet("claims").(*utils.CustomClaims)
	ID, _ := strconv.Atoi(claims.ID)
	user := models.User{Id:ID,Name:claims.Name,Email:claims.Email,Password:claims.Password}
	ctx.HTML(http.StatusOK,"index.html",gin.H{"threads":threads,"user":user})
	}





func LoginAccountHandle(ctx *gin.Context)  {
	fmt.Println("登录验证")
	var user models.User
	user.Password = ctx.PostForm("password")
	user.Email = ctx.PostForm("email")
	userdb, err := models.UserByEmail(user.Email)
	if err != nil {
		fmt.Println("Cannot find user")
	}
	if userdb.Password == common.Encrypt(user.Password){
		ids :=strconv.Itoa(userdb.Id)
		token, _ := utils.JWTS.GenerateToken(ids,userdb.Name,userdb.Email,userdb.Password)
		ctx.SetCookie("token", token, 999999, "/", "localhost", false, true)
		ctx.Redirect(http.StatusMovedPermanently,"/index")

		return
	}

}


func SignupHandle(ctx *gin.Context)  {
	threads, err := models.Threads();
	if err != nil{
		ctx.JSON(http.StatusOK,gin.H{"cuowu":"err"})
	}
	ctx.HTML(http.StatusOK,"signup.html",threads)
}


// 注册新用户
func SignupAccountHandle(ctx *gin.Context)  {
	var user models.User
	user.Name = ctx.PostForm("name")
	user.Email = ctx.PostForm("email")
	user.Password = ctx.PostForm("password")
	if err := user.Create();err != nil{
		ctx.JSON(http.StatusOK,gin.H{"cuowu":err})
		return
	}
	ctx.Redirect(http.StatusMovedPermanently,"/login")
}

func LoginOutHandle(ctx *gin.Context)  {
	ctx.SetCookie("token", "", -100, "/", "localhost", false, true)
	ctx.Redirect(http.StatusMovedPermanently,"/index")
}





