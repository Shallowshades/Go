package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	runtime包
		协程管理相关的api
*/

func show1(msg string) {
	for i := 0; i < 2; i++ {
		fmt.Println(msg)
	}
}

func show2(msg string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("i: %v\n", i)
		if i > 5 {
			//exit this routine
			runtime.Goexit()
		}
	}
}

func a() {
	for i := 0; i < 10; i++ {
		fmt.Printf("ai: %v\n", i)
		time.Sleep(time.Microsecond * 10)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Printf("bi: %v\n", i)
		time.Sleep(time.Microsecond * 10)
	}
}

func main() {
	//start a goroutine
	go show1("hello")

	//main routine
	for i := 0; i < 2; i++ {
		//让出cpu时间片，重新等待安排-
		runtime.Gosched()
		fmt.Println(runtime.Version())
	}
	//fmt.Println("main routine end...")

	go show2("kk")
	time.Sleep(time.Second)

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
	runtime.GOMAXPROCS(3)
	go a()
	go b()
	time.Sleep(time.Second)
}
