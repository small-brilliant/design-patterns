package proxy

import "fmt"

type Status uint8

const (
	Stopped Status = iota
	Started
)

type LifeCycle struct {
	name   string
	status Status
}

func (l *LifeCycle) Start() {
	l.status = Started
	fmt.Printf("%s plugin started.\n", l.name)
}

func (l *LifeCycle) Stop() {
	l.status = Stopped
	fmt.Printf("%s plugin stopped.\n", l.name)
}

func (l *LifeCycle) Status() Status {
	return l.status
}

type DbOutput struct {
	LifeCycle
	// 操作数据库的远程代理
	proxy KvDb
}

func (d *DbOutput) Send(msg string) {
	if d.status != Started {
		fmt.Printf("%s is not running, output nothing.\n", d.name)
		return
	}
	record := Record{
		Key:   "db",
		Value: msg,
	}
	reply := false
	err := d.proxy.Save(record, &reply)
	if err != nil || !reply {
		fmt.Println("Save msg to db server failed.")
	}
}
func (d *DbOutput) Init() {
	d.proxy = CreateClient()
	d.name = "db output"
}
