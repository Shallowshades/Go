package main

import "fmt"

func main() {
	//定义一个channel
	//无缓冲必须双方等到双方同时交接，才能继续执行
	c := make(chan int) //没有缓存，顶多存（放）一个元素

	go func() {
		defer fmt.Println("goroutine ending...")

		fmt.Println("goroutine running...")

		c <- 666
	}()

	//channel 隐含同步机制， 主goroutine等待chan， 子goroutine阻塞等待
	num := <-c //此语句执行先于defer fmt.Println("goroutine ending...")

	fmt.Println("num = ", num)
	fmt.Println("main goroutine ending...")

}
