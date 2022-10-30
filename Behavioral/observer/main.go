package main

type Listener interface {
	OnTearchCommint()
}
type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

type Student1 struct {
	Badthing string
}

func (s *Student1) OnTearchCommint() {

}

type Student2 struct {
	Badthing string
}

func (s *Student2) OnTearchCommint() {

}
