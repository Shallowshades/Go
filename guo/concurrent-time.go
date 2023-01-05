package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func(c chan int) {
		ticker := time.NewTicker(time.Millisecond * 100)
		defer close(c)
		defer ticker.Stop()
		for range ticker.C {
			select {
			case c <- 1:
			case c <- 2:
			case c <- 3:
			}
		}
	}(c)

	sum := 0
	for v := range c {
		fmt.Printf("v: %v\n", v)
		sum += v
		if sum >= 10 {
			break
		}
	}
}
