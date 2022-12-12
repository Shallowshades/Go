package main

import "fmt"

func func1() {
	fmt.Println("A")
}

func func2() {
	fmt.Println("B")
}

func func3() {
	fmt.Println("C")
}

func func_defer() {
	defer func1()
	defer func2()
	defer func3()
}

func deferFunc() int {
	fmt.Println("defer func called...")
	return 0
}

func returnFunc() int {
	fmt.Println("return func called...")
	return 0
}

/*
return 先执行
*/
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
} //走到这里才出栈

func main() {
	//写入defer关键字，位置随意
	//当前函数结束之前触发
	/*
		执行顺序是先入栈后出栈
	*/

	fmt.Println("------------")
	defer fmt.Println("main end1")
	defer fmt.Println("main end2")
	fmt.Println("main::Hello go 1")
	fmt.Println("main::Hello go 2")

	fmt.Println("------------")
	func_defer()

	fmt.Println("------------")
	returnAndDefer()

	fmt.Println("----------")
}
