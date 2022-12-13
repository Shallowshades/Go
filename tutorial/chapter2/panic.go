package main

import "fmt"

func fa() {
	fmt.Println("func A")
}

func fb() {
	panic("panic in B")
}

func fc() {
	fmt.Println("func C")
}

func main() {
	fa()
	fb()
	fc()
}
