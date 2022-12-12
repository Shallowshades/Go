package main

import (
	"fmt"
)

// const定义枚举类型
const (
	//可以在const()添加一个关键字iota，每行iota都会累加1，第一行默认是0
	//iota只能配合const使用
	Cloud = 10 * iota
	Alice
	Tifa
	Barret
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f

	g, h = iota * 2, iota * 3
	i, j
)

func main() {
	//常量(只读), 不允许修改
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("Cloud = ", Cloud)
	fmt.Println("Alice = ", Alice)
	fmt.Println("Tifa = ", Tifa)
	fmt.Println("Barret = ", Barret)

	fmt.Println("a = ", a, ", b = ", b)
	fmt.Println("c = ", c, ", d = ", d)
	fmt.Println("e = ", e, ", f = ", f)
	fmt.Println("g = ", g, ", h = ", h)
	fmt.Println("i = ", i, ", j = ", j)

}
