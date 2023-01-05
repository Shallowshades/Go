package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
	通道
		保证同步的一种机制
	类型
		有缓存的、无缓存的
	无缓存
		是阻塞的，读多写少，读到的为零值
*/

var buf = make(chan int)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(256)
	fmt.Println("send: ", value)
	buf <- value
}

func main() {

	defer close(buf)
	go send()
	fmt.Println("wait...")
	value := <-buf
	fmt.Println("recv: ", value)
	fmt.Println("end...")
}
