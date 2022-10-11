package main

import "fmt"

type Docker struct{}

func (d *Docker) treatEye() {
	fmt.Println("医生治疗眼睛")
}
func (d *Docker) treatNose() {
	fmt.Println("医生治疗鼻子")
}

type Command interface {
	Treat()
}
type CommandTreatEye struct {
	Docker *Docker
}

func (c *CommandTreatEye) Treat() {
	c.Docker.treatEye()
}

type CommandTreatNose struct {
	Docker *Docker
}

func (c *CommandTreatNose) Treat() {
	c.Docker.treatNose()
}

type Nurse struct {
	CmdList []Command
}

func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat()
	}
}
func main() {
	Docker := &Docker{}
	CmdEye := &CommandTreatEye{Docker}
	CmdNose := &CommandTreatNose{Docker}

	nurse := new(Nurse)
	nurse.CmdList = append(nurse.CmdList, CmdEye)
	nurse.CmdList = append(nurse.CmdList, CmdNose)

	nurse.Notify()
}
