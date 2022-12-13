package main

import (
	"fmt"
	"unicode/utf8"
)

// slice类似下面结构体的聚合类型
type IntSlice struct {
	ptr      *int
	len, cap int
}

func main() {
	{
		//一个slice由三个部分构成：指针、长度和容量
		//slice的底层确实引用一个数组对象
		//指针指向第一个slice元素对应的底层数组元素的地址
		//要注意的是slice的第一个元素并不一定就是数组的第一个元素
		//长度对应slice中元素的数目；长度不能超过容量
		//容量一般是从slice的开始位置到底层数据的结尾位置
		//内置的len和cap函数分别返回slice的长度和容量
		months := [...]string{
			1:  "January",
			2:  "February",
			3:  "March",
			4:  "April",
			5:  "May",
			6:  "June",
			7:  "July",
			8:  "August",
			9:  "September",
			10: "October",
			11: "November",
			12: "December",
		}

		//多个slice之间可以共享底层的数据，并且引用的数组部分区间可能重叠
		Q1 := months[1:4]
		Q2 := months[4:7]
		summer := months[6:9]
		fmt.Printf("%T\t%+[1]v\n", months)
		fmt.Printf("%T\t%+[1]v\n", Q1)
		fmt.Printf("%T\t%+[1]v\n", Q2)
		fmt.Printf("%T\t%+[1]v\n", summer)

		//切片操作超出len(s)则是意味着扩展了slice,不超过cap(s)
		//超出cap(s)的上限将导致一个panic异常
		endlessSummer := summer[:5]
		fmt.Println(endlessSummer)
		//fmt.Println(summer[:20]) //panic: out of range

	}

	{
		//slice值包含指向第一个slice元素的指针，因此向函数传递slice将允许在函数内部修改底层数组的元素。
		//换句话说，复制一个slice只是对底层的数组创建了一个新的slice别名
		s := []int{0, 1, 2, 3, 4, 5}
		fmt.Println(s)
		reverse(s)
		fmt.Println(s)

		reverse(s)
		//单独反转部分元素
		reverse(s[:2])
		reverse(s[2:])
		reverse(s)
		fmt.Println(s)

		//混合初始化 初始化序列 + 索引对
		s1 := []int{1, 2, 5: 5}
		fmt.Println(s1)
	}

	{
		//和数组不同的是，slice之间不能比较
		//bytes.Equal函数来判断两个字节型slice是否相等（[]byte）
		//对于其他类型的slice，必须展开每个元素进行比较
		s1 := []string{"Cloud", "Alice", "Tifa", "Barret"}
		s2 := s1[:3]
		s3 := s1[2:]
		//fmt.Println(s2 == s3) //slice can only be compared to nil
		fmt.Println(equal(s2, s3))

		//slice唯一合法的比较操作是和nil比较
		if s1 == nil {
			fmt.Println("s1 == nil")
		} else {
			fmt.Println("s1 != nil")
		}

		//一个零值的slice等于nil
		//一个nil值的slice并没有底层数组。
		//一个nil值的slice的长度和容量都是0
		//但是也有非nil值的slice的长度和容量也是0的
		s4 := []int{}
		s5 := make([]int, 3)[3:]
		fmt.Printf("%T\tlen(s4)=%d\tcap(s4)=%d\n", s4, len(s4), cap(s4))
		fmt.Printf("%T\tlen(s5)=%d\tcap(s5)=%d\n", s5, len(s5), cap(s5))
		if s4 == nil {
			fmt.Println("s4 == nil")
		} else {
			fmt.Println("s4 != nil")
		}
		if s5 == nil {
			fmt.Println("s5 == nil")
		} else {
			fmt.Println("s5 != nil")
		}

		var s6 []int
		if s6 == nil {
			fmt.Println("s6 == nil")
		} else {
			fmt.Println("s6 != nil")
		}

		//测试一个slice是否是空的,使用len(s)==0来判断,而不应该用s==nil来判断
		//除了和nil相等比较外，一个nil值的slice的行为和其它任意0长度的slice一样
		var s7 []int    // len(s) == 0, s == nil
		s7 = nil        // len(s) == 0, s == nil
		s7 = []int(nil) // len(s) == 0, s == nil
		s7 = []int{}    // len(s) == 0, s != nil
		fmt.Printf("%T\tlen(s7)=%d\tcap(s7)=%d\n", s7, len(s7), cap(s7))
	}

	{
		//内置的make函数创建一个指定元素类型、长度和容量的slic
		//容量部分可以省略，在这种情况下，容量将等于长度
		len, cap := 3, 5
		s1 := make([]int, len)
		s2 := make([]int, len, cap) // same as make([]T, cap)[:len]
		fmt.Println(s1, s2)

		//在底层，make创建了一个匿名的数组变量，然后返回一个slice；
		//只有通过返回的slice才能引用底层匿名的数组变量。
		//在第一种语句中，slice是整个数组的view。
		//在第二个语句中，slice只引用了底层数组的前len个元素，但是容量将包含整个的数组。额外的元素是留给未来的增长用的。
	}

	{
		//append
		var runes []rune
		for _, r := range "Hello, 世界" {
			runes = append(runes, r)
		}
		fmt.Printf("%q\n", runes)
	}

	{
		var x, y []int
		for i := 0; i < 10; i++ {
			y = appendInt(x, i)
			fmt.Printf("%d cap=%d \t%v\n", i, cap(y), y)
			x = y
		}
	}

	{
		s := []int{0, 1}
		fmt.Printf("s=%v\tlen=%d\tcap=%d\n", s, len(s), cap(s))
		s = append(s, 2, 3, 4)
		fmt.Printf("s=%v\tlen=%d\tcap=%d\n", s, len(s), cap(s))
	}

	{
		//1,2,4,..,256,512,848,1280
		// s := []int{0, 1}
		// for i := 2; i < 1000; i++ {
		// 	s = append(s, i)
		// 	fmt.Printf("i = %d\tlen=%d\tcap=%d\n", i, len(s), cap(s))
		// }
	}

	{
		var x []int
		x = append(x, 1)
		x = append(x, 2, 3)
		x = append(x, 4, 5, 6)
		x = append(x, x...)
		fmt.Println(x)
	}

	{
		var x []int
		x = appendInts(x, 1)
		x = appendInts(x, 2, 3)
		x = appendInts(x, 4, 5, 6)
		x = appendInts(x, x...)
		fmt.Println(x)
	}

	{
		//输入的slice和输出的slice共享一个底层数组
		data := []string{"one", "", "three"}
		fmt.Printf("%q\n", nonempty(data))
		fmt.Printf("%q\n", data)
	}

	{
		//移动
		s1 := []int{5, 6, 7, 8, 9}
		fmt.Println(remove(s1, 2))
		//填充
		s2 := []int{5, 6, 7, 8, 9}
		fmt.Println(remove2(s2, 2))
	}

	{
		//test stack
		st := newStack()
		for i := 0; i < 5; i++ {
			st.push(i)
		}

		for !st.empty() {
			fmt.Println(*st)
			fmt.Println(st.peak())
			fmt.Println(st.pop())
		}
	}

	{
		//test 4.3
		a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println(a)
		reverse2(&a)
		fmt.Println(a)
		reverse(a[:])
		//test4.4
		fmt.Println(rotate(a[:], 4))
		//test4.5
		s := []string{"Cloud", "Alice", "Alice", "Tifa", "Tifa", "Barret", "Alice", "Zack", "Zack"}
		fmt.Println(s)
		fmt.Println(nonrepeat(s))
		//test4.6
		b := []byte("abc \t\r\ndef \f g h 你好，    世    界ijk  l")
		fmt.Println(string(b))
		fmt.Println(string(nonrepeatspace(b)))
		//test4.7
		b = []byte("你好a, Hello, 世界")
		fmt.Println(string(b))
		fmt.Println(string(reverse3(b)))
	}

}

// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 练习 4.3： 重写reverse函数，使用数组指针代替slice。
func reverse2(s *[10]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// 练习 4.4： 编写一个rotate函数，通过一次循环完成旋转。
func rotate(s []int, n int) []int {
	//将头添加到尾部，并舍去头
	for i := 0; i < n; i++ {
		s = append(s, s[0])[1:]
	}
	return s
}

// 练习 4.7： 修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存？
// 有些rune占好几个byte 暂时想不通
// 此版本新分配内存
func reverse3(b []byte) []byte {
	c := []byte{}
	for i, n := len(b), 1; i > 0; i -= n {
		_, n = utf8.DecodeLastRune(b[:i])
		for j := i - n; j < i; j++ {
			c = append(c, b[j])
		}
	}
	return c
}

// equal compares whether two slices are equal
func equal(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

// append int version
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

// append ints... version
func appendInts(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen > cap(x) {
		zcap := zlen
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	copy(z[len(x):], y)
	return z
}

// nonempty函数将在原有slice内存空间之上返回不包含空字符串的列表
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// nonempty2用append实现，也是共用底层数组
func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

// 练习 4.5： 写一个函数在原地完成消除[]string中相邻重复的字符串的操作。
func nonrepeat(strings []string) []string {
	out := strings[:1]
	for _, c := range strings {
		if c != out[len(out)-1] {
			out = append(out, c)
		}
	}
	return out
}

// 删除slice中间的某个元素并保存原有的元素顺序,用copy移动后面的数
func remove(s []int, idx int) []int {
	copy(s[idx:], s[idx+1:])
	return s[:len(s)-1]
}

// 删除元素后不用保持原来顺序, 直接用最后一个数填充删除的位置
func remove2(s []int, idx int) []int {
	s[idx] = s[len(s)-1]
	return s[:len(s)-1]
}

// byte is a space?
func IsSpace(b byte) bool {
	switch b {
	//水平制表符(\t)、垂直制表符(\v)、换页符(\f)、回车(\r)、1000 0101、1010 0000
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}

// 练习 4.6： 编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考unicode.IsSpace）替换成一个空格返回
func nonrepeatspace(bs []byte) []byte {
	out := bs[:0]
	flag := false
	for _, b := range bs {
		if !IsSpace(b) {
			out = append(out, b)
			flag = false
		} else if !flag {
			flag = true
			out = append(out, b)
		}
	}
	return out
}

// 模拟栈
type stack struct {
	s []int
}

func newStack() *stack {
	return &stack{}
}

func (st *stack) empty() bool {
	return len(st.s) == 0
}

func (st *stack) peak() int {
	return st.s[len(st.s)-1]
}

func (st *stack) push(x int) {
	st.s = append(st.s, x)
}

func (st *stack) pop() int {
	defer func() {
		st.s = st.s[:len(st.s)-1]
	}()
	return st.peak()
}

func (st *stack) Write() string {
	return fmt.Sprintf("%v", st.s)
}
