package main

import "fmt"

/*
	简单工厂模式  + “开闭原则” =    工厂方法模式
*/

// 水果类(抽象接口)
type Fruit interface {
	Show()
}

// 工厂类(抽象接口)
type Factory interface {
	CreateFruit() Fruit
}

// 具体水果类
type Apple struct{}

func (apple *Apple) Show() {
	fmt.Println("Apple...")
}

type Banana struct{}

func (banana *Banana) Show() {
	fmt.Println("Banana...")
}

type Pear struct{}

func (pear *Pear) Show() {
	fmt.Println("Pear...")
}

// +++
type Lemon struct{}

func (lemon *Lemon) Show() {
	fmt.Println("Lemon...")
}

// 具体工厂类
type AppleFactory struct{}

func (fac *AppleFactory) CreateFruit() Fruit {
	return new(Apple)
}

type BananaFactory struct{}

func (fac *BananaFactory) CreateFruit() Fruit {
	return new(Banana)
}

type PearFactory struct{}

func (fac *PearFactory) CreateFruit() Fruit {
	return new(Pear)
}

// +++
type LemonFactory struct{}

func (fac *LemonFactory) CreateFruit() Fruit {
	return new(Lemon)
}

// 业务逻辑
func main() {

	/*
		本案例为了突出根据依赖倒转原则与面向接口编程特性。
		一些变量的定义将使用显示类型声明方式
	*/

	var af Factory = new(AppleFactory)
	var a Fruit = af.CreateFruit()
	a.Show()
	fmt.Printf("%T, %+v\n", af, af)
	fmt.Printf("%T, %+v\n", a, a)

	var bf Factory = new(BananaFactory)
	var b Fruit = bf.CreateFruit()
	b.Show()
	fmt.Printf("%T, %+v\n", bf, bf)
	fmt.Printf("%T, %+v\n", b, b)

	var pf Factory = new(PearFactory)
	var p Fruit = pf.CreateFruit()
	p.Show()
	fmt.Printf("%T, %+v\n", pf, pf)
	fmt.Printf("%T, %+v\n", p, p)

	//+++
	lf := new(LemonFactory)
	l := lf.CreateFruit()
	l.Show()
	fmt.Printf("%T, %+v\n", lf, lf)
	fmt.Printf("%T, %+v\n", l, l)
}
