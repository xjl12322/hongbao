package session

import (
	"errors"
	uuid "github.com/satori/go.uuid"
	"hongbao/dome"
	"sync"
)

type MemorySeesionMgr struct {

	sessionMap map[string]dome.Session
	rwlock sync.RWMutex


}


// 构造函数
func NewMemorySeesionMgr() *MemorySeesionMgr {
	sr := &MemorySeesionMgr{
		sessionMap: make(map[string]dome.Session, 1024),
	}
	return sr
}
// 构造函数

func (s *MemorySeesionMgr) Init(addr string, options ...string) (err error) {
	return
}


func (s *MemorySeesionMgr)CreateSession()(session dome.Session,err error)  {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	// go get github.com/satori/go.uuid
	// 用uuid作为sessionId
	uuid.NewV4()
	id := uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建个session
	session = dome.NewMemorySession(sessionId)
	//加入到大map中
	s.sessionMap[sessionId] = session
	

	return
}

func (s *MemorySeesionMgr) Get(sessionId string)(session dome.Session,err error)  {
	s.rwlock.Lock()
	defer s.rwlock.Unlock()
	session,ok := s.sessionMap[sessionId]
	if !ok{
		err = errors.New("session not exeit")
		return
	}
	return
}











