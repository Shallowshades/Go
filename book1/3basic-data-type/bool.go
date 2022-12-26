package main

import "fmt"

/*
布尔值可以和&&（AND）和||（OR）操作符结合，并且有短路行为

&&的优先级比||高（助记：&&对应逻辑乘法，||对应逻辑加法，乘法比加法优先级要高）

布尔值并不会隐式转换为数字值0或1，反之亦然。必须使用一个显式的if语句辅助转换

! if i（int） 会报错


*/

// 需要经常做类似的转换，包装成一个函数会更方便

// btoi convert a bool to int
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// itob convert a int to bool
func itob(i int) bool { return i != 0 }

func main() {

	b := 1
	if itob(b) {
		fmt.Println("b != 0")
	} else {
		fmt.Println("b == 0")
	}
}
