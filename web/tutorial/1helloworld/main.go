package main

import (
	"fmt"
	"net/http"
)

/*
http.ResponseWriter 	写入text/html响应
http.Request			本次请求的信息，URL，标头等
*/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
}

func main() {

	//将请求处理程序注册到默认HTTP Server
	http.HandleFunc("/", helloHandler)

	//侦听HTTP链接
	http.ListenAndServe(":8080", nil)
}
