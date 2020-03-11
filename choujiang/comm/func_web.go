package comm

import (
	"fmt"
	"hongbao/choujiang/models"
	"log"
	"net"
	"net/http"
	"net/url"
	"strconv"
)
// 得到客户端IP地址
func ClientIP(request *http.Request)string  {
	host,_,_:= net.SplitHostPort(request.RemoteAddr)
	return host

}

// 跳转重定向URL
func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location", url)
	writer.WriteHeader(http.StatusFound)
}


// 从cookie中得到当前登录的用户
func GetLoginUser(request *http.Request) *models.ObjLoginuser{
	c,err := request.Cookie("lottery_loginuser")
	if err !=nil{
		return nil
	}
	params,err := url.ParseQuery(c.Value)   //a=1 这样的字符转map字典
	if err != nil{
		return nil
	}
	uid,err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	// Cookie最长使用时长
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix()-now > 86400*30 {
		return nil
	}
	//// IP修改了是不是要重新登录判断
	//ip := params.Get("ip")
	//if ip != ClientIP(request) {
	//	return nil
	//}
	// 登录信息
	loginuser := &models.ObjLoginuser{}
	loginuser.Uid = uid
	loginuser.Username = params.Get("username")
	loginuser.Now = now
	loginuser.Ip = ClientIP(request)
	loginuser.Sign = params.Get("sign")
	if err != nil {
		log.Println("fuc_web GetLoginUser Unmarshal ", err)
		return nil
	}
	return loginuser
}



// 根据登录用户信息生成加密字符串
func createLoginuserSign(loginuser *models.ObjLoginuser){
	str := fmt.Sprintf("uid=%d&username=%s&secret=%s&now=%d",loginuser.Uid,loginuser.Username,loginuser.Now)
	sign :=

}



