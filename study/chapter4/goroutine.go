package main

import (
	"fmt"
	"time"
)

/*
Go 程
	Go 程（goroutine）是由 Go 运行时管理的轻量级线程。
	
	Go 程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。
	sync 包提供了这种能力，不过在 Go 中并不经常用到，因为还有其它的办法
*/

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	go say("World")
	say("Hello")
}
