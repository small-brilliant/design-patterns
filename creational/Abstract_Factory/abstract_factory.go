package abstractfactory

import "fmt"

type Lunch interface {
	Cook()
}
type rise struct{}

func (r *rise) Cook() {
	fmt.Println("煮米饭")
}

type Tomato struct{}

func (t *Tomato) Cook() {
	fmt.Println("炒土豆丝")
}

type LunchFactory interface {
	CreateFood() Lunch
	CreateVegetable() Lunch
}
type SimpleLunchFactory struct{}

func (s *SimpleLunchFactory) CreateFood() Lunch {
	return &rise{}
}
func (s *SimpleLunchFactory) CreateVegetable() Lunch {
	return &Tomato{}
}

func NewSimpleLuchFactory() LunchFactory {
	return &SimpleLunchFactory{}
}
