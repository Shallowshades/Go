package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
1.变量的声明

var 语句用于声明一个变量列表，跟函数的参数列表一样，'类型在最后'。
var 语句可以出现在包或函数级别。

2.变量的初始化

变量声明可以包含初始值，每个变量对应一个。
如果初始化值已存在，则可以省略类型；变量会从初始值中获得类型。

3.短变量声明

在函数中，简洁赋值语句 := 可在类型明确的地方代替 var 声明。
函数外的每个语句都必须以关键字开始（var, func 等等），因此 := 结构不能在函数外使用。

4.基本类型

Go 的基本类型有
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // uint8 的别名
rune // int32 的别名 // 表示一个 Unicode 码点
float32 float64
complex64 complex128

同导入语句一样，变量声明也可以“分组”成一个语法块。
int, uint 和 uintptr 在 32 位系统上通常为 32 位宽，在 64 位系统上则为 64 位宽。
当需要一个整数值时应使用 int 类型，除非你有特殊的理由使用固定大小或无符号的整数类型。

5.零值

没有明确初始值的变量声明会被赋予它们的 零值。

零值是：
数值类型为 0，
布尔类型为 false，
字符串为 ""（空字符串）。

6.类型转换

表达式 T(v) 将值 v 转换为类型 T。
与 C 不同的是，Go 在不同类型的项之间赋值时需要显式转换。

7.类型推导

在声明一个变量而不指定其类型时（即使用不带类型的 := 语法或 var = 表达式语法），变量的类型由右值推导得出。
当右值声明了类型时，新变量的类型与其相同
不过当右边包含未指明类型的数值常量时，新变量的类型就可能是 int, float64 或 complex128 了，这取决于常量的精度

8.常量const

常量的声明与变量类似，只不过是使用 const 关键字。
常量可以是字符、字符串、布尔值或数值。
常量不能用 := 语法声明。

9.数值常量

数值常量是高精度的 值。
一个未指定类型的常量由上下文来决定其类型。
（int 类型最大可以存储一个 64 位的整数，有时会更小。）
（int 可以存放最大64位的整数，根据平台不同有时会更少。）
*/

var gA, gB, gC bool

func needInt(x int) int           { return x*10 + 1 }
func needFloat(x float64) float64 { return x * 0.1 }

func main() {

	//1.声明
	{
		var i int
		fmt.Println("i = ", i)
	}

	//2.初始化
	{
		var i, j int = 1, 2
		var Malina, Ranna, Mlisen = true, false, "no!"
		fmt.Printf("i = %v, j = %v\n", i, j)
		fmt.Printf("Malina = %v, Ranna = %v, Mlisen = %v\n", Malina, Ranna, Mlisen)
	}

	//3.短变量声明
	{
		i, j := 1, 2
		Malina, Ranna, Mlisen := true, false, "no!"
		fmt.Printf("i = %v, j = %v\n", i, j)
		fmt.Printf("Malina = %v, Ranna = %v, Mlisen = %v\n", Malina, Ranna, Mlisen)
	}

	//4.类型
	{
		var (
			Cloud  bool       = false
			Alice  uint64     = 1<<64 - 1
			Tifa   complex128 = cmplx.Sqrt(-5 + 12i)
			Barret rune       = (1 << 31) - 1
		)

		fmt.Printf("Type is %T, Value is %v\n", Cloud, Cloud)
		fmt.Printf("Type is %T, Value is %v\n", Alice, Alice)
		fmt.Printf("Type is %T, Value is %v\n", Tifa, Tifa)
		fmt.Printf("Type is %T, Value is %v\n", Barret, Barret)
	}

	//5.零值
	{
		var a int
		var b bool
		var c string
		fmt.Println(a, "---")
		fmt.Println(b, "---")
		fmt.Println(c, "---")
	}

	//6.类型抓换
	{
		var x, y int = 3, 4
		var f float64 = math.Sqrt(float64(x*x + y*y)) //无法直接将不同类型的值赋值
		var z uint = uint(f)
		fmt.Printf("x = %v, y = %v, z = %v\n", x, y, z)
	}

	//7.类型推到
	{
		var i int
		j := i            // j 也是一个 int
		k := 42           // int
		f := 3.142        // float64
		g := 0.867 + 0.5i // complex128
		fmt.Printf("Type of i is %T\n", i)
		fmt.Printf("Type of j is %T\n", j)
		fmt.Printf("Type of k is %T\n", k)
		fmt.Printf("Type of f is %T\n", f)
		fmt.Printf("Type of g is %T\n", g)
	}

	//8.常量
	{
		const World = "世界"
		fmt.Println("Hello", World)
		fmt.Println("Happy", math.Pi, "Day")
		const Truth = true
		fmt.Println("Go rules?", Truth)

	}

	//9.数值常量
	{
		const (
			// 将 1 左移 100 位来创建一个非常大的数字
			// 即这个数的二进制是 1 后面跟着 100 个 0
			Big = 1 << 100
			// 再往右移 99 位，即 Small = 1 << 1，或者说 Small = 2
			Small = Big >> 99
		)
		fmt.Println(needInt(Small))
		fmt.Println(needFloat(Small))
		fmt.Println(needFloat(Big))
	}
}
