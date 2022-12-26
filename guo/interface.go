package main

import "fmt"

type usb interface {
	read()
	write()
}

type computer struct {
	name string
}

func (c computer) read() {
	fmt.Printf("c.name: %v\t", c.name)
	fmt.Println("read...")
}

func (c computer) write() {
	fmt.Printf("c.name: %v\t", c.name)
	fmt.Println("write...")
}

type mobile struct {
	model string
}

func (m *mobile) read() {
	fmt.Printf("m.model: %v\t", m.model)
	fmt.Println("read...")
}

func (m *mobile) write() {
	fmt.Printf("m.model: %v\t", m.model)
	fmt.Println("write...")
}

type OpenClose interface {
	open()
	close()
}

type Door struct{}

func (d *Door) open() {
	fmt.Println("open the door...")
}

type Pet interface {
	eat(string) string
}

type dog struct {
	name string
}

func (d *dog) eat(name string) string {
	fmt.Printf("%s is eating %s", d.name, name)
	d.name = "white"
	return "eat well"
}

func main() {

	//interface
	{
		var u usb

		//computer 采用值接受，可以接受值或者指针
		u = &computer{name: "lenveno"}
		u.read()
		u.write()

		//mobile 采用指针接受，只能接受指针
		u = &mobile{model: "5G"}
		u.read()
		u.write()
	}

	//必须实现接口的所有函数声明
	{
		//var oc OpenClose

		//oc = &Door{} 未实现close
	}

	//接口的接收者
	{
		var p Pet
		p = &dog{"black"} //dog的eat为指针接收者，使用指针
		fmt.Printf("p.eat(): %v\n", p.eat("chicken"))
		fmt.Println(p)
	}

}
