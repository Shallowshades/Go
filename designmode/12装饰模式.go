package main

import "fmt"

//===========抽象层===========

// 抽象的构件
type Phone interface {
	Show() //构件的功能
}

// 装饰器的抽象类
// （该类本应该为interface，但是Golang interface语法不可以有成员属性）
type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {}

//===========实现层===========

// 具体的构件
type HuaWei struct{}

func (hw *HuaWei) Show() {
	fmt.Println("show you the Huawei phone")
}

type XiaoMi struct{}

func (xm *XiaoMi) Show() {
	fmt.Println("show you the xiaomi phone")
}

// 具体的装饰器类
type MoDecorator struct {
	Decorator //继承抽象的装饰器类（主要继承Phone成员属性）
}

func (md *MoDecorator) Show() {
	md.phone.Show()                       //调用被装饰构件的原方法
	fmt.Println("phone with facial mode") //装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("phone with plastic ke")
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}

//===========业务层===========

func main() {

	var huawei Phone
	huawei = new(HuaWei)
	huawei.Show()

	fmt.Println("--------------")
	var moHuaWei Phone
	moHuaWei = NewKeDecorator(huawei)
	moHuaWei.Show()

	fmt.Println("--------------")
	var keHuaWei Phone
	keHuaWei = NewKeDecorator(huawei)
	keHuaWei.Show()

	fmt.Println("--------------")
	var xiaomi Phone
	xiaomi = new(XiaoMi)
	var keMoXiaoMi Phone
	keMoXiaoMi = NewKeDecorator(NewMoDecorator(xiaomi))
	keMoXiaoMi.Show()

}
