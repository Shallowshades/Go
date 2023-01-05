package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

var x int32 = 100

func add() {
	atomic.AddInt32(&x, 1)
}

func sub() {
	atomic.AddInt32(&x, -1)
}

func test() {
	cnt := 0
	for j := 0; j < 100; j++ {
		for i := 0; i < 10000; i++ {
			go add()
			go sub()
		}
		time.Sleep(time.Second * 1)
		fmt.Println("end x = ", atomic.LoadInt32(&x))
		if atomic.LoadInt32(&x) == 100 {
			cnt++
		} else {
			fmt.Println("error...")
		}
		atomic.StoreInt32(&x, 100)
	}
	fmt.Println("cnt = ", cnt)
}

var num int32 = 100

func cas() {
	//cas: compare and swap, old and new
	b := atomic.CompareAndSwapInt32(&num, 100, 200)
	fmt.Printf("b: %v\n", b)
	fmt.Printf("num: %v\n", num)
}

func main() {
	for i := 0; i < 5; i++ {
		go cas()
	}
	time.Sleep(time.Second)
}
