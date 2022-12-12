package main

import (
	"fmt"
	"net/http"
)

/*
一个基本的 HTTP服务器有几个关键的工作需要处理。

处理动态请求：处理来自浏览网站、登录其帐户或发布图像的用户的传入请求。
提供静态资产：为浏览器提供 JavaScript、CSS 和图像，为用户创造动态体验。
接受连接： HTTP 服务器必须侦听特定端口才能接受来自 Internet 的连接。
*/
func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "welcome to my website!")
	fmt.Fprintf(w, "GET parameters : %s\n", r.URL.Query().Get("Accept-Encoding"))
	fmt.Fprintf(w, "POST parameters : %s\n", r.FormValue("email"))
}

func main() {

	//处理动态请求
	//net/http包包含一个http.HandlerFunc函数，该函数接受两个参数
	//1.path 访问路径
	//2.func 执行函数
	http.HandleFunc("/", helloHandler)

	//提供静态资产
	//http.FileServer 指定静态文件的存放路径
	fs := http.FileServer(http.Dir("static/"))
	//去除url的前缀
	http.Handle("/static", http.StripPrefix("/static/", fs))

	//接受连接
	//http.ListenAndServe
	http.ListenAndServe(":8080", nil)
}
