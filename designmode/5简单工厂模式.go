package main

import "fmt"

/*
	业务逻辑层 ---> 工厂模块 ---> 基础类模块
*/

// ==========抽象层============

type Fruit interface {
	Show()
}

// ==========基础类模块============

type Apple struct {
	Fruit ////为了易于理解显示继承(此行可以省略)
}

func (apple *Apple) Show() {
	fmt.Println("Apple...")
}

type Banana struct {
	Fruit
}

func (banana *Banana) Show() {
	fmt.Println("Banana...")
}

type Pear struct {
	Fruit
}

func (pear *Pear) Show() {
	fmt.Println("Pear...")
}

// ==========工厂模块============

type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruit {
	var fruit Fruit

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}

	return fruit
}

// ==========业务逻辑层============

func main() {

	factory := new(Factory)

	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
