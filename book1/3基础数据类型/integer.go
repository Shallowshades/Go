package main

import "fmt"

/*
Unicode字符rune类型是和int32等价的类型，通常用于表示一个Unicode码点。这两个名称可以互换使用。
同样byte也是uint8类型的等价类型，byte类型一般用于强调数值是一个原始的数据而不是一个小的整数。

无符号的整数类型uintptr，没有指定具体的bit大小但是足以容纳指针。uintptr类型只有在底层编程时才需要，特别是Go语言和C语言函数库或操作系统接口相交互的地方

Go语言中关于算术运算、逻辑运算和比较运算的二元运算符，它们按照优先级递减的顺序排列
	*      /      %      <<       >>     &       &^
	+      -      |      ^
	==     !=     <      <=       >      >=
	&&
	||
二元运算符有五种优先级。
在同一个优先级，使用左优先结合规则，但是使用括号可以明确优先顺序，使用括号也可以用于提升优先级
例如mask & (1 << 28)

算术运算符+、-、*和/可以适用于整数、浮点数和复数，但是取模运算符%仅用于整数间的运算。

对于不同编程语言，%取模运算的行为可能并不相同。
在Go语言中，%取模运算符的符号和被取模数的符号总是一致的，因此-5%3和-5%-3结果都是-2。

一个算术运算的结果，不管是有符号或者是无符号的，如果需要更多的bit位才能正确表示的话，就说明计算结果是溢出了。"超出的高位的bit位部分将被丢弃"	。如果原始的数值是有符号类型，而且最左边的bit位是1的话，那么最终结果可能是负的

对于整数，+x是0+x的简写，-x则是0-x的简写；对于浮点数和复数，+x就是x，-x则是x 的负数。

&      位运算 AND
|      位运算 OR
^      位运算 XOR
&^     位清空（AND NOT）
<<     左移
>>     右移

使用了Printf函数的%b参数打印二进制格式的数字；其中%08b中08表示打印至少8个字符宽度，不足的前缀部分用0填充。

在x<<n和x>>n移位运算中，
决定了移位操作的bit数部分必须是无符号数；
被操作的x可以是有符号数或无符号数。算术上，一个x<<n左移运算等价于乘以2^n，一个x>>n右移运算等价于除以2^n。

对于每种类型T，如果转换允许的话，类型转换操作T(x)将x转换为T类型。许多整数之间的相互转换并不会改变数值；它们只是告诉编译器如何解释这个值。但是对于将一个大尺寸的整数类型转为一个小尺寸的整数类型，或者是将一个浮点数转为整数，可能会改变数值或丢失精度

浮点数到整数的转换将丢失任何小数部分，然后向数轴零方向截断。
应该避免对可能会超出目标类型表示范围的数值做类型转换，因为截断的行为可能依赖于具体的实现

任何大小的整数字面值都可以用以0开始的八进制格式书写，例如0666；
或用以0x或0X开头的十六进制格式书写，例如0xdeadbeef。
十六进制数字可以用大写或小写字母。
如今八进制数据通常用于POSIX操作系统上的文件访问权限标志，十六进制数字则更强调数字值的bit位模式

fmt的两个使用技巧。
通常Printf格式化字符串包含多个%参数时将会包含对应相同数量的额外操作数，但是%之后的[1]副词告诉Printf函数再次使用第一个操作数。第二，%后的#副词告诉Printf在用%o、%x或%X输出时生成0、0x或0X前缀。

字符面值通过一对单引号直接包含对应字符。最简单的例子是ASCII中类似'a'写法的字符面值，但是我们也可以通过转义的数值来表示任意的Unicode码点对应的字符

字符使用%c参数打印，或者是用%q参数打印带单引号的字符
*/

func main() {

	{
		var x uint = 1 //不同的编译器即使在相同的硬件平台上可能产生不同的大小 32 or 64
		fmt.Println(x << 63)
	}

	{
		var u uint8 = 255
		fmt.Println(u, u+1, u*u) // "255 0 1"

		var i int8 = 127
		fmt.Println(i, i+1, i*i) // "127 -128 1"

	}

	// &^     位清空（AND NOT）
	// 清空x中bit位为1且y中bit位为1的bit位
	{
		var x, y uint8 = 255, 67
		fmt.Printf("x = %08b\n", x)
		fmt.Printf("y = %08b\n", y)
		fmt.Println("&^")
		fmt.Printf("z = %08b\n", x&^y)
	}

	{
		var x uint8 = 1<<1 | 1<<5
		var y uint8 = 1<<1 | 1<<2

		fmt.Printf("%08b\n", x) // "00100010", the set {1, 5}
		fmt.Printf("%08b\n", y) // "00000110", the set {1, 2}

		fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
		fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
		fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
		fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}

		for i := uint(0); i < 8; i++ {
			if x&(1<<i) != 0 { // membership test
				fmt.Println(i) // "1", "5"
			}
		}

		fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
		fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	}

	{
		medals := []string{"gold", "silver", "bronze"}
		//如果len函数返回一个无符号数，那么i也将是无符号的uint类型，然后条件i >= 0则永远为真
		for i := len(medals) - 1; i >= 0; i-- {
			fmt.Println(medals[i]) // "bronze", "silver", "gold"
		}
	}

	{
		var apples int32 = 1
		var oranges int16 = 2
		var compote int = int(apples) + int(oranges)
		fmt.Println(compote)

		f := 3.141 // a float64
		i := int(f)
		fmt.Println(f, i) // "3.141 3"
		f = 1.99
		fmt.Println(int(f)) // "1"
	}

	{
		o := 0666
		fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
		x := int64(0xdeadbeef)
		fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
		// Output:
		// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	}

	{
		ascii := 'a'
		unicode := '国'
		newline := '\n'
		fmt.Printf("%d %[1]c %[1]q\n", ascii)   // "97 a 'a'"
		fmt.Printf("%d %[1]c %[1]q\n", unicode) // "22269 国 '国'"
		fmt.Printf("%d %[1]q\n", newline)       // "10 '\n'"
	}
}
