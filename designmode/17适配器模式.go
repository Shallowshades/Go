package main

import "fmt"

// 适配的目标
type V5 interface {
	Use5V()
}

// 业务类，依赖V5的接口
type Phone struct {
	v V5
}

// Phone有成员要写构造函数
func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("charge...")
	p.v.Use5V()
}

// 被适配的对象，适配者
type V220 struct{}

func (v *V220) Use220V() {
	fmt.Println("use 220v to charge...")
}

// 电源适配器
type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("use adapter to charge...")
	//调用适配者的方法
	a.v220.Use220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}

// 业务层
func main() {

	iphone := NewPhone(NewAdapter(new(V220)))
	iphone.Charge()
}
