package main

import "fmt"

type Phone interface {
	Show()
}

// 抽象的装饰器，该类本应该是interface，
// 由于Golang语法的interface不可以有成员属性，所以定义为struct来重写show方法
type Decorator struct {
	phone Phone
}
type Decorator2 struct {
	phone Phone
}

func (a *Decorator) Show()  {}
func (a *Decorator2) Show() {}

// 实现层
type Huawei struct{}

func (a *Huawei) Show() {
	fmt.Println("华为手机")
}

type Mi struct{}

func (a *Mi) Show() {
	fmt.Println("小米手机")
}

// 具体的装饰器
type MoDecorator struct {
	Decorator
}

// 具体的装饰器
func (md *MoDecorator) Show() {
	md.phone.Show()
	fmt.Println("贴膜")
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

// 手机壳
type KDecorator struct {
	Decorator
}

func (md *KDecorator) Show() {
	md.phone.Show()
	fmt.Println("手机壳")
}

func NewKDecorator(phone Phone) Phone {
	return &KDecorator{Decorator{phone}}
}

type SDecorator struct {
	Decorator2
}

func (md *SDecorator) Show() {
	md.phone.Show()
	fmt.Println("给手机壳添加绳子")
}

func NewSDecorator(phone Phone) Phone {
	return &SDecorator{Decorator2{phone}}
}
func main() {
	hw := &Huawei{}
	// hw.Show()
	//
	md := NewMoDecorator(hw)
	hkd := NewKDecorator(hw)
	sk := NewSDecorator(hw)
	md.Show()
	hkd.Show()
	sk.Show()
	kd := NewKDecorator(&Mi{})
	kd.Show()
}
