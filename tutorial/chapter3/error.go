package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{time.Now(), "it didn't work"}
}

/*
练习：
	Sqrt 接受到一个负数时，应当返回一个非 nil 的错误值。复数同样也不被支持。
*/

type ErrNegativeSqrt float64

// 在Error方法内调用 fmt.Sprint(e) 会让程序陷入死循环。
// 可以通过先转换 e 来避免这个问题
func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number %v", float64(e))
}

// 牛顿法求平方根
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {

	if err := run(); err != nil {
		fmt.Println(err)
	}

	//practice
	{
		fmt.Println(Sqrt(2))
		fmt.Println(Sqrt(-2))
	}
}
