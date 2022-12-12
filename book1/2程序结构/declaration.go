package main

import (
	"fmt"
)

/*
命名规则：
	一个名字必须以一个字母（Unicode字母）或下划线开头，后面可以跟任意数量的字母、数字或下划线。大写字母和小写字母是不同的：heapSort和Heapsort是两个不同的名字。

关键字
	break      default       func     interface   select
	case       defer         go       map         struct
	chan       else          goto     package     switch
	const      fallthrough   if       range       type
	continue   for           import   return      var

预定义名字
	内建常量: true false iota nil
	内建类型: int int8 int16 int32 int64
          	uint uint8 uint16 uint32 uint64 uintptr
          	float32 float64 complex128 complex64
          	bool byte rune string error
	内建函数: make len cap new append copy close delete
          	complex real imag
          	panic recover
这些内部预先定义的名字并不是关键字，可以在定义中重新使用它们。在一些特殊的场景中重新定义它们也是有意义的，但是也要注意避免过度而引起语义混乱。

如果一个名字是在函数内部定义，那么它就只在函数内部有效。如果是在函数外部定义，那么将在当前包的所有文件中都可以访问。名字的开头字母的大小写决定了名字在包外的可见性。如果一个名字是大写字母开头的（译注：必须是在函数外部定义的包级名字；包级函数名本身也是包级名字），那么它将是导出的，也就是说可以被外部的包访问，例如fmt包的Printf函数就是导出的，可以在fmt包外部访问。包本身的名字一般总是用小写字母。

推荐使用 驼峰式 命名
而不是优先用下划线分隔

var、const、type和func，分别对应变量、常量、类型和函数实体对象的声明

如果初始化表达式被省略，那么将用零值初始化该变量。 数值类型变量对应的零值是0，布尔类型变量对应的零值是false，字符串类型对应的零值是空字符串，接口或引用类型（包括slice、指针、map、chan和函数）变量对应的零值是nil。数组或结构体等聚合类型对应的零值是每个元素或字段都是对应该类型的零值。

“:=”是一个变量声明语句，而“=”是一个变量赋值操作

用new创建变量和普通变量声明语句方式创建变量没有什么区别，除了不需要声明一个临时变量的名字外，还可以在表达式中使用new(T)。换言之，new函数类似是一种语法糖，而不是一个新的基础概念。

每次调用new函数都是返回一个新的变量的地址

当然也可能有特殊情况：如果两个类型都是空的，也就是说类型的大小是0，例如struct{}和[0]int，有可能有相同的地址（依赖具体的语言实现）（译注：请谨慎使用大小为0的类型，因为如果类型的大小为0的话，可能导致Go语言的自动垃圾回收器有不同的行为，具体请查看runtime.SetFinalizer函数相关文档）。

一个变量的有效周期只取决于是否可达，因此一个循环迭代内部的局部变量的生命周期可能超出其局部作用域。同时，局部变量可能在函数返回之后依然存在。

编译器会自动选择在栈上还是在堆上分配局部变量的存储空间，但可能令人惊讶的是，这个选择并不是由用var还是new声明变量的方式决定的。
*/

// 包一级的各种类型的声明语句的顺序无关紧要
const boilingF = 212.0 //包一级范围

func main() {

	{
		var f = boilingF
		var c = (f - 32) * 5 / 9
		fmt.Printf("boiling point = %g°F or %g°C\n", f, c)
	}

	{
		const freezingF, boilingF = 32.0, 212.0
		fmt.Printf("%g°F = %g°C\n", freezingF, fToC(freezingF))
		fmt.Printf("%g°F = %g°C\n", boilingF, fToC(boilingF))
	}

	{
		var p = func() *int {
			v := 1
			return &v
		}
		fmt.Println(p())
		fmt.Println(p() == p())
	}

	{
		newInt1 := func() *int {
			return new(int)
		}

		newInt2 := func() *int {
			var dummy int
			return &dummy
		}

		x := newInt1()
		y := newInt2()
		fmt.Println(
			x, y,
			// 最后插入的逗号不会导致编译错误，这是Go编译器的一个特性
			// 小括弧另起一行缩进，和大括弧的风格保存一致
		)
	}

	{

	}
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

var global *int

/*
f函数里的x变量必须在堆上分配，因为它在函数退出后依然可以通过包一级的global变量找到，虽然它是在函数内部定义的；用Go语言的术语说，这个x局部变量从函数f中逃逸了。
*/
func f() {
	var x int
	x = 1
	global = &x
}

/*
相反，当g函数返回时，变量*y将是不可达的，也就是说可以马上被回收的。因此，*y并没有从函数g中逃逸，编译器可以选择在栈上分配*y的存储空间（译注：也可以选择在堆上分配，然后由Go语言的GC回收这个变量的内存空间），虽然这里用的是new方式。其实在任何时候，并不需为了编写正确的代码而要考虑变量的逃逸行为，要记住的是，逃逸的变量需要额外分配内存，同时对性能的优化可能会产生细微的影响。
*/
func g() {
	y := new(int)
	*y = 1
}

/*
如果将指向短生命周期对象的指针保存到具有长生命周期的对象中，特别是保存到全局变量时，会阻止对短生命周期对象的垃圾回收（从而可能影响程序的性能）。
*/
