package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	dup, err := os.Open("dup.txt")
	if err != nil {
		panic(err)
	}
	counts := make(map[string]int)
	input := bufio.NewScanner(dup)
	for input.Scan() {
		counts[input.Text()]++
	}
	dup.Close()
	fmt.Println("=============")
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
