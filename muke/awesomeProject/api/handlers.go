package main

import (
	"awesomeProject/api/dbops"
	"awesomeProject/api/defs"
	"awesomeProject/api/session"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	res,_ := ioutil.ReadAll(r.Body)
	ubody := &defs.UserCreadential{}
	if err := json.Unmarshal(res,ubody); err!=nil{
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	if err := dbops.AddUserCredential(ubody.Username,ubody.Pwd); err != nil{
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}

	id := session.GenerateNewSessionId(ubody.Username)
	su := &defs.SignedUp{Success:true,SessionId:id}
	if resp,err := json.Marshal(su); err != nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)
		return
	}else{
		sendNormalResponse(w,string(resp),201)
	}

}

func Login(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {
	res,_ := ioutil.ReadAll(r.Body)
	log.Printf("%s:",res)
	ubody := &defs.UserCreadential{}
	if err := json.Unmarshal(res,ubody);err != nil{
		log.Printf("%s0",err)
		sendErrorResponse(w,defs.ErrorRequestBodyParseFailed)
		return
	}
	uname := p.ByName("username")
	log.Printf("Login url name %s",uname)
	log.Printf("Login body name %s",ubody.Username)
	if uname != ubody.Username{
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return
	}

	log.Printf("%s",ubody.Username)
	pwd,err := dbops.GetUserCredential(ubody.Username)
	log.Printf("Login pwd %s",ubody.Pwd)
	log.Printf("Login body %s",ubody.Pwd)
	if err != nil || len(pwd) == 0 || pwd!= ubody.Pwd{
		sendErrorResponse(w,defs.ErrorNotAuthUser)
		return
	}
	id := session.GenerateNewSessionId(ubody.Username)
	si := &defs.SignedUp{Success:true,SessionId:id}

	if resp,err := json.Marshal(si);err!=nil{
		sendErrorResponse(w,defs.ErrorInternalFaults)

	}else{
		sendNormalResponse(w,string(resp),200)

	}


}

func GetUserInfo(w http.ResponseWriter,r *http.Request,p httprouter.Params)  {

	if !ValidateUser(w,r){
		log.Printf("unathorized user\n")
		return
	}
	uname := p.ByName("username")
	u,err := dbops.GetUserCredential(uname)
	if err != nil{
		log.Printf("error in GetUserInfo: %s",err)
		sendErrorResponse(w,defs.ErrorDBError)
		return
	}
	ui := &defs.UserCreadential{}


}