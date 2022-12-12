package main

import "fmt"

/*
	类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。
*/

type ClothesShop struct{}

func (cs *ClothesShop) OnShop() {
	fmt.Println("休闲的装扮...")
}

type ClothesWork struct{}

func (cw *ClothesWork) OnWork() {
	fmt.Println("工作的装扮...")

}

func main() {

	cw := new(ClothesWork)
	cw.OnWork()

	cs := new(ClothesShop)
	cs.OnShop()
}
