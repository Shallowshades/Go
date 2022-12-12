package main

import "fmt"

/*
	依赖于抽象(接口)，不要依赖具体的实现(类)，也就是针对接口编程。
*/

/*
将模块分为3个层次，抽象层、实现层、业务逻辑层
实现层 -> 抽象层
业务层 -> 抽象层
*/

// ============> 抽象层 <===============

// 抽象层车的接口
type Car interface {
	//抽象层车的run
	Run()
}

// 抽象层司机的接口
type Driver interface {
	//抽象层司机的drive
	Drive(car Car)
}

// ============> 实现层 <===============

type Benz struct {
}

func (benz *Benz) Run() {
	fmt.Println("Benz is running...")
}

type BWM struct {
}

func (bwm *BWM) Run() {
	fmt.Println("BWM is running...")
}

type Zhangsan struct {
}

func (zhangsan *Zhangsan) Drive(car Car) {
	fmt.Println("zhangsan is driving...")
	car.Run()
}

type Lisi struct {
}

func (lisi *Lisi) Drive(car Car) {
	fmt.Println("lisi is driving...")
	car.Run()
}

// ============> 逻辑层 <===============

func main() {

	var bwm Car
	bwm = &BWM{}

	var zhangsan Driver
	zhangsan = &Zhangsan{}

	zhangsan.Drive(bwm)

	var benz Car
	benz = &Benz{}

	var lisi Driver
	lisi = &Lisi{}

	lisi.Drive(benz)

	zhangsan.Drive(benz)
	lisi.Drive(bwm)

}
