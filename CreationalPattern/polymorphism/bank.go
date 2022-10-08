package main

import "fmt"

type Bank interface {
	Do()
}
type SaveBanker struct{}
type WithdrawBanker struct{}
type TansformerBanker struct{}

//存款业务
func (this *SaveBanker) Do() {
	fmt.Println("进行了 存款业务...")
}

//取钱业务
func (this *WithdrawBanker) Do() {
	fmt.Println("进行了 取钱业务...")
}

//转账业务
func (this *TansformerBanker) Do() {
	fmt.Println("进行了 转账业务...")
}

// 实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankBusiness(b Bank) {
	//通过接口来向下调用，(多态现象)
	b.Do()
}

func main() {
	BankBusiness(&SaveBanker{})
	BankBusiness(&WithdrawBanker{})
	BankBusiness(&TansformerBanker{})
}
