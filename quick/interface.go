package main

import "fmt"

//interface{}是万能数据接口
func myFunc(arg interface{}) {
	fmt.Println("myFunc is called...")
	fmt.Println(arg)

	//给interface提供“类型断言”的机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not string")
	} else {
		fmt.Println("arg is string type")
		fmt.Printf("value type is %T\n", value)
	}

}

type Book struct {
	author string
}

func main() {
	book := Book{"Zhangsan"}
	myFunc(book)
	myFunc(100)
	myFunc("abc")
	myFunc(3.14)
}
