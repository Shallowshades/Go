package main

import "fmt"

func foo1(a string, b int) int {
	fmt.Println("---foo1---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	c := 100
	return c
}

//返回多个返回值，匿名
func foo2(a string, b int) (int, int) {
	fmt.Println("---foo2---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 44, 55
}

//返回多个返回值，有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---foo3---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	//给有名称的返回值变量赋值, 不赋值默认为0
	r1 = 1000
	r2 = 2000

	return

}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---foo4---")
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 1000, 2000
}
func main() {

	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("Cloud", 999)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)

	ret1, ret2 = foo3("Alice", 888)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)

	ret1, ret2 = foo4("Tifa", 777)
	fmt.Println("ret1 = ", ret1, ", ret2 = ", ret2)

}
