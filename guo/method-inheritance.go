package main

import "fmt"

type Animal struct {
	name string
	age  int
}

func (a Animal) eat() {
	fmt.Println("animal eat...")
}

func (a Animal) sleep() {
	fmt.Println("animal sleep...")
}

type Dog struct {
	a     Animal
	color string
}

type Cat struct {
	Animal
	ability string
}

func main() {
	dog := Dog{
		a:     Animal{"cloud", 8},
		color: "black",
	}
	fmt.Printf("dog: %v\n", dog)
	//由于a Animal是个显示变量，所以要显示调用
	dog.a.eat()
	dog.a.sleep()
	fmt.Printf("dog.color: %v\n", dog.color)
	fmt.Printf("dog.a.age: %v\n", dog.a.age)

	cat := Cat{Animal{"Alice", 8}, "play"}
	fmt.Printf("cat: %v\n", cat)
	//Animal在Cat中是一个匿名变量，所以可以隐式调用
	cat.eat()
	cat.sleep()
	fmt.Printf("cat.ability: %v\n", cat.ability)
	fmt.Printf("cat.age: %v\n", cat.age)
}
