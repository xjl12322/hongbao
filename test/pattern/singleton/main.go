package main

import (
	"fmt"
	"sync"
)

//单例模式


//Singleton 是单例模式类
type Singleton struct{}


var singleton *Singleton

var once sync.Once

//GetInstance 用于获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		fmt.Println("生成单利")
		singleton = &Singleton{}
	})

	return singleton
}

func main() {
	//只执行一遍
	GetInstance()
	GetInstance()
	GetInstance()
}





