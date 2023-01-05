package main

import (
	"fmt"
	"time"
)

var chanInt = make(chan int, 0)
var chanStr = make(chan string)

func main() {

	go func() {
		defer close(chanInt)
		defer close(chanStr)
		chanInt <- 100
		chanStr <- "hello"
	}()

	for {
		select {
		case r := <-chanInt: //如果不关闭会读默认值
			fmt.Printf("chanInt: %v\n", r)
		case r := <-chanStr:
			fmt.Printf("chanStr: %v\n", r)
		default:
			fmt.Println("default")
		}
		time.Sleep(1 * time.Second)
	}
}
