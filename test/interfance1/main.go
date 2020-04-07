package main

import "fmt"

type Car interface {
	run()
}
type Driver interface {
	Driver(Car)
}

type BMW struct {

}

func (bmw *BMW)run()  {
	fmt.Println("bmw star")
}
type BC struct {

}
func (bc *BC)run()  {
	fmt.Println("bc star")
}

type Zhang3 struct {

}
type Li4 struct {

}
func (bc *Zhang3)Driver(cat Car)  {
	fmt.Println("zhang3 开始开车")
	cat.run()
}

func (bc *Li4 )Driver(cat Car)  {
	fmt.Println("li4 开始开车")
	cat.run()
}

func main() {
	var bmw Car

	bmw = &BMW{}
	var zhang3 Zhang3

	zhang3.Driver(bmw)


}