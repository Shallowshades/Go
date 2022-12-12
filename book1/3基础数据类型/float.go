package main

import (
	"fmt"
	"math"
)

/*
因为float32的有效bit位只有23个，其它的bit位用于指数和符号；当整数大于23bit能表达的范围时，float32的表示将出现误差
IEEE754
1 8 23
1 11 52

Printf函数的%g参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，使用%e（带指数）或%f的形式打印可能更合适。所有的这三个打印形式都可以指定打印的宽度和控制打印精度。

NaN非数，一般用于表示无效的除法操作结果0/0或Sqrt(-1).

函数math.IsNaN用于测试一个数是否是非数NaN，math.NaN则返回非数对应的值。
虽然可以用math.NaN来表示一个非法的结果，但是测试一个结果是否是非数NaN则是充满风险的，因为NaN和任何数都是不相等的
在浮点数中，NaN、正无穷大和负无穷大都不是唯一的，每个都有非常多种的bit模式表示

如果一个函数返回的浮点数结果可能失败，最好的做法是用单独的标志报告失败，像这样：
func compute() (value float64, ok bool) {
    // ...
    if failed {
        return 0, false
    }
    return result, true
}

服务器必须设置Content-Type头部：
w.Header().Set("Content-Type", "image/svg+xml")
*/

func main() {

	{
		var f float32 = 16777216 // 1 << 24
		fmt.Println(f == f+1)    // "true"!

		for x := 0; x < 8; x++ {
			fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
		}

		var z float64
		fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

		nan := math.NaN()
		fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"
	}

}
