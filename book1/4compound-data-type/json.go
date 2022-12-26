package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// JavaScript对象表示法（JSON）是一种用于发送和接收结构化信息的标准协议。
// 成员Tag一般用原生字符串面值的形式书写。
// json开头键名对应的值用于控制encoding/json包的编码和解码的行为，并且encoding/...下面其它的包也遵循这个约定。
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"` //omitempty 省略，当为零值时 as, false即为零值
	Actors []string
}

var movies = []Movie{
	{
		Title: "Casablanca",
		Year:  1942,
		Color: false,
		Actors: []string{"Humphrey Bogart",
			"Ingrid Bergman"},
	},
	{
		Title:  "Cool Hand Luke",
		Year:   1967,
		Color:  true,
		Actors: []string{"Paul Newman"},
	},
	{
		Title: "Bullitt",
		Year:  1968,
		Color: true,
		Actors: []string{"Steve McQueen",
			"Jacueline Bisset"},
	},
	// ...
}

func main() {

	//将一个Go语言中类似movies的结构体slice转为JSON的过程叫编组（marshaling）。
	//编组通过调用json.Marshal函数完成
	//只有导出的结构体成员才会被编码，大写字母开头
	{
		data, err := json.Marshal(movies)
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s\n", err)
		}
		fmt.Printf("%s\n", data)
	}
	{
		data, err := json.MarshalIndent(movies, "", "	")
		if err != nil {
			log.Fatalf("JSON marshaling failed: %s\n", err)
		}
		fmt.Printf("%s\n", data)
	}

	//通过定义合适的Go语言数据结构，可以选择性地解码JSON中的成员。
	{
		data, _ := json.Marshal(movies)

		//仅解码名称
		var titles []struct{ Title string }
		if err := json.Unmarshal(data, &titles); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s\n", err)
		}
		fmt.Println(titles)

		//仅解码Actors
		var Actors []struct{ Actors []string }
		if err := json.Unmarshal(data, &Actors); err != nil {
			log.Fatalf("JSON unmarshaling failed: %s\n", err)
		}
		fmt.Println(Actors)

		// 解码
		var ms []Movie
		json.Unmarshal(data, &ms)
		fmt.Printf("ms: %+v\n", ms)
	}
}
