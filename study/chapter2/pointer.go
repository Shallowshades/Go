package main

import "fmt"

/*
	Go 拥有指针。指针保存了值的内存地址。
	类型 *T 是指向 T 类型值的指针。其零值为 nil。
	var p *int
	& 操作符会生成一个指向其操作数的指针。
	i := 42
	p = &i
	* 操作符表示指针指向的底层值。
	fmt.Println(*p) // 通过指针 p 读取 i
	*p = 21         // 通过指针 p 设置 i
	这也就是通常所说的“间接引用”或“重定向”。

	与 C 不同，Go 没有指针运算。
*/

func main() {

	var p *int
	i := 42
	p = &i

	//查看地址
	fmt.Printf("%v\n", &i)
	fmt.Printf("%p\n", &i)
	fmt.Printf("%v\n", p)
	fmt.Printf("%p\n", p)
}
