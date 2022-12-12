package main

import (
	"fmt"
	"runtime"
	"time"
)

/*
	co-routine 协程

golang对协程的处理
co-routine -> goroutine
大小几KB

	GMP
G		gotoutine协程
P		processor处理器
M		thread线程

	调度器设计策略
复用线程 ：      Handoff
利用并行 ：
抢占	： 最多 10ms
全局G队列 ： 优先偷取全局的G队列中的goroutine，其次偷取其他的本地队列中的goroutine
*/

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %v\n", i)
		time.Sleep(1 * time.Second)
	}
}

func test1() {
	//创建一个go程去执行newTask()流程
	go newTask()
}

func main() {

	// i := 0
	// for {
	// 	i++
	// 	fmt.Printf("main goroutine : i = %d\n", i)
	// 	time.Sleep(1 * time.Second)
	// }

	//fmt.Println("main goroutine exit...") //若主go程结束，子go程直接结束

	//无参的
	//创建一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit() //退出goroutine
			fmt.Println("B")
		}()

		fmt.Println("A")
	}()

	//有参的
	//无法拿到返回值，需要特殊机制
	go func(a, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	for {
		time.Sleep(1 * time.Second)
	}
}
