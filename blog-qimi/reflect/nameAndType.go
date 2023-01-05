package main

import (
	"fmt"
	"reflect"
)

type myInt int64

func reflectType(x interface{}) {
	t := reflect.TypeOf(x)
	fmt.Println("--------")
	fmt.Printf("typeof:%v\ntype:%v\nkind:%v\n", t, t.Name(), t.Kind())
	fmt.Println("--------")
}

func main() {
	var a *float32 // 指针
	var b myInt    // 自定义类型
	var c rune     // 类型别名
	var d float32
	var e float64
	var f complex128
	reflectType(a) // type: kind:ptr
	reflectType(b) // type:myInt kind:int64
	reflectType(c) // type:int32 kind:int32
	reflectType(d)
	reflectType(e)
	reflectType(f)

	type person struct {
		name string
		age  int
	}
	type book struct{ title string }
	var p = person{
		name: "Cloud",
		age:  18,
	}
	var k = book{title: "《Golang》"}
	reflectType(p) // type:person kind:struct
	reflectType(k) // type:book kind:struct
}
