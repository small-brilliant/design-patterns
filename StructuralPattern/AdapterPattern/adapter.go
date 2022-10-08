package main

import "fmt"

type Target interface {
	Execute()
}
type Adaptee struct {
}

func (a *Adaptee) specificExecute() {
	fmt.Println("开始充电。。。")
}

type Adapter struct {
	*Adaptee
}

func (a *Adapter) Execute() {
	a.specificExecute()
}
func NewAdapter() *Adapter {
	return &Adapter{}
}

func main() {
	adapter := NewAdapter()
	adapter.specificExecute()
}
