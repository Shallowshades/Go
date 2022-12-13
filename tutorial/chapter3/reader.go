package main

import (
	"fmt"
	"io"
	"strings"
)

/*
Reader
	io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。
	Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。
	io.Reader 接口有一个 Read 方法：
	func (T) Read(b []byte) (n int, err error)

	Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。
*/

/*
	实现一个 Reader 类型，它产生一个 ASCII 字符 'A' 的无限流。
*/

type MyReader struct{}

func (reader MyReader) Read(b []byte) (n int, err error) {
	b[0], n, err = 'A', 1, nil
	return
}

/*
练习：rot13Reader
	有种常见的模式是一个 io.Reader 包装另一个 io.Reader，然后通过某种方式修改其数据流。

	例如，gzip.NewReader 函数接受一个 io.Reader（已压缩的数据流）并返回一个同样实现了 io.Reader 的 *gzip.Reader（解压后的数据流）。

	编写一个实现了 io.Reader 并从另一个 io.Reader 中读取数据的 rot13Reader，通过应用 rot13 代换密码对数据流进行修改。

	rot13Reader 类型已经提供。实现 Read 方法以满足 io.Reader。
*/

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, err error) {
	n, err = r.r.Read(b)
	for i := range b {
		b[i] -= 13
		if b[i] < 'a' {
			b[i] += 26
		}
	}
	return
}

func main() {

	/*
		创建了一个 strings.Reader 并以每次 8 字节的速度读取它的输出。
	*/
	{
		r := strings.NewReader("Hello,Reader!")
		b := make([]byte, 8)
		for {
			n, err := r.Read(b)
			fmt.Printf("n = %v err = %v b = %c\n", n, err, b)
			fmt.Printf("b[:n] = %q\n", b[:n])
			if err == io.EOF {
				break
			}
		}
	}

	//practice
	{
		r := new(MyReader)
		b := make([]byte, 8)
		for i := 0; i < 3; i++ {
			n, err := r.Read(b)
			fmt.Printf("n = %v err = %v b = %c\n", n, err, b)
			fmt.Printf("b[:n] = %q\n", b[:n])
			if err == io.EOF {
				break
			}
		}
	}

	//
	{

	}
}
