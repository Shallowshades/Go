package main

import (
	"fmt"
	"time"
)

/*
go build debug.go debug.out
GODEBUG=schedtrace=1000 ./debug.out

SCHED 						调试信息
4033ms: 					从程序启动到输出经历的时间
gomaxprocs=4 				P的数量（虚拟机分配了两个CPU核心，支持超线程）
idleprocs=4 				空闲的P的数量
threads=5 					线程数量（包括M0，GODEBUG调试的线程）
spinningthreads=0 			处于自旋状态的线程
idlethreads=3 				处于空闲状态的线程
runqueue=0 					全局G队列的数量
[0 0 0 0]					每个P的本地队列中G的数量

*/

func main() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello GMP")
	}
}
