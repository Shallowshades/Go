package main

import "fmt"

type Opener interface {
	Open()
}

type Closer interface {
	Close()
}

type Door struct{}

func (d *Door) Open() {
	fmt.Println("Open the Door")
}

func (d *Door) Close() {
	fmt.Println("Close the Door")
}

type WoodenDoor struct {
	Door
	Color string
	x, y  int
}

func NewDoor(color string, x, y int) *WoodenDoor {
	return &WoodenDoor{
		Color: color,
		x:     x,
		y:     y,
	}
}

func (d *WoodenDoor) Open() {
	fmt.Println("Open the WoodenDoor")
}

func (d *WoodenDoor) Close() {
	fmt.Println("Close the WoodenDoor")
}

func main() {

	// var o Opener
	// var c Closer
	// var d *Door

	// d = NewDoor("Black", 100, 200)
	// o = d
	// c = d

	// o.Open()
	// c.Close()

	var o Opener
	var c Closer

	o = &Door{}
	c = &Door{}
	o.Open()
	c.Close()

	o = NewDoor("Black", 1, 2)
	c = NewDoor("Black", 1, 2)

	o.Open()
	c.Close()
}
