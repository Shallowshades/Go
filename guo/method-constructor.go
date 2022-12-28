package main

import (
	"fmt"
	"os"
)

type Person struct {
	name string
	age  int
}

// NewPerson
//
//	@param name
//	@param age
//	@return *Person
//	@return error
func NewPerson(name string, age int) (*Person, error) {
	if name == "" || age < 0 {
		return nil, fmt.Errorf("name is empty or age is less than zero")
	}
	return &Person{name, age}, nil
}

func main() {

	person, err := NewPerson("Cloud", -18)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("person: %v\n", person)
}
