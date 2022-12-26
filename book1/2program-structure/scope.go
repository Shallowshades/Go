package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

/*
一个声明语句将程序中的实体和一个名字关联，比如一个函数或一个变量。声明语句的作用域是指源代码中可以有效使用这个名字的范围。

不要将作用域和生命周期混为一谈。
声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时的属性。
一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时的概念。

句法块是由花括弧所包含的一系列语句
句法块内部声明的名字是无法被外部块访问的。这个块决定了内部声明的名字的作用域范围。

词法块是声明在代码中并未显式地使用花括号包裹起来

？？？
对全局的源代码来说，存在一个整体的词法块，称为全局词法块；对于每个包；每个for、if和switch语句，也都有对应词法块；每个switch或select的分支也有独立的词法块；当然也包括显式书写的词法块（花括弧包含的语句）。

声明语句对应的词法域决定了作用域范围的大小
对于内置的类型、函数和常量，比如int、len和true等是在全局作用域的，因此可以在整个程序中直接使用。
任何在函数外部（也就是包级语法域）声明的名字可以在同一个包的任何源文件中访问的。
对于导入的包，例如tempconv导入的fmt包，则是对应源文件级的作用域，因此只能在当前的文件中访问导入的fmt包，当前包的其它源文件无法访问在当前源文件导入的包。
还有许多声明语句，比如tempconv.CToF函数中的变量c，则是局部作用域的，它只能在函数内部（甚至只能是局部的某些部分）访问。
控制流标号，就是break、continue或goto语句后面跟着的那种标号，则是函数级的作用域。
一个程序可能包含多个同名的声明，只要它们在不同的词法域就没有关系。
当编译器遇到一个名字引用时，它会对其定义进行查找，查找过程从最内层的词法域向全局的作用域进行。如果查找失败，则报告“未声明的名字”这样的错误。如果该名字在内部和外部的块分别声明过，则内部块的声明首先被找到。在这种情况下，内部声明屏蔽了外部同名的声明，让外部的声明的名字无法被访问。

在函数中词法域可以深度嵌套，因此内部的一个声明可能屏蔽外部的声明。

并不是所有的词法域都显式地对应到由花括弧包含的语句；还有一些隐含的规则
for语句创建了两个词法域：花括弧包含的是显式的部分，是for的循环体部分词法域，另外一个隐式的部分则是循环的初始化部分，比如用于迭代变量i的初始化。隐式的词法域部分的作用域还包含条件测试部分和循环后的迭代部分（i++），当然也包含循环体词法域。
for循环的隐式词法域包含显示词法域，不然显示词法如何使用隐式词法域中的i

if和switch语句也会在条件部分创建隐式词法域，还有它们对应的执行体词法域。

第二个if语句嵌套在第一个内部，因此第一个if语句条件初始化词法域声明的变量在第二个if中也可以访问。switch语句的每个分支也有类似的词法域规则：条件部分为一个隐式词法域，然后是每个分支的词法域。

在包级别，声明的顺序并不会影响作用域范围，因此一个先声明的可以引用它自身或者是引用后面的一个声明，这可以让我们定义一些相互嵌套或递归的类型或函数。
但是如果一个变量或常量递归引用了自身，则会产生编译错误。
*/

func f() {}

var g = "g"

func main() {

	{
		f := "f"
		fmt.Println(f) // "f"; local var f shadows package-level func f
		fmt.Println(g) // "g"; package-level var
		//fmt.Println(h) // compile error: undefined: h
	}

	/*
		下面的例子同样有三个不同的x变量，每个声明在不同的词法域，一个在函数体词法域，一个在for隐式的初始化词法域，一个在for循环体词法域；只有两个块是显式创建的
	*/
	{
		x1 := "hello"
		for _, x2 := range x1 {
			x3 := x2 + 'A' - 'a'
			fmt.Printf("%c", x3)
		}
		fmt.Println()
	}

	/*
		if和switch语句也会在条件部分创建隐式词法域，还有它们对应的执行体词法域。

		第二个if语句嵌套在第一个内部，因此第一个if语句条件初始化词法域声明的变量在第二个if中也可以访问。switch语句的每个分支也有类似的词法域规则：条件部分为一个隐式词法域，然后是每个分支的词法域。
	*/
	{
		if x := 1; x == 0 {
			fmt.Println(x)
		} else if y := 2; x == y {
			fmt.Println(x, y)
		} else {
			fmt.Println(x, y)
		}
		//fmt.Println(x, y) // compile error: x and y are not visible here
	}
	/*
		{
			if f, err := os.Open(fname); err != nil { // compile error: unused: f
			 	return err
			 }
			 f.ReadByte() // compile error: undefined f
			 f.Close()    // compile error: undefined f
		}
	*/
	{
		f, err := os.Open("1.txt")
		if err != nil {
			os.Exit(1)
		}
		io.Copy(os.Stdout, f)
		fmt.Println()
		f.Close()
	}

	/*
		可以这么写，但这不是Go语言推荐的做法
		Go语言的习惯是在if中处理错误然后直接返回，这样可以确保正常执行的语句不需要代码缩进。
	*/
	{
		if f, err := os.Open("1.txt"); err != nil {
			os.Exit(0)
		} else {
			// f and err are visible here too
			io.Copy(os.Stdout, f)
			fmt.Println()
			f.Close()
		}
	}

	{
		fmt.Println(cwd)
	}
}

var cwd string

/*
虽然cwd在外部已经声明过，但是:=语句还是将cwd和err重新声明为新的局部变量。因为内部声明的cwd将屏蔽外部的声明，因此下面的代码并不会正确更新包级声明的cwd变量。
*/
/*
func init() {
	cwd, err := os.Getwd() // compile error: unused: cwd
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
*/

/*
增加一个局部cwd的打印语句，就可能导致编译器检测失效

全局的cwd变量依然是没有被正确初始化的，而且看似正常的日志输出更是让这个BUG更加隐晦
*/
/*
func init() {
	cwd, err := os.Getwd() // wrong
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
	fmt.Println("working directory = ", cwd)
}
*/

/*
有许多方式可以避免出现类似潜在的问题。最直接的方法是通过单独声明err变量，来避免使用:=的简短声明方式：
*/
func init() {
	var err error
	cwd, err = os.Getwd()
	if err != nil {
		log.Fatalf("os.Getwd failed: %v", err)
	}
}
