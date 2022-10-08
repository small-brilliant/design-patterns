package abstractfactory

import "testing"

func TestNewSimpleLuchFactory(t *testing.T) {
	factory := NewSimpleLuchFactory()
	factory.CreateFood().Cook()
	factory.CreateVegetable().Cook()
}
