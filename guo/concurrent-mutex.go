package main

import (
	"fmt"
	"sync"
	"time"
)

var x int = 100
var lock sync.Mutex

func add() {
	lock.Lock()
	defer lock.Unlock()
	x++
	fmt.Printf("x: %v\n", x)
}

func sub() {
	lock.Lock()
	defer lock.Unlock()
	x--
	fmt.Printf("x: %v\n", x)
}

func main() {

	for i := 0; i < 10000; i++ {
		go add()
		go sub()
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("end x: %v\n", x)
}
