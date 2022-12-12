package main

import (
	"fmt"
)

/*
						变量
		(	type 					value  )     ---> pair
	static type concrete type

	static type : int、string...
	concrete type : interface所指向的具体的数据类型，系统看得见的类型
*/

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

// 具体类型
type Book struct {
}

func (this *Book) ReadBook() {
	fmt.Println("Read a Book")
}

func (this *Book) WriteBook() {
	fmt.Println("Write a Book")
}

func main() {

	{
		var a string
		//pair<static type : string, value : "Cloud">
		a = "Cloud"

		//pair<type : string, value : "Cloud">
		var allType interface{}
		allType = a
		str, _ := allType.(string)
		fmt.Println(str)
	}
	/*
		{
			// linux

			// tty:pair<type:*os.File, value:"/dev/tty"文件描述符>
			tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
			if err != nil {
				fmt.Println("open file error", err)
				return
			}

			//r:pair<type:,value:>
			var r io.Reader
			//r:pair<type:*os.File, value:"/dev/tty"文件描述符>
			r = tty

			//w:pair<type:,value:>
			var w io.Writer
			//w:pair<type:os.*File, value:"dev/tty"文件描述符>
			w = r.(io.Writer)
			w.Write([]byte("HELLO THIS is A TEST"))
		}
	*/
	{
		//b:pair<type:Book,value:Book{}地址>
		b := &Book{}

		//r:pair<type:,value:>
		var r Reader
		//r:pair<type:Book,value:book{}地址>
		r = b

		r.ReadBook()

		var w Writer
		//w:pair<type:Book, value:book{}地址>
		w = r.(Writer)
		w.WriteBook()
	}
}
