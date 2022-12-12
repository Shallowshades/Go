package main

import (
	"crypto/sha256"
	"fmt"
	"unsafe"
)

func main() {

	//数组
	{
		//数组是一个由固定长度的特定类型元素组成的序列，一个数组可以由零个或多个元素组成。
		//因为数组的长度是固定的，因此在Go语言中很少直接使用数组。
		//和数组对应的类型是Slice（切片），它是可以增长和收缩的动态序列，slice功能也更灵活，但是要理解slice工作原理的话需要先理解数组。
		//内置的len函数将返回数组中元素的个数
		{
			var a [3]int
			fmt.Println(unsafe.Sizeof(a))
			fmt.Println(a[0])        //first
			fmt.Println(a[len(a)-1]) //last

			for i, v := range a {
				fmt.Printf("%d %d\n", i, v)
			}

			for _, v := range a {
				fmt.Printf("%d\n", v)
			}

			//默认情况下，数组的每个元素都被初始化为元素类型对应的零值，对于数字类型来说就是0。
			//也可以使用数组字面值语法用一组值来初始化数组
			var q [3]int = [3]int{1, 2, 3}
			var r [3]int = [3]int{2, 3}
			fmt.Println("q = ", q)
			fmt.Println(r[2])
		}

		{
			//在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始化值的个数来计算
			q := [...]int{1, 2, 3} //同上q
			fmt.Printf("%T\n", q)
		}

		{
			//数组的长度是数组类型的一个组成部分，因此[3]int和[4]int是两种不同的数组类型。
			//数组的长度必须是常量表达式，因为数组的长度需要在编译阶段确定
			q := [3]int{1, 2, 3}
			//q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int
			fmt.Println(q)
		}

		{
			//数组、slice、map和结构体字面值的写法都很相似。
			//上面的形式是直接提供顺序初始化值序列，但是也可以指定一个索引和对应值列表的方式初始化
			type Currency int
			const (
				USD Currency = iota //美元
				EUR                 //欧元
				GBP                 //英镑
				RMB
			)
			symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
			fmt.Println(USD, symbol[USD])

			//此种方式初始化索引的顺序是无关紧要的，而且没用到的索引可以省略，和前面提到的规则一样，未指定初始值的元素将用零值初始化。
			r := [...]int{99: -1}
			fmt.Printf("r[98] = %d r[99] = %d\n", r[98], r[99])

			//如果一个数组的元素类型是可以相互比较的，那么数组类型也是可以相互比较的
			a := [2]int{1, 2}
			b := [...]int{1, 2}
			c := [2]int{1, 3}
			//d := [3]int{1, 2, 3}
			fmt.Println(a == b)
			fmt.Println(a != c)
			//fmt.Println(a > d) //mismatched types [2]int and [3]int
		}

		{
			//crypto/sha256包的Sum256函数对一个任意的字节slice类型的数据生成一个对应的消息摘要
			//消息摘要有256bit大小，因此对应[32]byte数组类型
			//如果两个消息摘要是相同的，那么可以认为两个消息本身也是相同
			//理论上有HASH码碰撞的情况，但是实际应用可以基本忽略
			//如果消息摘要不同，那么消息本身必然也是不同的
			c1 := sha256.Sum256([]byte("x"))
			c2 := sha256.Sum256([]byte("X"))
			fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
			//x 16进制 t bool T 数据类型
		}
	}
}
