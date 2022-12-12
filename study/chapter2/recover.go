package main

import "fmt"

func fa() {
	fmt.Println("func A")
}

/*
recover()必须搭配defer使用。
defer一定要在可能引发panic的语句之前定义。
*/

func fb() {
	defer func() {
		err := recover()
		//如果程序出现了panic错误，可以通过recover恢复过来
		if err != nil {
			fmt.Println("recover in B")
		}
	}()
	panic("panic in B")
}

func fc() {
	fmt.Println("func C")
}

func main() {

	fa()
	fb()
	fc()

}
