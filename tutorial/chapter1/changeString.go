package main

import "fmt"

/*
要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。
无论哪种转换，都会重新分配内存，并复制字节数组。
*/
func changeString() {
	s1 := "big"
	byteS1 := []byte(s1)
	byteS1[0] = 'B'
	fmt.Println(string(byteS1))

	s2 := "兔子"
	runeS2 := []rune(s2)
	runeS2[0] = '狗'
	fmt.Println(string(runeS2))
}

/*
统计一个字符串中汉字的字数
"hello沙河小王子"

unicode编码中汉字大于255
*/

func countChinese(str string) int {
	var n int
	s := []rune(str)
	for _, v := range s {
		if v >= 256 {
			n++
		}
	}
	return n
}

/*
有一堆数字，如果除了一个数字以外，其他数字都出现了两次，那么如何找到出现一次的数字？
*/
func xorNumber(arr []int) int {
	num := 0
	for _, v := range arr {
		num ^= v
	}
	return num
}

func main() {

	changeString()
	fmt.Println(countChinese("hello沙河小王子"))
	fmt.Println(xorNumber([]int{23, 465, 34, 56, 23, 465, 34}))
}
