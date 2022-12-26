package main

import (
	"fmt"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {

	{
		dilbert = Employee{
			1, "Cloud", "Shenluo", "Manager", 10000, 10,
		}
		fmt.Println(dilbert)
		dilbert.Salary -= 5000
		position := &dilbert.Position
		*position = "Senior " + *position
		fmt.Println(dilbert)

		var employeeOfTheMonth *Employee = &dilbert
		employeeOfTheMonth.Position += " (proactive team player)"
		fmt.Println(dilbert)
		//(*employeeOfTheMonth).Position += " (proactive team player)"

		fmt.Println(EmployeeByID(dilbert.ManagerID).Position)
		id := dilbert.ID
		EmployeeByID(id).Salary = 0
		fmt.Printf("%#+v\n", dilbert)
	}

	//如果结构体成员名字是以大写字母开头的，那么该成员就是导出的；这是Go语言导出规则决定的。
	//一个结构体可能同时包含导出和未导出的成员。
	//一个命名为S的结构体类型将不能再包含S类型的成员：因为一个聚合的值不能包含它自身。
	//但是S类型的结构体可以包含*S指针类型的成员，这可以让我们创建递归的数据结构，比如链表和树结构等

	//结构体类型的零值是每个成员都是零值。通常会将零值作为最合理的默认值。例如，对于bytes.Buffer类型，结构体初始值就是一个随时可用的空缓存，sync.Mutex的零值也是有效的未锁定状态。有时候这种零值可用的特性是自然获得的，但是也有些类型需要一些额外的工作

	//如果结构体没有任何成员的话就是空结构体，写作struct{}。它的大小为0，也不包含任何信息，但是有时候依然是有价值的。
	//用map来模拟set时，代替map中的value，只强调key，但节约的空间有限，且语法复杂，所以通常避免使用。
	{
		seen := make(map[string]struct{}) // set of strings
		s := "jack"
		if _, ok := seen[s]; !ok {
			seen[s] = struct{}{}
			// ...first time seeing s...
		}

	}

	//结构体字面值初始化
	//结构体值也可以用结构体字面值表示，结构体字面值可以指定每个成员的值。
	{
		//1.顺序指定，需记忆，一般在包内部或者较小的结构体使用
		p1 := Point{1, 2}
		//2.键值对初始化，默认零值，可部分或全部初始化，无顺序要求
		p2 := Point{Y: 2, X: 1}
		fmt.Println(p1)
		fmt.Println(p2)

		//两种不同形式的写法不能混合使用
		//不能在外部包中用第一种顺序初始化结构体中未导出的成员
		//企图隐式使用未导出成员的行为也是不允许的
		//var _ = p.T{a: 1, b: 2}
		//var _ = p.T{1, 2}

		fmt.Println(Scale(Point{1, 2}, 5))

		//pp1 = pp2
		pp1 := &Point{1, 2}
		pp2 := new(Point)
		*pp2 = Point{1, 2}
		fmt.Println(pp1, pp2)
	}

	//结构体比较
	//如果结构体的全部成员都是可以比较的，那么结构体也是可以比较的，那样的话两个结构体将可以使用==或!=运算符进行比较
	{
		p := Point{1, 2}
		q := Point{2, 1}
		fmt.Println(p.X == q.X && p.Y == q.Y)
		fmt.Println(p == q)

		type address struct {
			hostname string
			port     int
		}
		hits := make(map[address]int)
		hits[address{"golang.org", 443}]++
	}

	//结构体嵌入和匿名成员
	{
		type Circle struct {
			X, Y, Radius int
		}
		type Wheel struct {
			X, Y, Radius, Spokes int
		}

		var w Wheel
		w.X = 8
		w.Y = 8
		w.Radius = 5
		w.Spokes = 20
	}

	//相似和重复之处,为了便于维护而将相同的属性独立出来
	{
		type Point struct {
			X, Y int
		}

		type Circle struct {
			Center Point
			Radius int
		}

		type Wheel struct {
			Circle Circle
			Spokes int
		}

		//结构体类型变的清晰了，但是访问每个成员变得繁琐
		var w Wheel
		w.Circle.Center.X = 8
		w.Circle.Center.Y = 8
		w.Circle.Radius = 5
		w.Spokes = 20
	}

	//Go语言有一个特性让我们只声明一个成员对应的数据类型而不指名成员的名字；这类成员就叫匿名成员。匿名成员的数据类型必须是命名的类型或指向一个命名的类型的指针。
	{
		type Point struct {
			X, Y int
		}
		type Circle struct {
			Point
			Radius int
		}
		type Wheel struct {
			Circle
			Spokes int
		}

		//得益于匿名嵌入的特性，可以直接访问叶子属性而不需要给出完整的路径.在访问子成员的时候可以忽略任何匿名成员部分
		var w Wheel
		w.X = 8
		w.Y = 8
		w.Radius = 5
		w.Spokes = 20

		//结构体字面值并没有简短表示匿名成员的语法
		//w1 := Wheel{8, 8, 5, 20}                       // compile error: unknown fields
		//w2 := Wheel{X: 8, Y: 8, Radius: 5, Spokes: 20} // compile error: unknown fields
		fmt.Println(w)
		//fmt.Println(w1)
		//fmt.Println(w2)

		//结构体字面值必须遵循形状类型声明时的结构
		w1 := Wheel{Circle{Point{8, 8}, 5}, 20}
		w2 := Wheel{
			Circle: Circle{
				Point: Point{
					X: 8,
					Y: 8,
				},
				Radius: 5,
			},
			Spokes: 20,
		}
		fmt.Println(w1)
		fmt.Println(w2)

		w1.X = 42
		fmt.Printf("%#v\n", w1)

		//匿名成员并不要求是结构体类型；其实任何命名的类型都可以作为结构体的匿名成员
		
		//但是为什么要嵌入一个没有任何子成员类型的匿名成员类型呢？
		//答案是匿名类型的方法集。简短的点运算符语法可以用于选择匿名成员嵌套的成员，也可以用于访问它们的方法。实际上，外层的结构体不仅仅是获得了匿名成员类型的所有成员，而且也获得了该类型导出的全部的方法。这个机制可以用于将一些有简单行为的对象组合成有复杂行为的对象。组合是Go语言中面向对象编程的核心，
	}
}

// EmployeeByID
//
//	@param id
//	@return *Employee
func EmployeeByID(id int) *Employee {
	return &dilbert
}

// 使用一个二叉树来实现一个插入排序
type tree struct {
	value       int
	left, right *tree
}

// Sort sorts values in place
func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

// appendValues appends the elements of t to values in order
// and returns the resulting slice.
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

type Point struct {
	X, Y int
}

func Scale(p Point, factor int) Point {
	return Point{p.X * factor, p.Y * factor}
}

// 考虑效率的话，较大的结构体通常会用指针的方式传入和返回
func Bonus(e *Employee, percent int) int {
	return e.Salary * percent / 100
}

// 如果要在函数内部修改结构体成员的话，用指针传入是必须的；因为在Go语言中，所有的函数参数都是值拷贝传入的，函数参数将不再是函数调用时的原始变量。
func AwardAnnualRaise(e *Employee) {
	e.Salary = e.Salary * 105 / 100
}
