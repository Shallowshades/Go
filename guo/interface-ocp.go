// package
package main

// import
import "fmt"

/*
	OCP 开闭原则
*/

// Pet
type Pet interface {
	eat()
	sleep()
}

// Dog
type Dog struct {
}

// Cat
type Cat struct {
}

// eat
//
//	@receiver d
func (d Dog) eat() {
	fmt.Println("dog eat...")
}

// sleep
//
//	@receiver d
func (d Dog) sleep() {
	fmt.Println("dog sleep...")
}

// eat
//
//	@receiver c
func (c Cat) eat() {
	fmt.Println("cat eat...")
}

// sleep
//
//	@receiver c
func (c Cat) sleep() {
	fmt.Println("cat sleep...")
}

// +++

// Pig
type Pig struct {
}

// eat
//
//	@receiver p
func (p Pig) eat() {
	fmt.Println("pig eat...")
}

// sleep
//
//	@receiver p
func (p Pig) sleep() {
	fmt.Println("pig sleep...")
}

// Person
type Person struct {
}

// care
//
//	@receiver per
//	@param pet
func (per Person) care(pet Pet) {
	pet.eat()
	pet.sleep()
}

func main() {

	dog := Dog{}
	cat := Cat{}

	person := Person{}
	person.care(dog)
	person.care(cat)

	//+++
	pig := Pig{}
	person.care(pig)
}
