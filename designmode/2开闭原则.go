package main

import "fmt"

/*
	一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。
*/

/*
//不使用开闭原则时
//----------原有代码
//Banker银行业务
type Banker struct{}
//存款业务
func (this *Banker) Save(){
	fmt.Println("save...")
}
//转账业务
func (this *Banker) Transfer(){
	fmt.Println("transfer...")
}
//支付业务
func (this* Banker)Pay(){
	fmt.Println("pay...")
}

//-----------新增业务
//股票业务
func (this* Banker)Share(){
	fmt.Println("share...")
}
//...

*/

// ---------使用开闭原则
// 抽象的银行业务员
type AbstractBanker interface {
	//接口的最大的意义就是实现多态的思想
	DoBusi()
}

// 存款的业务员
type SaveBanker struct {
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("save...")
}

// 转账的业务员
type TransferBanker struct {
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("transfer...")
}

// 支付的业务员
type PayBanker struct {
}

func (pb *PayBanker) DoBusi() {
	fmt.Println("pay...")
}

// 实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankerBusiness(banker AbstractBanker) {
	banker.DoBusi()
}

func main() {

	sb := &SaveBanker{}
	sb.DoBusi()
	BankerBusiness(&SaveBanker{})

	tb := &TransferBanker{}
	tb.DoBusi()
	BankerBusiness(&TransferBanker{})

	pb := &PayBanker{}
	pb.DoBusi()
	BankerBusiness(&PayBanker{})
}
