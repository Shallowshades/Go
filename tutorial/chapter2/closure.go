package main

import (
	"fmt"
	"strings"
)

/*
闭包指的是一个函数和与其相关的引用环境组合而成的实体。
简单来说，闭包=函数+引用环境。

封闭在函数体内的函数
*/

func addr() func(int) int {
	var x int //此值在闭包创建时初始化
	return func(y int) int {
		x += y
		return x
	}
}

func addr2(x int) func(int) int {
	return func(y int) int {
		x += y
		return x
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		//没有后缀的情况，加上后缀
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}
	return add, sub
}

func main() {

	/*
		变量f是一个函数并且它引用了其外部作用域中的x变量，此时f就是一个闭包。
		在f的生命周期内，变量x也一直有效。
	*/
	{
		var f = addr()
		fmt.Println(f(10))
		fmt.Println(f(20))
		fmt.Println(f(30))

		f1 := addr()
		fmt.Println(f1(30))
		fmt.Println(f1(40))
		fmt.Println(f1(50))
	}

	{
		var f = addr2(10)  //10
		fmt.Println(f(10)) //20
		fmt.Println(f(20)) //40
		fmt.Println(f(30)) //70

		f1 := addr2(40)     //40
		fmt.Println(f1(50)) //90
		fmt.Println(f1(60)) //150
		fmt.Println(f1(70)) //220
	}

	{
		jpgFunc := makeSuffixFunc(".jpg")
		goFunc := makeSuffixFunc(".go")
		txtFunc := makeSuffixFunc(".txt")
		fmt.Println(jpgFunc("test"))
		fmt.Println(goFunc("test"))
		fmt.Println(txtFunc("test"))
	}

	{
		f1, f2 := calc(10)
		fmt.Println(f1(1), f2(2))
		fmt.Println(f1(3), f2(4))
		fmt.Println(f1(5), f2(6))
	}

}
