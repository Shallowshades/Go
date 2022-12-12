package main

import "fmt"

/*
在Go语言的函数中return语句在底层并不是原子操作，
它分为给返回值赋值和RET指令两步。
而defer语句执行的时机就在返回值赋值操作后，RET指令执行前。
*/

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}
func f5() (x int) {
	defer func(x *int) {
		*x++
	}(&x)
	return 5
}
func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {

	fmt.Println(f1()) //5
	fmt.Println(f2()) //6
	fmt.Println(f3()) //5
	fmt.Println(f4()) //5
	fmt.Println(f5()) //6

	//defer注册要延迟执行的函数时该函数所有的参数都需要确定其值
	x, y := 1, 2
	defer calc("AA", x, calc("A", x, y)) //calc("AA",1,3)
	x = 10
	defer calc("BB", x, calc("B", x, y)) //calc("BB",10,12)
	y = 20
}
