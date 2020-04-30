package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/chitchat/pkg/utils"
	"net/http"
	"time"
)

//// 通过 Cookie 判断用户是否已登录
//func session(writer http.ResponseWriter, request *http.Request) (sess models.Session, err error) {
//	cookie, err := request.Cookie("_cookie")
//	if err == nil {
//		sess = models.Session{Uuid: cookie.Value}
//		if ok, _ := sess.Check(); !ok {
//			err = errors.New("Invalid session")
//		}
//	}
//	return
//}
//
//func Session()gin.HandlerFunc  {
//	return func(ctx *gin.Context) {
//		  ctx.Cookie("_cookie")
//	}
//}

func Loginrequits() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		fmt.Println("中间件")
		//url := ctx.Request.URL.Path
		//if url == "index"{
		//
		//}
		jwt := utils.JWTS
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.Abort()
			ctx.HTML(http.StatusOK, "login.html", gin.H{})
		} else {
			claims, err := jwt.ParseToken(cookie)
			if err != nil {
				ctx.Abort()
				ctx.HTML(http.StatusOK, "login.html", gin.H{})

			} else if time.Now().Unix() > claims.ExpiresAt {
				ctx.Abort()
				ctx.HTML(http.StatusOK, "login.html", gin.H{})
			}
			ctx.Set("claims",claims)
			ctx.Next()
		}
	}
}


func Indexrequited() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		jwt := utils.JWTS
		cookie, err := ctx.Cookie("token")
		if err != nil {
			ctx.Set("claims",nil)
			ctx.Next()
		} else {
			claims, err := jwt.ParseToken(cookie)
			if err != nil {
				ctx.Set("claims",nil)
				ctx.Next()
			} else if time.Now().Unix() > claims.ExpiresAt {
				ctx.Set("claims",nil)
				ctx.Next()
			}
			ctx.Set("claims",claims)
			ctx.Next()
		}
	}
}














