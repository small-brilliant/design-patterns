package Factory

import "fmt"

// 需要什么就去工厂拿什么
type Restaurant interface {
	GetFood()
}

// 饭店名字老味道
type LaoWeiDao struct {
}

func (d *LaoWeiDao) GetFood() {
	fmt.Println("烤鸭，上菜")
}

type XinWeiDao struct {
}

func (q *XinWeiDao) GetFood() {
	fmt.Println("烧鸡，上菜")
}

func NewRestaurant(name string) Restaurant {
	switch name {
	case "l":
		return &LaoWeiDao{}
	case "x":
		return &XinWeiDao{}
	}
	return nil
}
