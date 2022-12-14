package main

import (
	"fmt"
	"time"
)

/*
信道
	信道是带有类型的管道，可以通过它用信道操作符 <- 来发送或者接收值。

	ch <- v    // 将 v 发送至信道 ch。
	v := <-ch  // 从 ch 接收值并赋予 v。
	（“箭头”就是数据流的方向。）
*/

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum //将sum送入c
}

func main() {

	/*
		和映射与切片一样，信道在使用前必须创建：

		ch := make(chan int)
		默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 Go 程可以在没有显式的锁或竞态变量的情况下进行同步。
	*/
	{
		s := []int{7, 2, 8, -9, 4, 0}
		c := make(chan int)
		go sum(s[:len(s)/2], c)
		go sum(s[len(s)/2:], c)
		x, y := <-c, <-c
		fmt.Println(x, y, x+y)
	}

	/*
		带缓冲的信道
			信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：

			ch := make(chan int, 100)
			仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
	*/
	{
		ch := make(chan int, 2)
		ch <- 1
		ch <- 2
		//ch <- 3
		//fatal error: all goroutines are asleep - deadlock!
		//当缓冲区满时，若没有读取，则陷入死锁
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		//fmt.Println(<-ch)
		//fatal error: all goroutines are asleep - deadlock!
		//当缓冲区空时，若进行读取，则陷入死锁
	}

	/*
		range 和 close
			发送者可通过 close 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配第二个参数来测试信道是否被关闭

		若没有值可以接收且信道已被关闭，那么在执行完
		v, ok := <-ch
		之后 ok 会被设置为 false。

		循环 for i := range c 会不断从信道接收值，直到它被关闭。

		*注意：* 只有发送者才能关闭信道，而接收者不能。向一个已经关闭的信道发送数据会引发程序恐慌（panic）。

		*注意：* 信道与文件不同，通常情况下无需关闭它们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。
	*/
	{
		fibonacci := func(n int, c chan int) {
			x, y := 0, 1
			for i := 0; i < n; i++ {
				c <- x
				x, y = y, x+y
			}
			close(c)
		}

		c := make(chan int, 10)
		go fibonacci(cap(c), c)
		for i := range c {
			fmt.Print(i, " ")
		}
		fmt.Println()
	}

	/*
		select 语句
			select 语句使一个 Go 程可以等待多个通信操作。

		select 会阻塞到某个分支可以继续执行为止，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。
	*/
	{
		fibonacci := func(c, quit chan int) {
			x, y := 0, 1
			for {
				select {
				case c <- x:
					x, y = y, x+y
				case <-quit:
					fmt.Println("quit")
					return
				}
			}
		}

		c := make(chan int)
		quit := make(chan int)
		go func() {
			for i := 0; i < 10; i++ {
				fmt.Print(<-c, " ")
			}
			fmt.Println()
			quit <- 0
		}()
		fibonacci(c, quit)
	}

	/*
		默认选择
			当 select 中的其它分支都没有准备好时，default 分支就会执行。

		为了在尝试发送或者接收时不发生阻塞，可使用 default 分支
	*/
	{
		tick := time.Tick(100 * time.Millisecond)
		boom := time.After(500 * time.Millisecond)
		for {
			flag := false
			select {
			case <-tick:
				fmt.Println("tick.")
			case <-boom:
				fmt.Println("BOOM!")
				flag = true
			default:
				fmt.Print("      .")
				time.Sleep(50 * time.Millisecond)
			}

			if flag == true {
				break
			}
		}
	}

}
