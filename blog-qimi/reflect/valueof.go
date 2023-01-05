package main

import (
	"fmt"
	"reflect"
)

func reflectValue(x any) {
	v := reflect.ValueOf(x)
	k := v.Kind()
	switch k {
	case reflect.Int64:
		// v.Int()从反射中获取整型的原始值，然后通过int64()强制类型转换
		fmt.Printf("type is int64, value is %d\n", int64(v.Int()))
	case reflect.Float32:
		fmt.Printf("type is float32, value is %f\n", float32(v.Float()))
	case reflect.Float64:
		fmt.Printf("type is float64, value is %f\n", float64(v.Float()))
	}
}

// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。
func reflectSetValue1(x any) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200) //修改的是副本，reflect包会引发panic
	}
}

func reflectSetValue2(x any) {
	v := reflect.ValueOf(x)
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(200) // 反射中使用 Elem()方法获取指针对应的值
	}
}

func main() {
	//通过反射获取值
	{
		var a float32 = 3.14
		var b int64 = 100
		reflectValue(a)
		reflectValue(b)
		// 将int类型的原始值转换为reflect.Value类型
		c := reflect.ValueOf(10)
		fmt.Printf("type c :%T\n", c)
	}
	//通过反射设置变量的值
	{
		var a int64 = 100
		//reflectSetValue1(a)  //panic
		//reflectSetValue1(&a) //无法修改
		//reflectSetValue2(a) //panic
		reflectSetValue2(&a) //success
		fmt.Println(a)
	}
	//IsNil()报告v持有的值是否为nil。v持有的值的分类必须是通道、函数、接口、映射、指针、切片之一；否则IsNil函数会导致panic
	//IsValid()返回v是否持有一个值。如果v是Value零值会返回假，此时v除了IsValid、String、Kind之外的方法都会导致panic
	{
		var a *int // *int类型空指针
		fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
		fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid()) //nil

		b := struct{}{}                                                               // 实例化一个匿名结构体
		fmt.Println("查找不存在的结构体成员:", reflect.ValueOf(b).FieldByName("abc").IsValid())  // 尝试从结构体中查找"abc"字段
		fmt.Println("查找不存在的结构体方法:", reflect.ValueOf(b).MethodByName("abc").IsValid()) // 尝试从结构体中查找"abc"方法

		c := map[string]int{}
		fmt.Println("map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("Cloud"))) // 尝试从map中查找一个不存在的键
	}
}
