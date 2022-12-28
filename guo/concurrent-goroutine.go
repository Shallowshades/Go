package main

import (
	"fmt"
	"time"
)

func show(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg: %v\n", msg)
		time.Sleep(time.Millisecond * 100)
	}
}

func main() {

	//main函数中的协程，如果main结束了，协程也会结束，mian函数结束意味着程序结束了吧
	//非main函数里的协程，函数结束了，只要main没结束，协程就会执行。
	go show("let's go")
	go show("go go go")
	time.Sleep(time.Second)
}
