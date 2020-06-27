package session

import (
	"awesomeProject/api/dbops"
	"awesomeProject/api/defs"
	"awesomeProject/api/utils"
	"sync"
	"time"
)



var sessionMap *sync.Map

func nowInMilli() int64 {
	return time.Now().UnixNano()/1000000 //时间错紧缺到毫秒
}

func deleteExpiredSession(sid string)  {
	sessionMap.Delete(sid)
	dbops.DeleteSession(sid)
}

func init()  {
	sessionMap = &sync.Map{}
	
}
//获取全部session
func LoadSessionsFromDB()  {
	r, err := dbops.RetrieveAllSessions()
	if err != nil{
		return
	}
	r.Range(func(k, v interface{}) bool {
		ss := v.(*defs.SimpleSession)
		sessionMap.Store(k,ss)
		return true
	})


}
//创建session
func GenerateNewSessionId(un string) string {
	id, _ :=utils.NewUUID()
	ct := time.Now().UnixNano()/1000000 //时间错紧缺到毫秒
	ttl := ct+30*60*1000  //超过30分钟过期
	ss := &defs.SimpleSession{Username:un,TTL:ttl}
	sessionMap.Store(id,ss)
	dbops.InsertSession(id,ttl,un)
	return id

}

func IsSessionExpired(sid string) (string, bool) {
	ss,ok := sessionMap.Load(sid)
	if ok {
		ct := nowInMilli()
		if ss.(*defs.SimpleSession).TTL <ct{
			deleteExpiredSession(sid)  //没有过期
			return "",true
		}
		return ss.(*defs.SimpleSession).Username,false
	}
	return "",true
}

