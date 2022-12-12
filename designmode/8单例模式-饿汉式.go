package main

import "fmt"

/*
	饿汉式
	线程安全，但是无论用不用到该实例都会创建一个
*/

type singleton struct{}

var instance *singleton = new(singleton)

func (s *singleton) SomeThing() {
	fmt.Println("a certain function...")
}

func GetInstance() *singleton {
	return instance
}

func main() {
	s := GetInstance()
	s.SomeThing()

	s1 := GetInstance()
	s1.SomeThing()

	if s == s1 {
		fmt.Println("s == s1")
	} else {
		fmt.Println("s != s1")
	}
}
