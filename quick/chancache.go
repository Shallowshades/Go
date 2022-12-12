package main

import (
	"fmt"
	"time"
)

func main() {
	//有缓存的通道
	//空的时候阻塞，满的时候阻塞
	c := make(chan int, 3)

	fmt.Println("len(c) = ", len(c), ", cap(c) = ", cap(c))

	go func() {
		defer fmt.Println("child goroutine end...")

		for i := 0; i < 5; i++ {
			c <- i
			fmt.Println("child goroutine is running, send data = ", i,
				"len(c) = ", len(c), ", cap(c) = ", cap(c))
		}
	}()

	for i := 0; i < 5; i++ {
		num := <-c
		fmt.Println("num = ", num)
	}

	time.Sleep(1 * time.Second)

	fmt.Println("main end...")
}
