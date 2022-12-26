package main

import (
	"fmt"
	"sort"
)

/*
在Go语言中，一个map就是一个哈希表的引用

key必须是支持==比较运算符的数据类型
虽然浮点数类型也是支持相等运算符比较的，但是将浮点数用做key类型则是一个坏的想法,最坏的情况是可能出现的NaN和任何浮点数都不相等。
*/
func main() {

	//make && delete
	{
		//内置make创建
		ages1 := make(map[string]int)
		fmt.Println(ages1)
		ages2 := map[string]int{
			"cloud": 18,
			"alice": 18,
		}
		/*
			相当于
			ages := make(map[string]int)
			ages["cloud"] = 31
			ages["alice"] = 34
		*/
		fmt.Println(ages2)

		//删除和查询操作是安全的，即使这些元素不在map中也没有关系；如果一个查找失败将返回value类型对应的零值
		delete(ages2, "alice")
		fmt.Println(ages2)
		delete(ages2, "Tifa")
		ages2["Tifa"] = ages2["Tifa"] + 1 //新插入Tifa:1
		fmt.Println(ages2)

		//简单写法
		ages2["alice"]++ //新插入alice:1
		fmt.Println(ages2)

		//map中的元素并不是一个变量，因此不能对map的元素进行取址操作
		//addr := &ages2["cloud"] // compile error: cannot take address of map element
		//fmt.Println(addr)

		//Map的迭代顺序是不确定的，并且不同的哈希函数实现可能导致不同的遍历顺序
		for name, age := range ages2 {
			fmt.Println(name, age)
		}

		//使用sort包的Strings函数对字符串slice进行排序
		//一开始就知道names的最终大小，因此给slice分配一个合适的大小将会更有效
		names := make([]string, 0, len(ages2))
		for name := range ages2 {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			fmt.Println(name, ages2[name])
		}

		//map类型的零值是nil，也就是没有引用任何哈希表
		var ages3 map[string]int
		fmt.Println(ages3 == nil)
		fmt.Println(len(ages3) == 0)

		//查找、删除、len和range循环都可以安全工作在nil值的map上，它们的行为和一个空的map类似
		//但是向一个nil值的map存入元素将导致一个panic异常
		//ages3["barret"] = 21 // panic: assignment to entry in nil map

		//向map存数据前必须先创建map

		//索引取值，并判断是否存在
		age, ok := ages2["barret"]
		if !ok {
			fmt.Println("barret does not exist in ages2")
		} else {
			fmt.Println("find it")
			fmt.Println("barret = ", age)
		}
		//经常将这两个结合起来使用
		if age, ok := ages2["cloud"]; !ok {
			fmt.Println("cloud does not exist in ages2")
		} else {
			fmt.Println("find it")
			fmt.Println("cloud = ", age)
		}
		//和slice一样，map之间也不能进行相等比较；唯一的例外是和nil进行比较
		fmt.Println(ages1 == nil)
		fmt.Println(ages2 == nil)
		fmt.Println(ages3 == nil)
		//fmt.Println(ages1 == ages2) error
		fmt.Println(equal(ages1, ages2))
	}

	//Go语言中并没有提供一个set类型，但是map中的key也是不相同的，可以用map实现类似set的功能
	// {
	// 	seen := make(map[string]bool)
	// 	input := bufio.NewScanner(os.Stdin)
	// 	for input.Scan() {
	// 		line := input.Text()
	// 		if line == "q" {
	// 			break
	// 		}
	// 		if !seen[line] {
	// 			seen[line] = true
	// 			fmt.Println(line)
	// 		}
	// 	}

	// 	if err := input.Err(); err != nil {
	// 		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
	// 		os.Exit(1)
	// 	}
	// }

	{

	}

}

// 判断两个map是否相等
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

// 将slice作为key,记录提交相同的字符串列表的次数
// 使用同样的技术可以处理任何不可比较的key类型
var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}
func add(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}

// Map的value类型也可以是一个聚合类型，比如是一个map或slice
// graph将一个字符串类型的key映射到一组相关的字符串集合,它们指向新的graph的key
var graph = make(map[string]map[string]bool)

// addEdge函数惰性初始化map是一个惯用方式，也就是说在每个值首次作为key时才初始化。hasEdge函数显示了如何让map的零值也能正常工作
func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

// 即使from到to的边不存在，graph[from][to]依然可以返回一个有意义的结果
func hasEdge(from, to string) bool {
	return graph[from][to]
}
