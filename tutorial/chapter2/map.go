package main

import (
	"fmt"
	"strings"
)

type Vertex struct {
	Lat, Long float64
}

func main() {

	/*
		映射将键映射到值。
		映射的零值为 nil 。nil 映射既没有键，也不能添加键。
		make 函数会返回给定类型的映射，并将其初始化备用。
	*/
	{
		var m map[string]Vertex
		fmt.Printf("(%v, %T)\n", m, m)
		m = make(map[string]Vertex)
		m["Bell Labs"] = Vertex{
			40.68433, -74.39967,
		}
		fmt.Println(m["Bell Labs"])
	}

	/*
		映射的文法
		映射的文法与结构体相似，不过必须有键名
	*/
	/*
		{
			var m = map[string]Vertex{
				"Bell Labs": Vertex{40.68433, -74.39967},
				"Google":    Vertex{37.42202, -122.08408},
			}
			fmt.Println(m)
		}
	*/

	/*
		若顶级类型只是一个类型名，你可以在文法的元素中省略它。
	*/
	{
		var m = map[string]Vertex{
			"Bell Labs": {40.68433, -74.39967},
			"Google":    {37.42202, -122.08408},
		}
		fmt.Println(m)
	}

	/*
		对映射的增删改查
	*/
	{
		m := make(map[string]int)
		m["Answer"] = 42
		fmt.Println("The value : ", m["Answer"])

		m["Answer"] = 48
		fmt.Println("The value : ", m["Answer"])

		delete(m, "Answer")
		fmt.Println("The value : ", m["Answer"])

		v, ok := m["Answer"]
		fmt.Println("The value : ", v, "Present?", ok)
	}

	/*
		case :
		实现 WordCount。它应当返回一个映射，其中包含字符串s中每个单词的个数。
	*/
	{
		m := func(s string) map[string]int {
			m := make(map[string]int)
			words := strings.Fields(s)
			for _, v := range words {
				m[v]++
			}
			return m
		}("I am learning Go!")

		for k, v := range m {
			fmt.Println(k, " : ", v)
		}
	}
}
