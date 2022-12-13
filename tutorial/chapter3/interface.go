package main

import (
	"fmt"
	"math"
)

/*
	接口
	接口类型 是由一组方法签名定义的集合。
*/

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型T 实现了接口I，但无需显式声明此事。
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

/*
func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
*/

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {

	/*
		接口类型的变量可以保存任何实现了这些方法的值。
	*/
	{
		var a Abser
		f := MyFloat(-math.Sqrt2)
		v := Vertex{3, 4}

		a = f //a Myfloat实现了Abser
		fmt.Println(a.Abs())

		a = &v //a *Vertex实现了Abser
		fmt.Println(a.Abs())

		//v是一个Vertex,而不是*Vertex，所以没有实现Abser
		//a = v error
	}

	/*
		接口与隐式实现
		类型通过实现一个接口的所有方法来实现该接口。
		既然无需专门显式声明，也就没有“implements”关键字。

		隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。

		因此，也就无需在每一个实现上增加新的接口名称，这样同时也鼓励了明确的接口定义。
	*/
	{
		var i I = &T{"Hello"}
		i.M()
	}

	/*
		接口也是值。它们可以像其它值一样传递。
		接口值可以用作函数的参数或返回值。
		在内部，接口值可以看做包含值和具体类型的元组：(value, type)
		接口值保存了一个具体底层类型的具体值。
		接口值调用方法时会执行其底层类型的同名方法。
	*/
	{
		var i I
		i = &T{"Hello"}
		describe(i)
		i.M()

		i = F(math.Pi)
		describe(i)
		i.M()
	}

	/*
		即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。
		在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它

		注意: 保存了nil具体值的接口其自身并不为nil。

		接口的具体值为nil
		但接口不为nil，有具体的类型
	*/
	{
		var i I

		var t *T
		i = t
		describe(i)
		i.M()

		i = &T{"Hello"}
		describe(i)
		i.M()
	}

	/*
		nil接口值既不保存值也不保存具体类型。
		为 nil 接口调用方法会产生运行时错误，
		因为接口的元组内并未包含能够指明该调用哪个具体方法的类型。
	*/
	{
		var i I
		describe(i)
		//i.M() //runtime error
	}

	/*
		指定了零个方法的接口值被称为 "空接口"
		interface{}空接口可保存任何类型的值。
		（因为每个类型都至少实现了零个方法。）

		空接口被用来处理未知类型的值。
		例如，fmt.Print可接受类型为 interface{}的任意数量的参数。
	*/
	{
		var i interface{}
		describe(i)

		i = 42
		describe(i)

		i = "Hello"
		describe(i)
	}

}
