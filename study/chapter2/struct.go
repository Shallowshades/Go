package main

import (
	"fmt"
	"unsafe"
)

/*
	结构体
		一个结构体（struct）就是一组字段（field）。
	结构体字段
		结构体字段使用点号来访问。
	结构体指针
		结构体字段可以通过结构体指针来访问。
		有一个指向结构体的指针 p，可以通过 (*p).X来访问其字段 X。
		语言也允许我们使用隐式间接引用，直接写 p.X 就可以。
	结构体文法
		结构体文法通过直接列出字段的值来新分配一个结构体。（初始化列表）
		使用 Name: 语法可以仅列出 "部分" 字段。（字段名的顺序无关。）（键值对）
		特殊的前缀 & 返回一个指向结构体的指针。
*/

type Person struct {
	Name string
	City string
	Age  int8
}

type student struct {
	name string
	age  int
}

func main() {

	var p1 Person
	p1.Name = "Zack"
	p1.City = "???"
	p1.Age = 18
	fmt.Printf("p1=%v\n", p1)
	fmt.Printf("p1=%#v\n", p1)

	//结构体指针
	var p2 = new(Person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)

	//取结构体地址实例化
	p3 := &Person{}
	fmt.Printf("%T\n", p3)
	fmt.Printf("p3=%#v\n", p3)
	p3.Name = "Cloud"
	p3.Age = 18
	p3.City = "GONGGONGGA"
	fmt.Printf("p3=%#v\n", p3)

	var p4 Person
	fmt.Printf("p4=%#v\n", p4)

	p5 := Person{
		Name: "Alice",
		City: "MiDEJA",
		Age:  18,
	}
	fmt.Printf("p5=%#v\n", p5)

	p6 := &Person{
		Name: "Tifa",
		City: "SEVENHEAVEN",
		Age:  18,
	}
	fmt.Printf("p6=%#v\n", p6)

	/*
		初始化列表
		必须初始化结构体的所有字段。
		初始值的填充顺序必须与字段在结构体中的声明顺序一致。
		该方式不能和键值初始化方式混用。
	*/
	p7 := &Person{
		"Barret",
		"SEVENHEAVEN",
		30,
	}
	fmt.Println("p7 = ", p7)

	var v struct{}
	fmt.Println("size of v = ", unsafe.Sizeof(v))
	var p8 Person
	fmt.Println("size of p8 = ", unsafe.Sizeof(p8))

	//??? slice range 相关问题
	{
		m := make(map[string]*student)
		stus := []student{
			{name: "Cloud", age: 16},
			{name: "Alice", age: 17},
			{name: "Tifa", age: 18},
		}

		for _, stu := range stus {
			m[stu.name] = &stu
			fmt.Println(&stu)
		}

		for k, v := range m {
			fmt.Println(k, "=>", v)
		}
	}
}
