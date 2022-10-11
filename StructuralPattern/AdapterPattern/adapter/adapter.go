package main

import "fmt"

// 适配目标
type V5 interface {
	UseV5()
}

// 被适配的角色
type V220 struct{}

func (v *V220) User220v() {
	fmt.Println("使用220V的电压")
}

// 适配器
type Adapter struct {
	V220 *V220
}

func (a *Adapter) UseV5() {
	a.V220.User220v()
	fmt.Println("使用适配器进行220V到5V")
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

// 业务
type Phone struct {
	v V5
}

func (p *Phone) charge() {
	fmt.Println("chargin 5v")
	p.v.UseV5()

}
func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func main() {
	adapter := NewAdapter(new(V220))
	phone := NewPhone(adapter)
	phone.charge()
}
