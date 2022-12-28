package main

import "fmt"

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct{}

func (f Fish) fly() {
	fmt.Println("fly...")
}

func (f Fish) swim() {
	fmt.Println("swim...")
}

func main() {

	var ff FlyFish
	ff = &Fish{}
	ff.fly()
	ff.swim()
}
