package main

import (
	"fmt"
	"strings"
)

func main() {

	/*
		数组
			类型 [n]T 表示拥有 n 个 T 类型的值的数组。
		表达式
			var 变量名 [长度]类型
		数组的长度是其类型的一部分，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式来使用数组。
	*/
	{
		var a [2]string
		a[0] = "Hello"
		a[1] = "World"
		fmt.Println(a[0], a[1])
		fmt.Println(a)

		primes := [6]int{2, 3, 5, 7, 11, 13}
		fmt.Println(primes)
	}

	/*
		切片
			每个数组的大小都是固定的。而切片则为数组元素提供动态大小的、灵活的视角。在实践中，切片比数组更常用。

		类型 []T
			表示一个元素类型为 T 的切片。

		切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔：
			a[low : high] 左闭右开
	*/
	{
		primes := [6]int{2, 3, 5, 7, 11, 13}
		s := primes[1:4]
		fmt.Println(s)
	}

	/*
		切片就像数组的引用
			切片并不存储任何数据，它只是描述了底层数组中的一段。
		更改切片的元素会修改其底层数组中对应的元素。
		与它共享底层数组的切片都会观测到这些修改
	*/
	{
		arr := []int{1, 2, 3, 4}
		fmt.Println("arr[1:1] = ", arr[1:1])

		s1 := make([]int, 3) //[0 0 0]
		s2 := s1             //将s1直接赋值给s2，s1和s2共用一个底层数组
		s2[0] = 100
		fmt.Println(s1) //[100 0 0]
		fmt.Println(s2) //[100 0 0]
	}

	/*
		切片文法
			切片文法类似于没有长度的数组文法。

		这是一个数组文法：
			[3]bool{true, true, false}
		下面这样则会创建一个和上面相同的数组，然后构建一个引用了它的切片：
			[]bool{true, true, false}
	*/
	{
		q := []int{2, 3, 5, 7, 11, 13}
		fmt.Println(q)

		r := []bool{true, false, true, true, false, true}
		fmt.Println(r)

		s := []struct {
			i int
			b bool
		}{
			{2, true},
			{3, false},
			{5, true},
			{7, true},
			{11, false},
			{13, true},
		}
		fmt.Println(s)
	}

	/*
		切片的默认行为
			在进行切片时，可以利用它的默认行为来忽略上下界。
		切片下界的默认值为 0，上界则是该切片的长度。

		对于数组
		var a [10]int
		来说，以下切片是等价的：
		a[0:10]
		a[:10]
		a[0:]
		a[:]
	*/
	{
		s := []int{2, 3, 5, 7, 11, 13}

		s = s[1:4]
		fmt.Println(s)

		s = s[:2]
		fmt.Println(s)

		s = s[1:]
		fmt.Println(s)
	}

	/*
		nil 切片
			切片的零值是 nil。
		nil 切片的长度和容量为 0 且没有底层数组。
	*/
	{
		var s []int
		fmt.Println(s, len(s), cap(s))
		if s == nil {
			fmt.Println("nil!")
		}
	}

	/*
		用 make 创建切片
			切片可以用内建函数 make 来创建，这也是创建动态数组的方式。

		make 函数会分配一个元素为零值的数组并返回一个引用了它的切片
		a := make([]int, 5)  // len(a)=5

		要指定它的容量，需向 make 传入第三个参数：
		b := make([]int, 0, 5) // len(b)=0, cap(b)=5
	*/
	{
		a := make([]int, 5)
		PrintSlice("a", a)

		b := make([]int, 0, 5)
		PrintSlice("b", b)

		c := b[:2]
		PrintSlice("c", c)

		d := c[2:5]
		PrintSlice("d", d)

		/*
			a len=5 cap=5 [0 0 0 0 0]
			b len=0 cap=5 []
			c len=2 cap=5 [0 0]
			d len=3 cap=3 [0 0 0]
		*/
	}

	/*
		切片的切片
			切片可包含任何类型，甚至包括其它的切片
	*/
	{
		// 创建一个井字板（经典游戏）
		board := [][]string{
			{"·", "·", "·"},
			{"·", "·", "·"},
			{"·", "·", "·"},
		}

		// 两个玩家轮流打上 X 和 O
		board[0][0] = "X"
		board[2][2] = "O"
		board[1][2] = "X"
		board[1][0] = "O"
		board[0][2] = "X"

		for i := 0; i < len(board); i++ {
			fmt.Printf("%s\n", strings.Join(board[i], "  "))
		}
	}

	/*
		向切片追加元素
		func append(s []T, vs ...T) []T

		当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。
	*/
	{
		var s []int
		printSlice(s)

		s = append(s, 0)
		printSlice(s)

		s = append(s, 1)
		printSlice(s)

		s = append(s, 2, 3, 4)
		printSlice(s)
	}

	/*
		Range
			for 循环的 range 形式可遍历切片或映射。

		当使用 for 循环遍历切片时，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份'副本'。

		可以将下标或值赋予 _ 来忽略它
	*/
	{
		var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
		for i, v := range pow {
			fmt.Printf("2**%d = %d\n", i, v)
		}

		for _, v := range pow {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	/*
		slice append添加
	*/
	{
		s1 := []int32{0, 1, 2, 3, 4, 5}
		s2 := s1[:2]
		s3 := s1[3:]
		fmt.Println("s1 = ", s1)
		fmt.Println("s2 = ", s2)
		fmt.Println("s3 = ", s3)
		fmt.Printf("&s1 = %p\n", s1)
		fmt.Printf("&s2 = %p\n", s2)
		fmt.Printf("&s3 = %p\n", s3)

		//fmt.Println(s2[2]) 越界
		s2 = append(s2, 33) //修改底层数组，并不copy新的数组
		fmt.Println(s1[2])
		fmt.Println(s2[2])

		s3 = append(s3, 66) //copy新的数组
		fmt.Println(s1)
		fmt.Println(s3)
		fmt.Printf("&s1 = %p\n", s1)
		fmt.Printf("&s2 = %p\n", s2)
		fmt.Printf("&s3 = %p\n", s3)
		//fmt.Println(s1[6]) 越界
		fmt.Println(s3[3])
	}

	/*
		slice扩容原则

		step1:
			预估扩容后的容量
			OldCap 旧容量
			NewCap 新容量
			Cap    预估容量
			ints := []int{1,2}			OldCap = 2
			ints = append(ints,3,4,5)	至少Cap = 5?
			2*2 < 5 -> NewCap = 5
		规则：
			1.OldCap * 2 < Cap -> NewCap = Cap
			2.
				1> OldLen < 1024 -> NewCap = OldCap * 2
				2> OldLen >= 1024 -> NewCap = OldCap *1.25
		step2:
			newCap个元素需要多大内存
			预估容量 * 元素大小

			每个语言有自己的内存管理模块
			它会提前向os申请一批内存并管理起来
			匹配到足够大且正好合适的内存规格

			go中string是一种数据类型，占16字节
			前8字节是一个指针，指向字符串的地址
			后8字节是一个整数，标识字符串的长度
	*/
	{
		var s1 []int
		s1 = make([]int, 1000, 1000)
		fmt.Println(len(s1), cap(s1))
		s1 = append(s1, 10)
		fmt.Println(len(s1), cap(s1))
	}
	{
		var s1 []int
		s1 = make([]int, 1023, 1023)
		fmt.Println(len(s1), cap(s1))
		s1 = append(s1, 10)
		fmt.Println(len(s1), cap(s1))
	}
	{
		var s1 []int
		s1 = make([]int, 1024, 1024)
		fmt.Println(len(s1), cap(s1))
		s1 = append(s1, 10)
		fmt.Println(len(s1), cap(s1))
	}
	{
		var s1 []int
		s1 = make([]int, 1025, 1025)
		fmt.Println(len(s1), cap(s1))
		s1 = append(s1, 10)
		fmt.Println(len(s1), cap(s1))
	}
	{

	}
}

/*
练习：

	实现 Pic。它应当返回一个长度为 dy 的切片，其中每个元素是一个长度为 dx，元素类型为 uint8 的切片。当你运行此程序时，它会将每个整数解释为灰度值（好吧，其实是蓝度值）并显示它所对应的图像。
*/
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dx)
	for i := 0; i < dx; i++ {
		s[i] = make([]uint8, dy)
		for j := 0; j < dy; j++ {
			s[i][j] = uint8((i + j) / 2)
			//s[i][j] = uint8(i * j)
			//s[i][j] = uint8(i ^ j)
			//s[i][j] = uint8(i * int(math.Log2(float64(j))))
			//s[i][j] = uint8(i % (j + 1))
		}
	}
	return s
}

func PrintSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
