package main

import "fmt"

//别名
type myint int

type Book struct {
	title  string
	author string
}

func changeBook1(book Book) {
	//传递一个Book的副本
	book.author = "666"
}

func changeBook2(book *Book) {
	//指针传递
	book.author = "777"
}

type Hero struct {
	Name  string
	Ad    int
	Level int
}

/*
	函数首字母大小写决定本包api对不对外界开放，大写对外开放，小写不开放
*/

func (this *Hero) Show() {
	fmt.Println("Name = ", this.Name)
	fmt.Println("Ad = ", this.Ad)
	fmt.Println("Level = ", this.Level)
}
func (this *Hero) GetName() {
	fmt.Println("Name = ", this.Name)
}
func (this *Hero) SetName(newName string) { //直接传递为值传递，指针传递方可修改
	this.Name = newName
}

type Human struct {
	name string
	sex  string
}

func (this *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (this *Human) Walk() {
	fmt.Println("Human.Walk()...")
}

type SuperMan struct {
	Human //SuperMan类继承了Human类的方法
	level int
}

//重定义父类的方法Eat()
func (this *SuperMan) Eat() {
	fmt.Println("SuperMan.Eat()...")
}

//子类的新方法
func (this *SuperMan) Fly() {
	fmt.Println("SuperMan.Fly()...")
}

func (this *SuperMan) print() {
	fmt.Println(*this)
}

//本质是一个指针
type AnimalIF interface {
	Sleep()
	GetColor() string //获取动物的颜色
	GetType() string  //获取动物的种类
}

//具体的类
type Cat struct {
	color string
}

func (this *Cat) Sleep() {
	fmt.Println("Cat is Sleep")
}
func (this *Cat) GetColor() string {
	return this.color
}
func (this *Cat) GetType() string {
	return "Cat"
}

//具体的类
type Dog struct {
	color string
}

func (this *Dog) Sleep() {
	fmt.Println("Dog is Sleep")
}
func (this *Dog) GetColor() string {
	return this.color
}
func (this *Dog) GetType() string {
	return "Dog"
}
func showAnimal(animal AnimalIF) {
	animal.Sleep()
	fmt.Println("color = ", animal.GetColor())
	fmt.Println("kind = ", animal.GetType())
}

func main() {

	//别名
	{
		var a myint = 10
		fmt.Printf("a = %v type of a = %T\n", a, a)
	}

	//结构体
	{
		var book1 Book
		book1.title = "Golang"
		book1.author = "Zhangsan"
		fmt.Printf("%v\n", book1)
		changeBook1(book1)
		fmt.Printf("%v\n", book1)
		changeBook2(&book1)
		fmt.Printf("%v\n", book1)
	}

	//类的封装
	{
		hero := Hero{Name: "zhangsan", Ad: 100, Level: 1}
		hero.Show()
		hero.SetName("lisi") //
		hero.Show()
	}

	//继承
	{
		h := Human{"Zhangsan", "female"}
		h.Eat()
		h.Walk()

		//定义一个子类
		s := SuperMan{Human{"lisi", "female"}, 88}
		/*
			var s SuperMan
			s.name = "lisi"
			s.sex = "male"
			s.level = 88
		*/
		s.Walk()
		s.Eat()
		s.Fly()
		s.print()
	}

	//多态
	{
		var animal AnimalIF
		animal = &Cat{"Green"}
		animal.Sleep()
		animal = &Dog{"Yellow"}
		animal.Sleep()

		cat := Cat{"Black"}
		dog := Dog{"White"}
		showAnimal(&cat)
		showAnimal(&dog)
	}
}
