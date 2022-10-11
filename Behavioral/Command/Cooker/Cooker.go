package main

import "fmt"

type Cooker struct{}

func (c *Cooker) MakeChicken() {
	fmt.Println("烤串师傅烤了鸡肉串儿")
}
func (c *Cooker) MakeChuaner() {
	fmt.Println("烤串师傅烤了羊肉串儿")
}

type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeChicken()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

type Menu interface {
	Make()
}

type WaiterMM struct {
	MenuList []Menu
}

func (w *WaiterMM) Notify() {
	if w.MenuList == nil {
		return
	}

	for _, cmd := range w.MenuList {
		cmd.Make()
	}
}
func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker}
	cmdChuaner := CommandCookChuaner{cooker}

	mm := new(WaiterMM)
	mm.MenuList = append(mm.MenuList, &cmdChicken)
	mm.MenuList = append(mm.MenuList, &cmdChuaner)

	mm.Notify()
}
