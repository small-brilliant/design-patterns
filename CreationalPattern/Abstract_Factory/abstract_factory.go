package abstractfactory

import "fmt"

type Fruit interface {
	Show()
}

type Factory interface {
	CreateFruit() Fruit
}

// apple
type Apple struct{}

func (a *Apple) Show() {
	fmt.Println("我是苹果")
}

type AppleFactory struct{}

func (af *AppleFactory) CreateFruit() Fruit {
	return new(Apple)
}

// Banana
type Banana struct{}

func (a *Banana) Show() {
	fmt.Println("我是苹果")
}

type BananaFactory struct{}

func (af *BananaFactory) CreateFruit() Fruit {
	return new(Banana)
}

// Banana
type Pear struct{}

func (a *Pear) Show() {
	fmt.Println("我是苹果")
}

type PearFactory struct{}

func (af *PearFactory) CreateFruit() Fruit {
	return new(Pear)
}
