package main

import "fmt"

/*
Stringer
	fmt 包中定义的 Stringer 是最普遍的接口之一。

type Stringer interface {
	String() string
}
	Stringer 是一个可以用字符串描述自己的类型。
	fmt 包（还有很多包）都通过此接口来打印值。
*/

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

/*
练习：Stringer

	通过让 IPAddr 类型实现 fmt.Stringer 来打印点号分隔的地址。
	例如，IPAddr{1, 2, 3, 4} 应当打印为 "1.2.3.4"。
*/
type IPAddr [4]byte

func (addr IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", addr[0], addr[1], addr[2], addr[3])
}

func main() {

	{
		a := Person{"Arthur Dent", 42}
		z := Person{"Zaphod Beeblebrox", 9001}
		fmt.Println(a, z)
	}

	{
		hosts := map[string]IPAddr{
			"loopback":  {127, 0, 0, 1},
			"googleDNS": {8, 8, 8, 8},
		}

		for name, ip := range hosts {
			fmt.Printf("%v: %v\n", name, ip)
		}
	}
}
