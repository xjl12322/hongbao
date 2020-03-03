package session

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/kataras/iris/core/errors"
	"sync"
)
type RedisSession struct {
	sessionId string
	pool *redis.Pool
	//设置sessio
	sessionMap map[string]interface{}
	rwlock sync.RWMutex
	//内存中map是否被操作
	flag int
}

const (
	//内容数据无变化
	SessionFlagNone = iota
	//有变化
	SessionFlagModify
)
//构造函数
func NewRedisSession(id string,pool *redis.Pool)*RedisSession{
	s:= &RedisSession{
		sessionId:id,
		sessionMap:make(map[string]interface{},16),
		pool:pool,
		flag:SessionFlagNone,
	}
	return s
}
func (r *RedisSession) Set(key string,value interface{})(err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//设置值
	r.sessionMap[key] = value

	r.flag= SessionFlagModify
	return

}
func (r *RedisSession) Save()(err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//数据没变不需要存
	if r.flag != SessionFlagModify{
		return
	}
	//内存中的redis 进行序列化
	data,err := json.Marshal(r.sessionMap)
	if err !=nil {
		return
	}
	//获取redis 链接
	conn := r.pool.Get()
	_,err = conn.Do("SET",r.sessionId,string(data))
	//改状态
	r.flag = SessionFlagNone
	if err!=nil{
		return
	}

	return

}
func (r *RedisSession) Get(key string)(result interface{},err error) {
	//加锁
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	//判断内存
	if r.flag == SessionFlagNone{
		result,ok:= r.sessionMap[key]
		if !ok{
			err = errors.New("key not exists")
		}
		return

	}


	return

}

func (r *RedisSession)loadFromRedis()(err error){
	conn := r.pool.Get()
	reply,err := conn.Do("GET",r.sessionId)
	if err!=nil{
		return
	}
	//转字符串
	data,err := redis.String(reply,err)
	if err != nil{
		return
	}
	//取到东西反序列化到内存map中
	err = json.Unmarshal([]byte(data),r.sessionMap)
	if err != nil{
		return
	}
	return
}

func (r *RedisSession)Del(key string)(err error)  {
	r.rwlock.Lock()
	defer r.rwlock.Unlock()
	r.flag = SessionFlagModify
	delete(r.sessionMap,key)
	return



}
