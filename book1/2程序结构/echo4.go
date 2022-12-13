package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var sep = flag.String("s", " ", "separator")

func main() {

	flag.Parse()
	//如果在flag.Parse函数解析命令行参数时遇到错误，默认将打印相关的提示信息，然后调用os.Exit(2)终止程序。

	newFunction()
}

func newFunction() {
	fmt.Print(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println("")
	}
}
