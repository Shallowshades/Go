package main

import (
	"flag"
	"fmt"
	"time"
)

type Value interface {
	String() string
	Set(string) error
}

type Celsius float64

type Fahrenheit float64

// 练习 7.6： 对tempFlag加入支持开尔文温度。
type Kelvin float64

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func KToC(k Kelvin) Celsius {
	return Celsius(k - 273.15)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%g°C", c)
}

type celsiusFlag struct{ Celsius }

func (f *celsiusFlag) Set(s string) error {
	var unit string
	var value float64
	fmt.Sscanf(s, "%f%s", &value, &unit)
	switch unit {
	case "c", "C", "°C":
		f.Celsius = Celsius(value)
		return nil
	case "f", "F", "°F":
		f.Celsius = FToC(Fahrenheit(value))
		return nil
	case "k", "K", "°K":
		f.Celsius = KToC(Kelvin(value))
		return nil
	}
	return fmt.Errorf("invalid temperature %q", s)
}

func CelsiusFlag(name string, value Celsius, usage string) *Celsius {
	f := celsiusFlag{value}
	flag.CommandLine.Var(&f, name, usage)
	return &f.Celsius
}

var period = flag.Duration("period", 1*time.Second, "sleep period")

var temp = CelsiusFlag("temp", 20.0, "the temperature")

func main() {

	flag.Parse()
	//fmt.Printf("Sleeping for %v...\n", *period)
	//time.Sleep(*period)
	//fmt.Println("end...")

	fmt.Println(*temp)
}

//练习 7.7： 解释为什么帮助信息在它的默认值是20.0没有包含°C的情况下输出了°C。
//A:调用了Celsius的String方法进行的输出
