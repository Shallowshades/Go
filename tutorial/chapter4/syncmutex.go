package main

import (
	"fmt"
	"sync"
	"time"
)

/*
sync.Mutex
	信道非常适合在各个 Go 程间进行通信。

	sync.Mutex也可保证每次只有一个 Go 程能够访问一个共享的变量，从而避免冲突

	互斥（mutual*exclusion），通常使用 *互斥锁（Mutex）* 这一数据结构来提供这种机制。

	Go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法：
		Lock
		Unlock
	可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行。

	也可以用 defer 语句来保证互斥锁一定会被解锁。
*/

// 不使用锁的Counter
type Counter struct {
	v map[string]int
}

func (c *Counter) Inc(key string) {
	time.Sleep(10 * time.Millisecond)
	c.v[key] = c.v[key] + 1
}

func (c *Counter) Value(key string) int {
	return c.v[key]
}

// 使用锁的Counter
type SafeCounter struct {
	v      map[string]int
	muxtex sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.muxtex.Lock()
	c.v[key]++
	c.muxtex.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.muxtex.Lock()
	defer c.muxtex.Unlock()
	return c.v[key]
}

func main() {

	{

		fmt.Println("-----不使用锁的Counter------")
		c := Counter{v: make(map[string]int)}
		for i := 0; i < 200; i++ {
			go c.Inc("somekey")
		}

		time.Sleep(time.Second)
		fmt.Println(c.Value("somekey"))

	}

	{
		fmt.Println("-----使用锁的Counter------")
		c := SafeCounter{v: make(map[string]int)}
		for i := 0; i < 200; i++ {
			go c.Inc("somekey")
		}

		time.Sleep(time.Second)
		fmt.Println(c.Value("somekey"))
	}
}
