package main

import (
	"fmt"
	"time"
)

//------------------有关task的功能--------

//定义一个任务类型Task
type Task struct {
	f func()error  //一个task里面应该有一个具体的业务，业务名称叫f
}

//创建一个task任务
func NewTask(argf func()error) *Task {
	t:=Task{
		f:argf,
	}
	return &t
}

//task 也需要一个执行业务的方法
func (t *Task)Execute()  {
	t.f()//调用任务中已经绑定好的业务方法
}



//------------------有关协成池pool的功能--------
//定义一个pool协成池的类型
type Pool struct {
	//对外的task入口 EentryChannel
	EentryChannel chan *Task
	//对内部的task队列JobsChannel
	JobsChannel chan *Task
	//携程池最大work数量
	worker_num int
}


//创建pool的函数
func NewPool(cap int) *Pool {
	p:= Pool{
		EentryChannel:make(chan *Task),
		JobsChannel:make(chan *Task),
		worker_num:cap,
	}
	return &p
}

//协成池创建一个work 并让work去工作
func (p *Pool)worker(workerId int)  {
	//worker 具体的工作

	//1 永久的从jobschannel去取任务
	for task := range p.JobsChannel{
		//task 就是worker从job中拿到的任务
		//2一旦取到任务去执行这个任务
		task.Execute()
		fmt.Println("worker Id=  ",workerId)
	}

}

// 让协成池开始真正的工作 协成池启动的方法
func (p *Pool)run() {
	//1 根据worker_num 来创建worker去工作
	for i:=0;i<=p.worker_num;i++{
		go p.worker(i)
	}
	//2 从entrychannel去取任务，将取到的任务发送给jobbschannel
	for task := range p.EentryChannel{
		//一旦有task读取到
		p.JobsChannel <- task
	}
}
func main() {
	//1创建一个任务
	t:= NewTask(func() error {
		fmt.Println(time.Now())
		return nil
	})

	//2 创建一个Pool协成池
	p := NewPool(3)
	task_num := 0
	//3 将这些任务，交给协成池Pool
	go func() {
		for {
			//不断的向p中写入任务t，内个任务就是打印当前时间
			p.EentryChannel <- t
			task_num+=1
			fmt.Println("当前执行了",task_num,"个任务")
		}
	}()

	p.run()

}