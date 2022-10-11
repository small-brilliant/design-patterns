package main

import "fmt"

type Attack interface {
	Fight()
}

// 具体技能
type ESkill struct{}

func (e *ESkill) Fight() {
	fmt.Println("使用e技能将敌人击杀")
}

type Hero struct {
	Name   string
	Attack Attack
}

func (h *Hero) Skill() {
	fmt.Println("使用了技能")
	h.Attack.Fight()
}

// 适配者
type PowerOff struct{}

func (p *PowerOff) ShutDown() {
	fmt.Println("电脑即将关机。。。")
}

type Adapter struct {
	PowerOff *PowerOff
}

func (a *Adapter) Fight() {
	a.PowerOff.ShutDown()
}
func NewAdapter(PowerOff *PowerOff) *Adapter {
	return &Adapter{PowerOff}
}
func main() {
	a := NewAdapter(new(PowerOff))
	gl := Hero{"盖伦", a}
	gl.Skill()
}
