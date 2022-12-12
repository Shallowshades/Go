package main

import "fmt"

/*
type 类型名字 底层类型
	一个类型声明语句创建了一个新的类型名称，和现有类型具有相同的底层结构。
	新命名的类型提供了一个方法，用来分隔不同概念的类型，这样即使它们底层类型相同也是不兼容的。

	类型声明语句一般出现在包一级，因此如果新创建的类型名字的首字符大写，则在包外部也可以使用。
	对于中文汉字，Unicode标志都作为小写字母处理，因此中文的命名默认不能导出

	对于每一个类型T，都有一个对应的类型转换操作T(x)，用于将x转为T类型
	如果T是指针类型，可能会需要用小括弧包装T，比如(*int)(0)
	只有当两个类型的底层基础类型相同时，才允许这种转型操作，或者是两者都是指向相同底层结构的指针类型，这些转换只改变类型而不会影响值本身

	数值类型之间的转型也是允许的，并且在字符串和一些特定类型的slice之间也是可以转换的
	这类转换可能改变值的表现。
	例如，将一个浮点数转为整数将丢弃小数部分，将一个字符串转为[]byte类型的slice将拷贝一个字符串数据的副本。
	在任何情况下，运行时不会发生转换失败的错误
	错误只会发生在编译阶段

	底层数据类型决定了内部结构和表达方式，也决定是否可以像底层类型一样对内置运算符的支持。这意味着，Celsius和Fahrenheit类型的算术运算行为和底层的float64类型是一样的
*/

// 虽然有着相同的底层类型float64,但是它们是不同的数据类型
// 因此它们不可以被相互比较或混在一个表达式运算
type Celsius float64    //摄氏温度
type Fahrenheit float64 //华氏温度

const (
	AbsoluteZeroC Celsius = -273.15 //绝对零度
	FreezingC     Celsius = 0       //冰点温度
	BoilingC      Celsius = 100     //沸点温度
)

// Celsius(t)和Fahrenheit(t)是类型转换操作，它们并不是函数调用。类型转换不会改变值本身，但是会使它们的语义发生变化
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// 此处使用指针接收，%s将无法匹配输出， warning
// func (c *Celsius) String() string {
// 	return fmt.Sprintf("%g°C", c)
// }

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

func main() {

	{
		fmt.Printf("%g°C\n", BoilingC-FreezingC)
		boilingF := CToF(BoilingC)
		fmt.Printf("%g°F\n", boilingF-CToF(FreezingC))
		//fmt.Printf("%g°F\n", boilingF-FreezingC) compile error: type mismatch
	}

	{
		var c Celsius
		var f Fahrenheit
		fmt.Println(c == 0)
		fmt.Println(f >= 0)
		//fmt.Println(c == f) compile error: type mismatch

		//尽管看起来像函数调用，但是Celsius(f)是类型转换操作，它并不会改变值，仅仅是改变值的类型而已。测试为真的原因是因为c和f都是零值。
		fmt.Println(c == Celsius(f))
	}

	{
		c := FToC(212.0)

		fmt.Println("c = ", c)
		fmt.Println("c.String()", c.String())
		fmt.Printf("c = %v\n", c)
		fmt.Printf("c = %s\n", c)
		fmt.Println(c)
		fmt.Printf("%g\n", c)
		fmt.Println(float64(c))
	}
}
