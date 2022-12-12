package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

//runtime.GOMAXPROCS(256) //设置processser的最大数量

/*
	go run 生成trace.out文件
	go tool trace trace.out 可视化分析

*/

func main() {

	//创建一个trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	//调试业务
	fmt.Println("Hello,GMP")

	//停止trace
	trace.Stop()

}
