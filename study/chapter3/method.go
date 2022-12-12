package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {

	/*
		Go 没有类。不过可以为结构体类型定义方法。
		方法就是一类带特殊的 接收者 参数的函数。
		方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。
	*/
	{
		v := Vertex{3, 4}
		fmt.Println(v.Abs())
	}

	/*
		方法只是个带接收者参数的函数。
	*/
	{
		v := &Vertex{4, 5}
		len := func(v *Vertex) float64 {
			return math.Sqrt(v.X*v.X + v.Y*v.Y)
		}(v)
		fmt.Println(len)
	}

	/*
		也可以为非结构体类型声明方法。
		你只能为在同一包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。
		接收者的类型定义和方法声明必须在同一包内；不能为内建类型声明方法。
	*/
	{
		f := MyFloat(-math.Sqrt2)
		fmt.Println(f.Abs())
	}

	/*
		为指针接收者声明方法。
		这意味着对于某类型 T，接收者的类型可以用 *T 的文法。
		T 不能是像 *int 这样的指针
		指针接收者的方法可以修改接收者指向的值
		使用值接收者，那么方法会对原始值的副本进行操作
	*/
	{
		v := Vertex{3, 4}
		v.Scale(10)
		fmt.Println(v.Abs())
	}

	//
	{
		v := Vertex{3, 4}
		Scale(&v, 10)
		fmt.Println(Abs(v))
	}

	/*
		方法与指针重定向
		带指针参数的函数必须接受一个指针
		以指针为接收者的方法被调用时，接收者既能为值又能为指针
	*/

	{
		v := Vertex{3, 4}
		p := &Vertex{3, 4}

		// func (v *Vertex) Scale () float64{ ... }
		v.Scale(2)
		//方法为指针接收者时，亦可值传递
		//Go 会将语句 v.Scale(5) 解释为 (&v).Scale(5)
		p.Scale(3)
		//正常传递指针

		//func Scale(v *Vertex, f float64) float64{ ... }
		//Scale(v, 10)
		//error 函数形参为指针不能值传递
		Scale(p, 10)
		//正常传递指针

		fmt.Println(v, p)
	}

	/*
		方法与指针重定向
		接受一个值作为参数的函数必须接受一个指定类型的值
		以值为接收者的方法被调用时，接收者既能为值又能为指针
	*/
	{
		v := Vertex{3, 4}
		p := &Vertex{3, 4}

		//func Abs(v Vertex)float64{ ... }
		fmt.Println(Abs(v))
		//形参为值类型，传值正常
		//fmt.Println(Abs(p))
		//形参为值类型，传指针error

		//func (v Vertex)Abs()float64{ ... }
		fmt.Println(v.Abs())
		fmt.Println(p.Abs())
		//以值为接受者的方法，值和指针都能调用
		//p.Abs() 会被解释为 (*p).Abs()
	}

	/*
		选择值或指针作为接收者

		使用指针接收者的原因有二：
		1.方法能够修改其接收者指向的值。
		2.避免在每次调用方法时复制该值。若值的类型为大型结构体时会更加高效。


		通常来说，所有给定类型的方法都应该有值或指针接收者，但并不应该二者混用。
	*/

}
