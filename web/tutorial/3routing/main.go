package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func bookHandler(w http.ResponseWriter, r *http.Request) {
	//mux.Vars(r) 参数为http.Request, 返回映射片段map[string]string
	vars := mux.Vars(r)
	title := vars["title"]
	page := vars["page"]

	fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
}

func BookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "You've requested the book: %s\n", title)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "You're creating the book: %s\n", title)
}

func ReadBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "You're reading the book: %s\n", title)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "You're updating the book: %s\n", title)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	fmt.Fprintf(w, "You're deleting the book: %s \n", title)
}

func SecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You're using the secure protocol brower the website!")
}

func InsecureHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You're using the insecure protocol brower the website!")
}

func AllBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "You're getting all books!")
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]
	fmt.Fprintf(w, "You're getting the book : %s!\n", title)
}

func main() {

	//创建一个请求路由，主路由（接受所有的HTTP请求），作为参数传递给server
	r := mux.NewRouter()

	//将handlers注册给主路由
	//{title} {page} 占位符替换动态片段
	r.HandleFunc("/books/{title}/page/{page}", bookHandler)

	//限定请求方法
	r.HandleFunc("/books/{title}", CreateBook).Methods("POST")
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	//限定请求域名或子域名
	r.HandleFunc("/books/{title}", BookHandler).Host("192.168.29.135")

	//限定请求协议 HTTP/HTTPS
	r.HandleFunc("/secure", SecureHandler).Schemes("https")
	r.HandleFunc("/insecure", InsecureHandler).Schemes("http")

	//子路由和限定路由前缀
	bookRouter := r.PathPrefix("/books").Subrouter()
	bookRouter.HandleFunc("/", AllBooks)
	bookRouter.HandleFunc("/{title}", GetBook)

	//第一个参数 端口
	//第二个参数 路由（nil为使用net/http package中的默认路由）
	http.ListenAndServe(":8080", r)
}
