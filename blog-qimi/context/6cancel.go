package main

import (
	"context"
	"fmt"
)

/*
	gen函数在单独的goroutine中生成整数并将它们发送到返回的通道。
	gen的调用者在使用生成的整数之后需要取消上下文，以免gen启动的内部goroutine发生泄漏。
*/

func gen(ctx context.Context) <-chan int {
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // return结束该goroutine，防止泄露
			case dst <- n:
				n++
			}
		}
	}()
	fmt.Println("gen over")
	return dst
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	fmt.Println("main start...")
	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Println("main over...")
}
