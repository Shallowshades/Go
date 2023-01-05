package main

import (
	"fmt"
	"os"
	"time"
)

/*
	基于select的多路复用
*/

func countdown1() {
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		<-tick
	}
}

func countdown2(abort chan struct{}) {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
}

func countdown() {
	abort := make(chan struct{})
	go countdown1()
	go countdown2(abort)

	fmt.Println("Commencing countdown. Press return to abort.")
	select {
	case <-time.After(10 * time.Second):
		launch()
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
}

func countdown3() {
	abort := make(chan struct{})
	go countdown2(abort)
	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	// Tick函数挺方便，但是只有当程序整个生命周期都需要这个时间时使用它才比较合适。
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick: //goroutine泄露
			//
		case <-abort:
			fmt.Println("Launch abort!")
			return
		}
	}
	launch()
}

func launch() {
	fmt.Println("launch...")
}

func main() {

	//channel的零值是nil
	//channel and select case
	//如果多个case同时就绪时，select会随机地选择一个执行，这样来保证每一个channel都有平等的被select的机会。
	//当buffer既不为满也不为空时，select语句的执行情况就像是抛硬币的行为一样是随机的。
	{
		ch := make(chan int, 2)
		for i := 0; i < 10; i++ {
			select {
			case x := <-ch:
				fmt.Println(x) // "0" "2" "4" "6" "8"
			case ch <- i:

			}
		}
	}

	countdown3()

	//replace Tick
	{
		ticker := time.NewTicker(1 * time.Second)
		<-ticker.C    // receive from the ticker's channel
		ticker.Stop() // cause the ticker's goroutine to terminate
	}
}
