package main

import (
	"github.com/gin-gonic/gin"
	"hongbao/muke/awesomeProject/api/defs"
	"hongbao/muke/awesomeProject/api/session"
	"net/http"
)



var HEADER_DIELD_SESSION = "X-Session-Id"
var HEADER_DIELD_UNAME= "X-User-Name"

//session 效验
func validateUserSession() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sid := ctx.GetHeader(HEADER_DIELD_SESSION)
		if len(sid) == 0{
			ctx.Abort()
			return
		}
		uname,ok := session.IsSessionExpired(sid)
		if ok{
			ctx.Abort()
			return
		}
		ctx.Header(HEADER_DIELD_UNAME,uname)
		ctx.Next()

	}
}

//session 效验
//func validateUserSession(r *http.Request) bool {
//	sid := r.Header.Get(HEADER_DIELD_SESSION)
//	if len(sid) == 0{
//		return false
//	}
//
//	uname,ok := session.IsSessionExpired(sid)
//	if ok{
//		return false
//	}
//	r.Header.Add(HEADER_DIELD_UNAME,uname)
//	return true
//
//}

//user 校验
func ValidateUser(w http.ResponseWriter, r *http.Request)bool  {
	uname := r.Header.Get(HEADER_DIELD_UNAME)
	if len(uname ) == 0{

		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true

}











