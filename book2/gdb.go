package main

import (
	"fmt"
	"os"
)

func Println(a ...interface{}) (n int, err error) {
	println("draven")
	return fmt.Fprintln(os.Stdout, a...)
}

func main() {

	Println("Hello, world")
}
