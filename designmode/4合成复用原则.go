package main

import "fmt"

/*
	如果使用继承，会导致父类的任何变换都可能影响到子类的行为。如果使用对象组合，就降低了这种依赖关系。对于继承和组合，优先使用组合。
*/

type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("eat ...")
}

type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("sleep ...")
}

type CatC struct {
	c *Cat
}

func (cc *CatC) Sleep() {
	fmt.Println("sleep ...")
}

func NewCatC(c *Cat) *CatC {
	return &CatC{c}
}

func main() {

	cb := new(CatB)
	fmt.Printf("%p\n", &cb)
	cb.Eat()
	cb.Sleep()

	/*
		&和new的区别？？？
	*/
	//cc := NewCatC(new(Cat))
	cc := &CatC{new(Cat)}
	fmt.Printf("%p\n", &cc)
	cc.c.Eat()
	cc.Sleep()
}
