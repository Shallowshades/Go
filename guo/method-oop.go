package main

import "fmt"

/*
	带有receiver的函数称为方法
	function与类型进行了绑定
	从而模拟oop的行为
*/

type Person struct {
	name string
	age  int
}

func (p Person) eat() {
	fmt.Println("person eat...")
}

func (p Person) sleep() {
	fmt.Println("person sleep...")
}

func (p Person) work() {
	fmt.Println("person work...")
}

func main() {
	person := Person{
		name: "Cloud",
		age: 18,
	}

	fmt.Printf("person: %v\n", person)
	person.eat()
	person.work()
	person.sleep()

}
