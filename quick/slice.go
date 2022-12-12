package main

import "fmt"

//值拷贝传递
func printArray1(myArray [4]int) {
	for index, value := range myArray {
		fmt.Println("index = ", index, ", value = ", value)
	}
}

//引用传递
func printArray2(myArray []int) {
	for _, value := range myArray {
		fmt.Println("value = ", value)
	}
	myArray[0] = 100
}

func main() {
	//固定长度的数组
	var myArray1 [10]int
	myArray2 := [10]int{1, 2, 3, 4, 5, 6, 7, 8}
	myArray3 := [4]int{23, 34, 56, 78}
	for i := 0; i < len(myArray1); i++ {
		fmt.Print(myArray1[i], " ")
	}
	fmt.Println()

	//range
	for index, value := range myArray2 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	//查看数组的数据类型
	fmt.Printf("myArray1 types = %T\n", myArray1)
	fmt.Printf("myArray2 types = %T\n", myArray2)
	fmt.Printf("myArray3 types = %T\n", myArray3)

	printArray1(myArray3)
	//printArray1(myArray1) error

	//动态数组，切片slice
	//动态数组本身就是一个指针
	myArray4 := []int{1, 2, 3, 4}
	fmt.Printf("myArray4 types is %T\n", myArray4)
	printArray2(myArray4)
	fmt.Println("--------")
	for _, value := range myArray4 {
		fmt.Println("value = ", value)
	}

	//slice声明方式
	//1.声明并初始化
	slice1 := []int{1, 2, 3}
	fmt.Printf("len = %d, slice = %v\n", len(slice1), slice1)

	//2.声明，但并未分配空间
	var slice2 []int
	//slice2[0] = 100 //error
	slice2 = make([]int, 4) //make分配空间
	fmt.Printf("len = %d, slice = %v\n", len(slice2), slice2)

	//3.声明，并同时分配空间，默认值为0
	var slice3 []int = make([]int, 5)
	fmt.Printf("len = %d, slice = %v\n", len(slice3), slice3)

	//4.声明，分配空间，并推导slice为切片
	slice4 := make([]int, 6)
	fmt.Printf("len = %d, slice = %v\n", len(slice4), slice4)

	if slice4 == nil {
		fmt.Println("slice4是一个空切片")
	} else {
		fmt.Println("slice4是有空间的")
	}

	//追加
	//扩容， double
	numbers := make([]int, 3, 4)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 1)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 2)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers)

	//截取
	s1 := []int{1, 2, 3}
	s2 := s1[0:2] //左闭右开
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v\n", s2)
	s1[0] = 100
	fmt.Printf("s1 = %v\n", s1)
	fmt.Printf("s2 = %v\n", s2)

	//拷贝
	//copy可以将底层数组的slice一起进行copy
	s3 := make([]int, 3)
	copy(s3, s1) //将s1拷贝的s3中
	fmt.Printf("s3 = %v\n", s3)
}
