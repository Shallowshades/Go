package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker(ctx context.Context) {
LOOP:
	for {
		fmt.Println("db connect...")
		time.Sleep(time.Millisecond * 10) //假设正常连接数据库耗时10ms
		select {
		case <-ctx.Done(): //50ms后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
	fmt.Println("worker over")
}

func main() {
	//设置一个50ms的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second)
	fmt.Println("here")
	cancel() //通知子goroutine结束
	wg.Wait()
	fmt.Println("over")
}
