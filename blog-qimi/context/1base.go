// 初始的例子

package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func worker() {
	for {
		fmt.Println("worker")
		time.Sleep(time.Second)
	}
	//how to receive cmd to exit goroutine
	wg.Done()
}

func main() {

	wg.Add(1)
	go worker()
	//how to exit elegantly
	wg.Wait()
	fmt.Println("over")
}
