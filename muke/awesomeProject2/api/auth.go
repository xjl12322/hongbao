package main

import (
	"awesomeProject/api/defs"
	"awesomeProject/api/session"
	"net/http"
)



var HEADER_DIELD_SESSION = "X-Session-Id"
var HEADER_DIELD_UNAME= "X-Session-Id"

//session 效验
func validateUserSession(r *http.Request) bool {
	sid := r.Header.Get(HEADER_DIELD_SESSION)
	if len(sid) == 0{
		return false
	}
	
	uname,ok := session.IsSessionExpired(sid)
	if ok{
		return false
	}
	r.Header.Add(HEADER_DIELD_UNAME,uname)
	return true

}

//
func ValidateUser(w http.ResponseWriter, r *http.Request)bool  {
	uname := r.Header.Get(HEADER_DIELD_UNAME)
	if len(uname ) == 0{

		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return false
	}
	return true

}











