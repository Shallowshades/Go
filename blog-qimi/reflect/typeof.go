package main

import (
	"fmt"
	"reflect"
)

func reflectType(x any) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %v\n", v)
}

func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 100
	reflectType(b)
}
