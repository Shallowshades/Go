package main

import "fmt"

//
func sum() func(int) int {
	x := 0
	return func(y int) int {
		x += y
		return x
	}
}

func main() {

	f := sum()
	fmt.Println(f(10))
}
