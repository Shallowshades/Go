package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)
	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文及其父类存活的时间超过必要的时间。
	defer cancel()

	/*
		定义了一个50毫秒之后过期的deadline，然后我们调用context.WithDeadline(context.Background(), d)得到一个上下文（ctx）和一个取消函数（cancel），然后使用一个select让主程序陷入等待：等待1秒后打印overslept退出或者等待ctx过期后退出。
		因为ctx 50毫秒后就会过期，所以ctx.Done()会先接收到context到期通知，并且会打印ctx.Err()的内容。
	*/

LOOP:
	for {
		select {
		case <-time.After(10 * time.Millisecond):
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			break LOOP
		}
	}

	fmt.Println("main over")
}
