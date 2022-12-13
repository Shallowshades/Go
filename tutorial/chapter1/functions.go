package main

import "fmt"

/**
函数

函数可以没有参数或接受多个参数。

在本例中，add 接受两个 int 类型的参数。

1.注意类型在变量名之后。 返回类型后置

当连续两个或多个函数的已命名形参类型相同时，除最后一个类型以外，其它都可以省略。

2.多值返回

函数可以返回任意数量的返回值。

swap 函数返回了两个字符串。

3.命名返回值

Go 的返回值可被命名，它们会被视作定义在函数顶部的变量。

返回值的名称应当具有一定的意义，它可以作为文档使用。

没有参数的 return 语句返回已命名的返回值。也就是 直接 返回。

直接返回语句应当仅用在下面这样的短函数中。在长的函数中它们会影响代码的可读性。
*/

func add(x int, y int) int {
	return x + y
}

func mul(x, y int) int {
	return x * y
}

//多值返回
func swap(x, y string) (string, string) {
	return y, x
}

//返回值命名
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(34, 34))
	fmt.Println(mul(23, 23))
	a, b := swap("hello", "world")
	fmt.Printf("a = %v, b = %v\n", a, b)
	x, y := split(78)
	fmt.Printf("x = %v, y = %v\n", x, y)
}
