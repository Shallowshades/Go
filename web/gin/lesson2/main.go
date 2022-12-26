package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gender string
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	//2.解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("fail to parse tmpl file, err : ", err)
		return
	}
	//3.渲染模板
	// str := "Cloud"
	// user := User{
	// 	Name:   "Cloud",
	// 	Gender: "female",
	// 	Age:    18,
	// }
	m := map[string]interface{}{
		"Name":   "Cloud",
		"Gender": "female",
		"Age":    20,
	}
	err = t.Execute(w, m) //data any, 可以传入任意数据类型
	if err != nil {
		fmt.Println("fail to execute tmpl file, err : ", err)
		return
	}
}

func main() {

	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("HTTP server start failed, err : ", err)
		return
	}
}
