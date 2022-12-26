package main

import (
	"fmt"
	"os"
	"strconv"

	. "conv/lenconv"
	. "conv/tempconv"
	. "conv/weighconv"
)

/*
练习 2.2： 写一个通用的单位转换程序，用类似cf程序的方式从命令行读取参数，如果缺省的话则是从标准输入读取参数，然后做类似Celsius和Fahrenheit的单位转换，长度单位可以对应英尺和米，重量单位可以对应磅和公斤等。
*/
func main() {
	for _, arg := range os.Args[1:] {
		x, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := Fahrenheit(x)
		c := Celsius(x)
		fmt.Printf("%s = %s, %s = %s\n", f, FToC(f), c, CToF(c))

		k := Kilogram(x)
		p := Pound(x)
		fmt.Printf("%s = %s, %s = %s\n", k, KToP(k), p, PToK(p))

		m := Metre(x)
		F := Foot(x)
		fmt.Printf("%s = %s, %s = %s\n", m, MToF(m), F, FToM(F))
	}
}
