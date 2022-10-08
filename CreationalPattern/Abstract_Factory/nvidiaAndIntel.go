package abstractfactory

import "fmt"

type GPU interface {
	display()
}
type Memory interface {
	storage()
}
type CPU interface {
	calculate()
}
type MotherboardFactory interface {
	CreateGPU() GPU
	CreateMemory() Memory
	CreateCPU() CPU
}

// Intel
type IntelGPU struct{}

func (i *IntelGPU) display() {
	fmt.Println("IntelGPU的display工作")
}

type IntelMemory struct{}

func (i *IntelMemory) storage() {
	fmt.Println("IntelMemory的storage工作")
}

type IntelCPU struct{}

func (i *IntelCPU) calculate() {
	fmt.Println("IntelCPU的calculate工作")
}

type IntelFactory struct{}

func (t *IntelFactory) CreateGPU() GPU {
	return &IntelGPU{}
}
func (t *IntelFactory) CreateMemory() Memory {
	return &IntelMemory{}
}
func (t *IntelFactory) CreateCPU() CPU {
	return &IntelCPU{}
}
func NewIntelFactory() MotherboardFactory {
	return &IntelFactory{}
}

// Nvidia
type NvidiaGPU struct{}

func (i *NvidiaGPU) display() {
	fmt.Println("NvidiaGPU的display工作")
}

type NvidiaMemory struct{}

func (i *NvidiaMemory) storage() {
	fmt.Println("NvidiaMemory的storage工作")
}

type NvidiaCPU struct{}

func (i *NvidiaCPU) calculate() {
	fmt.Println("NvidiaCPU的calculate工作")
}

type NvidiaFactory struct{}

func (t *NvidiaFactory) CreateGPU() GPU {
	return &NvidiaGPU{}
}
func (t *NvidiaFactory) CreateMemory() Memory {
	return &NvidiaMemory{}
}
func (t *NvidiaFactory) CreateCPU() CPU {
	return &NvidiaCPU{}
}
func NewNvidiaFactory() MotherboardFactory {
	return &NvidiaFactory{}
}

// Kingston
type KingstonGPU struct{}

func (i *KingstonGPU) display() {
	fmt.Println("KingstonGPU的display工作")
}

type KingstonMemory struct{}

func (i *KingstonMemory) storage() {
	fmt.Println("KingstonMemory的storage工作")
}

type KingstonCPU struct{}

func (i *KingstonCPU) calculate() {
	fmt.Println("KingstonCPU的calculate工作")
}

type KingstonFactory struct{}

func (t *KingstonFactory) CreateGPU() GPU {
	return &KingstonGPU{}
}
func (t *KingstonFactory) CreateMemory() Memory {
	return &KingstonMemory{}
}
func (t *KingstonFactory) CreateCPU() CPU {
	return &KingstonCPU{}
}
func NewKingstonFactory() MotherboardFactory {
	return &KingstonFactory{}
}
