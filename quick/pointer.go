package main

import "fmt"

func changeValue1(p int) {
	p = 10
}

func changeValue2(p *int) {
	*p = 10
}

func swap1(a, b int) {
	var temp int = a
	a = b
	b = temp
}

func swap2(pa, pb *int) {
	var temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {
	var a int = 1
	changeValue1(a)
	fmt.Println("a = ", a)

	changeValue2(&a)
	fmt.Println("a = ", a)

	a, b := 1, 2
	swap1(a, b)
	fmt.Println("a = ", a, ", b = ", b)
	swap2(&a, &b)
	fmt.Println("a = ", a, ", b = ", b)

	var p *int = &a
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)

	var pp **int = &p
	fmt.Println(&p)
	fmt.Println(pp)
	fmt.Println(**pp)
}
