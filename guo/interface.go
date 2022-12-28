package main

import "fmt"

type usb interface {
	read()
	write()
}

type player interface {
	playMusic()
}

type video interface {
	playVideo()
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

func (c computer) playMusic() {
	fmt.Printf("c.name: %v\t", c.name)
	fmt.Println("play music...")
}

func (c computer) playVideo() {
	fmt.Printf("c.name: %v\t", c.name)
	fmt.Println("play video...")
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

type cat struct {
	name string
}

func (d *dog) eat(name string) string {
	fmt.Printf("[dog]%s is eating %s\n", d.name, name)
	d.name = "white"
	return "eat well"
}

func (c *cat) eat(name string) string {
	fmt.Printf("[cat]%s is eating %s\n", c.name, name)
	c.name = "white"
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
		p = &dog{"black"} //dog的eat为指针接收者，必须使用指针

		p.eat("chicken") //不使用指针接受的情况，即使接口接受的是指针，函数也无法修改其内部值
		fmt.Println(p)
	}

	//一个类型可以实现多个接口
	//一个接口可以被多个类型实现
	{
		var p player
		var v video

		c := &computer{"XiaoMi"}

		p = c
		p.playMusic()

		v = c
		v.playVideo()

		var pet Pet

		dog := &dog{"black"}
		cat := &cat{"kitty"}

		pet = dog
		pet.eat("bone")

		pet = cat
		pet.eat("fish")
	}

}
