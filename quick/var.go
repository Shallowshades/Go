package main

import "fmt"

//声明全局变量,方法1，2，3
var gA int = 100
var gB = 200

//声明全局变量，方法4
// := 只能够用在函数体内来声明
//gC := 300

func main() {

	//前三种使用场景无区别
	//1.声明一个变量 默认值是0
	var a int
	fmt.Println("a = ", a)
	fmt.Printf("type of a = %T\n", a)

	//2.声明一个变量并初始化
	var b int = 100
	fmt.Println("b = ", b)
	fmt.Printf("type of b = %T\n", b)

	var bb string = "abcd"
	fmt.Printf("bb = %s, type of bb = %T\n", bb, bb)

	//3.在初始化的时候，省略数据类型，通过值自动匹配当前变量的数据类型
	var c = 100
	fmt.Println("c = ", c)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("cc = %s, type of cc = %T\n", cc, cc)

	//4.（常用的方法）省略var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	f := "abcd"
	fmt.Println("f = ", f)
	fmt.Printf("type of f = %T\n", f)

	g := 3.14
	fmt.Println("g = ", g)
	fmt.Printf("type of g = %T\n", g)

	//--------
	fmt.Println("gA = ", gA, ", gB = ", gB)
	//fmt.Println("gC = ", gC)

	//声明多个变量
	var xx, yy int = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	kk, ll := 100, "Alice"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	//多行多变量声明
	var (
		vv int    = 100
		uu string = "Tifa"
		ww bool   = true
	)
	fmt.Println("vv = ", vv, ", uu = ", uu, ", ww = ", ww)
}
