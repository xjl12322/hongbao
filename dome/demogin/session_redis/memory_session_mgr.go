package session

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"hongbao/dome"
	"sync"
	"time"
)

type RedisSessionMgr struct {
	//redis 地址
	addr string
	//密码
	passwd string
	//连接池
	pool *redis.Pool
	//锁
	rwlock sync.RWMutex
	//大map
	sessionMap map[string]dome.Session
}


// 构造函数
func NewRedisSeesionMgr() dome.SessionMgr {
	sr := &RedisSessionMgr{
		sessionMap: make(map[string]dome.Session, 1024),
	}
	return sr
}
// 构造函数
func (r *RedisSessionMgr) Init(addr string, options ...string) (err error) {
	if len(options)>0{
		r.passwd = options[0]
	}

	//创建连接池
	r.pool = myPool(addr,r.passwd)
	r.addr = addr
	return
}

func myPool(addr,passwd string)(*redis.Pool)  {
	return &redis.Pool{
		MaxIdle:64,
		MaxActive:1000,
		IdleTimeout:240 *time.Second,
		Dial: func() (redis.Conn,error) {
			conn,err:= redis.Dial("tcp",addr)
			if  err != nil{
				return nil,err
			}
			//若有密码，判断
			if _,err := conn.Do("AUTH",passwd);err!= nil{
				conn.Close()
				return nil,err
			}
			return conn,err
		},
		//链接测试 开发使用 正式上线注释掉
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_,err :=c.Do("PING")
			return err
		},
	}
}


func (r *dome.RedisSession)CreateSession()(session dome.Session,err error)  {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()

	id := uuid.NewV4()
	// 转string
	sessionId := id.String()
	// 创建个session

	session = dome.NewRedisSession(sessionId,r.pool)
	//加入到大map中
	r.sessionMap[sessionId] = session
	

	return
}

func (r *dome.RedisSession) Get(sessionId string)(session dome.Session,err error)  {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	session,ok := r.sessionMap[sessionId]
	if !ok{
		err = errors.New("session not exeit")
		return
	}
	return
}











