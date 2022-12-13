package main

import (
	"fmt"
	"math"
	"runtime"
	"time"
)

/*
1.for

Go 只有一种循环结构：for 循环。
基本的 for 循环由三部分组成，它们用分号隔开：
初始化语句：在第一次迭代前执行
条件表达式：在每次迭代前求值
后置语句：在每次迭代的结尾执行
初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。
一旦条件表达式的布尔值为 false，循环迭代就会终止。

注意：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。

2.初始化语句和后置语句是可选的。

3.for 是 Go 中的 “while”

此时你可以去掉分号，因为 C 的 while 在 Go 中叫做 for。

4.无限循环

如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。

5.if

Go 的 if 语句与 for 循环类似，表达式外无需小括号 ( ) ，而大括号 { } 则是必须的。

6.if 的简短语句

同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。
该语句声明的变量作用域仅在 if 之内。
在 if 的简短语句中声明的变量同样可以在任何对应的 else 块中使用。

7.switch

switch 是编写一连串 if - else 语句的简便方法。它运行第一个值等于条件表达式的 case 语句。
Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。
实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。
除非以 fallthrough 语句结束，否则分支会自动终止。
Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。

8.switch 的求值顺序

switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

9.没有条件的 switch

没有条件的 switch 同 switch true 一样。
这种形式能将一长串 if-then-else 写得更加清晰。
*/

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	} else {
		return fmt.Sprint(math.Sqrt(x))
	}
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

// 牛顿法求平方根
func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func main() {

	//1.for
	{
		sum := 0
		for i := 0; i < 5; i++ {
			sum += i
		}
		fmt.Printf("sum = %v\n", sum)
	}

	//2.初始化语句和后置语句是可选的。
	{
		sum := 1
		for sum < 1000 { // for ; sum < 1000; {} 分号被自动删除
			sum += sum
		}
		fmt.Printf("sum = %v\n", sum)
	}

	//3.for 是 Go 中的 “while”
	{
		sum := 1
		for sum < 1000 {
			sum += sum
		}
		fmt.Printf("sum = %v\n", sum)
	}

	//4.无限循环
	{
		// for {

		// }
	}

	//5.if
	{
		fmt.Printf("sqrt(2) = %v, sqrt(-4) = %v\n", sqrt(2), sqrt(-4))
	}

	//6.if 的简短语句,该语句声明的变量作用域仅在 if 之内。
	{
		fmt.Printf("pow(3,2,10) = %v, pow(3,3,20) = %v\n", pow(3, 2, 10), pow(3, 3, 20))
	}

	//练习，牛顿法求平方根
	{
		for i := 1; i < 10; i++ {
			fmt.Printf("i = %v\tSqrt(i) = %v\tmath.Sqrt(i)=%v\n", i, Sqrt(float64(i)), math.Sqrt(float64(i)))
		}
	}

	//7.
	{
		fmt.Print("Go runs on ")
		switch os := runtime.GOOS; os {
		case "darwin":
			fmt.Println("OS X.")
		case "linux":
			fmt.Println("Linux")
		default:
			fmt.Printf("%s.\n", os)
		}
	}

	//8.switch 的求值顺序
	{
		fmt.Println("When's Saturday?")
		today := time.Now().Weekday()
		switch time.Saturday {
		case today + 0:
			fmt.Println("Today.")
		case today + 1:
			fmt.Println("Tomorrow.")
		case today + 2:
			fmt.Println("In two days.")
		default:
			fmt.Println("Too far away.")
		}
	}

	//9.没有条件的 switch, 同 switch true 一样。
	{
		t := time.Now()
		switch {
		case t.Hour() < 12:
			fmt.Println("Good morning!")
		case t.Hour() < 17:
			fmt.Println("Good afternoon!")
		default:
			fmt.Println("Good evening.")
		}
	}

	//fallthrough
	//搭配switch使用的关键字，默认在switch中，每个case都会有一个隐藏的break，如果想要去掉隐藏的break，可以使用fallthrough来进行取代
	{
		val := 2
		switch val {
		case 1:
			fmt.Println("val = 1")
		case 2:
			fmt.Println("val = 2")
			fallthrough //相当于nobreak
		case 3:
			fmt.Println("val = 3")
		case 4:
			fmt.Println("val = 4")
		default:
			fmt.Println("default")
		}
	}

}
