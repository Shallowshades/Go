package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

var instance *singleton

func (s *singleton) SomeThing() {
	fmt.Println("a certain function...")
}

/*
	懒汉式
	线程不安全（多线程并发时可能创建多个实例）
*/

func GetInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
		return instance
	}

	return instance
}

/*
	第一种改进：加锁
	但是当并发量大的时候，竞争造成性能下降
*/

var lock sync.Mutex

func GetInstanceThreadSafe() *singleton {
	//互斥锁，保证线程安全
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
	}

	return instance
}

/*
	第二种改进：原子操作
	借助"sync/atomic"来进行内存的状态存留来做互斥
	atomic就可以自动加载和设置标记
*/

var initialized uint32

func GetInstanceFinalVersion() *singleton {

	//如果标记被设置，直接返回，不加锁
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	//第一次加锁
	lock.Lock()
	defer lock.Unlock()

	if initialized == 0 {
		instance = new(singleton)
		//设置标记位
		atomic.StoreUint32(&initialized, 1)
	}

	return instance
}

/*
	Go 内部实现原子操作的一个函数，
	若第一次进入，则执行该函数
*/

var once sync.Once

func GetInstanceOnceVersion() *singleton {
	once.Do(func() {
		instance = new(singleton)
	})

	return instance
}

/*
once源码
func (o *Once) Do(f func()) {
	//判断是否执行过该方法，如果执行过则不执行
    if atomic.LoadUint32(&o.done) == 1 {
        return
    }
    // Slow-path.
    o.m.Lock()
    defer o.m.Unlock()
    if o.done == 0 {
        defer atomic.StoreUint32(&o.done, 1)
        f()
    }
}
*/

func main() {

	// {
	// 	s := GetInstance()
	// 	s.SomeThing()

	// 	s1 := GetInstance()
	// 	s1.SomeThing()

	// 	if s == s1 {
	// 		fmt.Println("s == s1")
	// 	} else {
	// 		fmt.Println("s != s1")
	// 	}
	// }

	// {
	// 	s := GetInstanceThreadSafe()
	// 	s.SomeThing()

	// 	s1 := GetInstanceThreadSafe()
	// 	s1.SomeThing()

	// 	if s == s1 {
	// 		fmt.Println("s == s1")
	// 	} else {
	// 		fmt.Println("s != s1")
	// 	}
	// }

	// {
	// 	s := GetInstanceFinalVersion()
	// 	s.SomeThing()

	// 	s1 := GetInstanceFinalVersion()
	// 	s1.SomeThing()

	// 	if s == s1 {
	// 		fmt.Println("s == s1")
	// 	} else {
	// 		fmt.Println("s != s1")
	// 	}
	// }

	{
		s := GetInstanceOnceVersion()
		s.SomeThing()

		s1 := GetInstanceOnceVersion()
		s1.SomeThing()

		if s == s1 {
			fmt.Println("s == s1")
		} else {
			fmt.Println("s != s1")
		}
	}

}
