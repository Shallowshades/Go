package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

// 练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。

var sep = flag.String("s", " ", "separator")

var (
	sha2 = flag.Bool("2", true, "use sha256")
	sha3 = flag.Bool("3", false, "use sha384")
	sha5 = flag.Bool("5", false, "use sha512")
)

func main() {
	flag.Parse()
	fmt.Println(*sha2)
	fmt.Println(*sha3)
	fmt.Println(*sha5)

	for {
		s := make([]byte, 1024, 1024)
		n, err := fmt.Scan(&s)
		if err != nil {
			fmt.Println(err)
			os.Exit(-1)
		}
		fmt.Printf("n = %d\n", n)

		switch {
		case *sha5:
			fmt.Printf("use sha512, ret = %#v\n", sha512.Sum512([]byte(s)))
		case *sha3:
			fmt.Printf("use sha384, ret = %#v\n", sha512.Sum384([]byte(s)))
		case *sha2:
			fmt.Printf("use sha256, ret = %#v\n", sha256.Sum256([]byte(s)))
		}
	}
}
