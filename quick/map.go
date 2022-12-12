package main

import "fmt"

//引用传递
func printMap(cityMap map[string]string) {
	for key, value := range cityMap {
		fmt.Println("key = ", key, "\tvalue = ", value)
	}
	cityMap["China"] = "Beijing"
}

func main() {
	//声明方式
	//1.先声明，不分配内存
	//map类型 key是string，value是string
	var myMap1 map[string]string
	if myMap1 == nil {
		fmt.Println("myMap1 is an empty map")
	} else {
		fmt.Println("myMap1 is't an empty map")
	}

	//使用前先分配内存空间
	myMap1 = make(map[string]string, 10)
	myMap1["one"] = "Cloud"
	myMap1["two"] = "Alice"
	myMap1["three"] = "Tifa"
	fmt.Println(myMap1)

	//2.声明并分配
	myMap2 := make(map[int]string)
	myMap2[1] = "Cloud"
	myMap2[2] = "Alice"
	myMap2[3] = "Tifa"
	fmt.Println(myMap2)

	//3.声明并初始化
	myMap3 := map[float32]string{
		1.01: "Cloud",
		2.02: "Alice",
		3.03: "Tifa",
	}
	fmt.Println(myMap3)

	//使用方式
	cityMap := make(map[string]string)
	//添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["America"] = "NewYork"
	//遍历
	for key, value := range cityMap {
		fmt.Println("key = ", key, "\tvalue = ", value)
	}
	//删除
	delete(cityMap, "China")
	for key, value := range cityMap {
		fmt.Println("key = ", key, "\tvalue = ", value)
	}
	//修改
	cityMap["America"] = "Washington"
	for key, value := range cityMap {
		fmt.Println("key = ", key, "\tvalue = ", value)
	}
	//函数传参 引用传递
	printMap(cityMap)
	cityMap["America"] = "Washington"
	for key, value := range cityMap {
		fmt.Println("key = ", key, "\tvalue = ", value)
	}
}
