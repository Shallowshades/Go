package main

import (
	"fmt"
	_ "lib1"     //调用init，但不使用api
	mylib "lib2" //别名
	. "lib3"     //将lib3包中的api直接导入本包中，不建议使用
)

/*
init顺序
import :
	import :
		import :
			const...
			var...
			init...
		...
	...
...
const...
var...
init...
main...
*/

func init() {
	fmt.Println("main. init()...")
}

func main() {
	//lib1.Lib1Test()
	mylib.Lib2Test()
	Lib3Test()
}
