package main

import (
	"fmt"
	"sync"
)

func SyncMap() {
	var m sync.Map

	m.Store("Cloud", "Zack")

	value, ok := m.Load("Cloud")
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	value, ok = m.Load("Alice")
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	m.Delete("Cloud")

	value, ok = m.Load("Cloud")
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	m.Store("Cloud", "1")
	m.Store("Alice", "2")
	m.Store("Tifa", "3")
	m.Store("Barret", "4")

	m.Range(func(key, value any) bool {
		fmt.Println("k: ", key, "v :", value)
		return true
	})

	//load and delete 加载一个值，存在则删除
	value, ok = m.LoadAndDelete("Cloud")
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	m.Range(func(key, value any) bool {
		fmt.Println("k: ", key, "v :", value)
		return true
	})

	// load or store 如果存在就load，不存在就store
	actual, loaded := m.LoadOrStore("Alice", "5")
	fmt.Println("actual : ", actual, " loaded : ", loaded)

	value, ok = m.Load("Alice")
	fmt.Printf("value: %v, ok: %v\n", value, ok)

	actual, loaded = m.LoadOrStore("Zack", "7")
	fmt.Println("actual : ", actual, " loaded : ", loaded)
}

func SyncCond() {
	var lock sync.Locker
	lock.Lock()
	defer lock.Unlock()

}

func main() {
	SyncMap()
}
