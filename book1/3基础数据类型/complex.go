package main

import (
	"fmt"
	"math/cmplx"
)

/*
	Go语言提供了两种精度的复数类型：complex64和complex128，分别对应float32和float64两种浮点数精度。内置的complex函数用于构建复数，内建的real和imag函数分别返回复数的实部和虚部

	如果一个浮点数面值或一个十进制整数面值后面跟着一个i，它将构成一个复数的虚部，复数的实部是0，3i = 0 + 3i
*/

func main() {

	{
		var x complex128 = complex(1, 2)
		var y complex128 = complex(3, 4)
		fmt.Println(x * y)
		fmt.Println(real(x * y))
		fmt.Println(imag(x * y))

		//如果一个浮点数面值或一个十进制整数面值后面跟着一个i，它将构成一个复数的虚部，复数的实部是0
		fmt.Println(1i * 1i)
	}

	{
		x := 1 + 2i
		y := 3 + 4i
		fmt.Println(x)
		fmt.Println(y)

		fmt.Println(cmplx.Sqrt(-1))
	}
}
