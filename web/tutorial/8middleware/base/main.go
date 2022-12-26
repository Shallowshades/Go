package main

import (
	"fmt"
	"log"
	"net/http"
)

// middle ware
// 中间件只需将http.HandlerFunc作为其参数之一，
// 将其包装并返回一个新http.HandlerFunc的供服务器调用。
func logging(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		f(w, r)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "foo")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "bar")

}

func main() {

	http.HandleFunc("/foo", logging(foo))
	http.HandleFunc("/bar", logging(bar))

	http.ListenAndServe(":8080", nil)

}
