package demo

import "fmt"

type Phone interface {
	Call()
}
type NokiaPhone struct {
}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("我是诺基亚")
}

type IPhone struct {
}

func (iPhone IPhone) Call() {
	fmt.Println("I am iPhone, I can call you!")
}
