package main

import (
	"fmt"
)

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {

	/*
		类型断言
		类型断言提供了访问接口值底层具体值的方式。

		这种语法和读取一个映射时的相同之处。
	*/
	{
		var i interface{} = "Hello"

		s := i.(string)
		//该语句断言接口值i保存了具体类型T，并将其底层类型为T的值赋予变量t
		//若i并未保存T类型的值，该语句就会触发一个恐慌
		fmt.Println(s)

		f, ok := i.(float64)
		//为了判断一个接口值是否保存了一个特定的类型，类型断言可返回两个值:其底层值以及一个报告断言是否成功的布尔值。
		//若i保存了一个T，那么t将会是其底层值，而ok为true。
		//否则，ok将为false而t将为T类型的零值，程序并不会产生恐慌。
		fmt.Println(f, ok)

		//下面语句会产生panic
		//f = i.(float64)
		//fmt.Println(f)
	}

	/*
		类型选择
		类型选择 是一种按顺序从几个类型断言中选择分支的结构。

		类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）， 它们针对给定接口值所存储的值的类型进行比较。
	*/
	{
		do(21)
		do("Hello")
		do(true)
	}

}
