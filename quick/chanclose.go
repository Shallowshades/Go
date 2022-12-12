package main

import "fmt"

func main() {

	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
			//close(c) //panic: send on closed channel
		}

		//关闭channel
		close(c)
	}()

	/*
		for {
			//ok为true表示channel未关闭，false表示channel已关闭
			//死循环读取，若不关闭channel，则会deadlock
			//即使channel已经关闭，也可以继续从channel中接受数据（相当于关闭的是发送端）
			//对于nil channel(没有make)，无论收发都会阻塞

			if data, ok := <-c; ok {
				fmt.Println("data = ", data)
			} else {
				break
			}
		}
	*/

	//可以使用range来迭代不断操作channel
	for data := range c {
		fmt.Println(data)
	}

	//select, 同一流程下监视多个channel的状态

	fmt.Println("Main Finished...")
}
