package main

import "fmt"

/*
	channel多路状态监控
*/
func fib(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			y = y + x
			x = y - x
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	//child go
	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	//main go
	fib(c, quit)
}
