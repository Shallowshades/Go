package main

import (
	"fmt"
	"io"
	"os"
)

/*
1.defer

defer 语句会将函数推迟到外层函数返回之后执行。
推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

2.defer 栈

推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。
*/

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close() //通过defer将src的关闭放于函数结束后，即使dst打开失败，也能关闭

	dst, err := os.Create(dstName)
	if err != nil { //如果在此处打开失败，则src在原始写法中未关闭
		return
	}
	defer dst.Close()

	written, err = io.Copy(dst, src)

	return
}

func a() {
	i := 0
	defer fmt.Println("i = ", i) // i = 0,在函数返回时打印0
	i++
	return
}

func b() {
	i := new(int)
	*i = 0
	defer fmt.Println("b() *i = ", *i) // i = 0,在函数返回时打印0,压栈时已经将所有的表达式计算完毕
	*i++
	return
}

/*
调用顺序是，先return 再defer
但此时到defer并未返回上一级函数，调用完毕才返回上一级函数
*/
func c() (i int) {
	defer func() {
		fmt.Println("c() i = ", i)
		i++
	}() //在函数结束时，i++,然后返回给上一级函数调用
	return 1 //此时可以理解为将返回值赋值给i
}

func main() {

	defer fmt.Println("world")
	fmt.Println("hello")

	{
		fmt.Println("counting")
		for i := 0; i < 10; i++ {
			defer fmt.Println(i)
		}
		fmt.Println("done")
	}

	{
		CopyFile("2.txt", "1.txt")
	}

	{
		fmt.Println("--------------")
		a()
		b()
		fmt.Println("c() i = ", c())
	}

	fmt.Println("--------------")
}
