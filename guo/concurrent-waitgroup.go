package main

import (
	"fmt"
	"sync"
)

/*
	sync.WaitGroup
		实现同步
*/

var wg sync.WaitGroup

func show(num int) {
	defer wg.Done()
	fmt.Printf("num: %v\n", num)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go show(i)
	}

	wg.Wait()

	//main routine
	fmt.Println("end...")
}
