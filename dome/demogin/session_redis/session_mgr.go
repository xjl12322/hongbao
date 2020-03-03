package session

import "hongbao/dome"

// 定义管理者，管理所有session
type SessionMgr interface {
	// 初始化
	Init(addr string,options ...string)(err error)
	CreateSession()(session dome.Session,err error)
	Get(sessionId string)(session dome.Session,err error)



}








