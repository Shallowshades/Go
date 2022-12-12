package main

import "fmt"

type SubSystemA struct{}

func (sa *SubSystemA) MethodA() {
	fmt.Println("A ...")
}

type SubSystemB struct{}

func (sa *SubSystemB) MethodB() {
	fmt.Println("B ...")
}

type SubSystemC struct{}

func (sa *SubSystemC) MethodC() {
	fmt.Println("C ...")
}

type SubSystemD struct{}

func (sa *SubSystemD) MethodD() {
	fmt.Println("D ...")
}

// 外观模式，提供一个外观类，简化成一个简单的接口供使用
type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
	f.c.MethodC()
	f.d.MethodD()
}

func main() {

	//不使用外观模式时
	sa := new(SubSystemA)
	sa.MethodA()
	sb := new(SubSystemB)
	sb.MethodB()

	//使用外观模式
	fmt.Println("-------------")
	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}
	//调用外观包裹方法
	f.MethodOne()

}
