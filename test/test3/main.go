package main

import "runtime"
import "fmt"

func main() {
	pc,file,line,ok := runtime.Caller(0)
	fmt.Println(pc)
	fmt.Println(file)
	fmt.Println(line)
	fmt.Println(ok)

}