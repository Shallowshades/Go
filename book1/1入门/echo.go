package main

import (
	"fmt"
	"os"
	"strings"
)

/*
程序的命令行参数可从 os 包的 Args 变量获取
os 包外部使用 os.Args 访问该变量
*/

func main() {

	{
		var s, sep string
		for i := 0; i < len(os.Args); i++ {
			s += sep + os.Args[i]
			sep = " "
		}
		fmt.Println(s)
	}

	{
		var s, sep string
		for _, arg := range os.Args { //空标识符（blank identifier），即 _（也就是下划线）
			s += sep + arg
			sep = " "
		}
		fmt.Println(s)
	}

	{
		fmt.Println(strings.Join(os.Args, " "))
	}

	{
		fmt.Println(os.Args)
	}

	{
		for _, v := range os.Args {
			fmt.Println(v)
		}
	}

	//做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异
	{

	}
}
