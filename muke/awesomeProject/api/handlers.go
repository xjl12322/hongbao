package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"hongbao/muke/awesomeProject/api/dbops"
	"hongbao/muke/awesomeProject/api/defs"
	"hongbao/muke/awesomeProject/api/session"
	"log"
)

//func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
//	fmt.Println("创建用户")
//	res,_ := ioutil.ReadAll(r.Body)
//	log.Printf("%s:",res)
//	ubody := &defs.UserCreadential{}
//	if err := json.Unmarshal(res,ubody);err != nil{
//		log.Printf("%s0",err)
//		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
//		return
//	}
//	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd); err != nil{
//		sendErrorResponse(w,defs.ErrorDBError)
//		return
//	}
//
//	id := session.GenerateNewSessionId(ubody.Username)
//	su := &defs.SignedUp{Success:true,SessionId:id}
//	if resp,err := json.Marshal(su); err != nil{
//		sendErrorResponse(w,defs.ErrorInternalFaults)
//		return
//	}else{
//		sendNormalResponse(w,string(resp),201)
//	}
//
//}

func CreateUser(r *gin.Context)  {
	fmt.Println("创建用户")
	username := r.PostForm("user_name")
	pwd := r.PostForm("pwd")
	ubody := &defs.UserCreadential{}
	ubody.Pwd = pwd
	ubody.Username = username
	//TODO:校验参数

	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd); err != nil{
		sendErrorResponse(r.Writer,defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true,SessionId:id}
	if resp,err := json.Marshal(su); err != nil{
		sendErrorResponse(r.Writer,defs.ErrorInternalFaults)
		return
	}else{
		sendNormalResponse(r.Writer,string(resp),201)
	}

}

func Login(r *gin.Context)  {
	fmt.Println("登录用户")
	username := r.PostForm("user_name")
	pwd := r.PostForm("pwd")
	ubody := &defs.UserCreadential{}
	ubody.Pwd = pwd
	ubody.Username = username
	//TODO:校验参数

	log.Printf("Login url name %s",username)
	log.Printf("Login body name %s",ubody.Username)
	if username != ubody.Username{
		sendErrorResponse(r.Writer,defs.ErrorNotAuthUser)
		return
	}

	log.Printf("%s",ubody.Username)
	pwd,err := dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd %s",ubody.Pwd)
	log.Printf("Login body %s",ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd!= ubody.Pwd{
		sendErrorResponse(r.Writer,defs.ErrorNotAuthUser)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedUp{Success:true,SessionId:id}

	if resp,err := json.Marshal(si);err!=nil{
		sendErrorResponse(r.Writer,defs.ErrorInternalFaults)

	}else{
		sendNormalResponse(r.Writer,string(resp),200)

	}


}

//func GetUserInfo(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
//
//	if !ValidateUser(w,r){
//		log.Printf("unathorized user\n")
//		return
//	}
//	uname := p.ByName("username")
//	u,err := dbops.GetUserCredential(uname)
//	if err != nil{
//		log.Printf("error in GetUserInfo: %s",err)
//		sendErrorResponse(w,defs.ErrorDBError)
//		return
//	}
//	ui := &defs.UserCreadential{}
//
//
//}