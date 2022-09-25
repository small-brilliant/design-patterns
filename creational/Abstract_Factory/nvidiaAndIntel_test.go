package abstractfactory

import (
	"fmt"
	"testing"
)

func TestMotherboardFactory(t *testing.T) {
	// 创建工厂
	intelF := NewIntelFactory()
	NvidiaF := NewNvidiaFactory()
	kingstonF := NewKingstonFactory()
	// Intel的CPU，Intel的显卡，Intel的内存
	fmt.Println("第一台电脑：")
	intelF.CreateCPU().calculate()
	intelF.CreateGPU().display()
	intelF.CreateMemory().storage()
	fmt.Println("第二台电脑：")
	intelF.CreateCPU().calculate()
	NvidiaF.CreateGPU().display()
	kingstonF.CreateMemory().storage()
}
