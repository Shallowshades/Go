package main

import "fmt"

/*
	MakeCoffee是template的子类
	MakeCoffee实现了Beverage的接口
	那 template 和 Beverage 是什么关系？？？
*/

// 抽象类，制作饮料，包裹一个模板的全部实现步奏
type Beverage interface {
	BoilWater() //煮开水
	Brew()      //冲泡
	PourInCup() //倒入杯中
	AddThings() //添加佐料

	WantAddThings() bool //是否加入佐料的Hook
}

// 封装一套流程模板,让具体的制作流程继承且实现
type template struct {
	b Beverage //模板包含一个饮料的接口
}

// 封装的固定模板
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}

	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()

	//让子类可以重写该方法来决定是否执行下面动作
	if t.b.WantAddThings() == true {
		t.b.AddThings()
	}
}

// 具体的模板子类 制作咖啡
type MakeCoffee struct {
	template //继承模板
}

func NewMakeCoffee() *MakeCoffee {
	makeCoffee := new(MakeCoffee)
	//b Beverage，是MakeCoffee的接口，这里需要给接口赋值，指向具体的子类来触发b全部接口方法的多态特性
	makeCoffee.b = makeCoffee
	return makeCoffee
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("boil water to 100°C...")
}

func (mc *MakeCoffee) Brew() {
	fmt.Println("brew coffee bean...")
}

func (mc *MakeCoffee) PourInCup() {
	fmt.Println("pour in cup...")
}

func (mc *MakeCoffee) AddThings() {
	fmt.Println("add milk and sweety...")
}

func (mc *MakeCoffee) WantAddThings() bool {
	return true //启动Hook条件
}

// 具体模板的子类 制作茶
type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("boil water to 100°C...")
}

func (mt *MakeTea) Brew() {
	fmt.Println("brew tea leaves...")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("pour in cup...")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("add lemon...")
}

func (mt *MakeTea) WantAddThings() bool {
	return false //启动Hook条件
}

func main() {

	//1.制作已被咖啡
	makeCoffee := NewMakeCoffee()
	makeCoffee.MakeBeverage()

	fmt.Println("---------------")

	//2.制作茶
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}
