package main

import "fmt"

//抽象层

//组装电脑
type Card interface { //显卡
	Display()
}

type Memory interface {//内存
	Storage()
}

type CPU interface {//cpu
	Calculate()
}
type Computer struct {  //电脑
	cpu CPU
	mem Memory
	card Card
}

//初始化电脑对象的方法
func NewComputer(cpu CPU,mem Memory,card Card) *Computer {
	return &Computer{
		cpu:cpu,
		mem:mem,
		card:card,

	}
}
func (this *Computer)work()  {
	this.cpu.Calculate()
	this.mem.Storage()
	this.card.Display()

}

//实现层
type IntelCPU struct {
	CPU
}
func (this *IntelCPU) Calculate()  {
	fmt.Println("intel cpu 开始计算了")
}


type IntelCard struct {
    Card
}
func (this *IntelCard)Display() {
	fmt.Println("intel 显卡 开始显示了")
}


type IntelMemory struct {
	Memory
}
func (this *IntelMemory)Storage() {
	fmt.Println("intel 内存 开始储存了")
}

type ZapaiCPU struct {
	CPU
}

func (za *ZapaiCPU)Calculate()  {
	fmt.Println("杂牌  cup 开始计算了")
}

type ZapaiCard struct {
	Card
}

func (za *ZapaiCard)Display()  {
	fmt.Println("杂牌  显卡 开始计算了")
}
type ZapaiMemory struct {
	Card
}

func (za *ZapaiMemory)Storage()  {
	fmt.Println("杂牌  内存开始计算了")
}


func main() {
	//组装一台intel
	com := NewComputer(&IntelCPU{},&IntelMemory{},&IntelCard{})
	//intel 电脑开始工作
	com.work()

	//杂牌电脑
	com2 := NewComputer(&ZapaiCPU{},&ZapaiMemory{},&ZapaiCard{})
	com2.work()

}