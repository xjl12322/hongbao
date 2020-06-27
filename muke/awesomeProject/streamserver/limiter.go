package main

import "log"
//bucket 算法
type ConnLimiter struct {
	concurrentConn int
	bucket chan int
}

func NewConnLimiter(cc int)*ConnLimiter  {
	return &ConnLimiter{
		concurrentConn:cc,
		bucket:make(chan int,cc),  //带缓存
	}
}


//留空机制
func (cl *ConnLimiter) GetConn()bool  {
	if len(cl.bucket) >=cl.concurrentConn{
		log.Printf("Reached the rate limitation")
		return false
	}
	cl.bucket<- 1
	return true

}
func (cl *ConnLimiter) ReleaseConn()  {
	c :=  <- cl.bucket
	log.Printf("New connction coming:%d",c)
}

