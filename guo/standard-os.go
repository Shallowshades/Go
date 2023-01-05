package main

import (
	"fmt"
	"os"
	"sort"
)

/*
	standard library : os
*/

// 创建文件
func createFile() {
	f, err := os.Create("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}
}

// 创建目录
func makeDir() {
	err := os.Mkdir("test", os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	err2 := os.MkdirAll("a/b/c", os.ModePerm)
	if err != nil {
		fmt.Printf("err2: %v\n", err2)
	}
}

// 删除目录或者文件
func remove() {
	err := os.Remove("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 删除目录下的所有文件或目录
func removeAll() {
	err := os.RemoveAll("test")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

func wd() {
	dir, _ := os.Getwd() // 获取工作目录
	fmt.Printf("dir: %v\n", dir)
	os.Chdir("/home/shallow/Go") //修改工作目录
	dir, _ = os.Getwd()
	fmt.Printf("dir: %v\n", dir)
}

// 文件重命名
func rename() {
	err := os.Rename("test.txt", "a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 读文件
func readFile() {
	buf, err := os.ReadFile("a.txt")
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("buf: %v\n", string(buf))
}

// 写文件
func writeFile() {
	err := os.WriteFile("a.txt", []byte("Cloud and Tifa"), os.ModePerm)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
}

// 打开关闭文件
func openCloseFile() {
	f, err := os.Open("a.txt") //只读
	defer f.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}

	f2, err2 := os.OpenFile("b.txt", os.O_RDWR|os.O_CREATE, 755)
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
	}
}

// 创建文件
func createFile2() {
	f, err := os.Create("c.txt")
	// == os.OpenFile(name, os.O_RDWR|os.O_CREATE, 666)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("f.Name(): %v\n", f.Name())
	}

	f2, err2 := os.CreateTemp("", "temp")
	if err2 != nil {
		fmt.Printf("err2: %v\n", err2)
	} else {
		fmt.Printf("f2.Name(): %v\n", f2.Name())
	}
}

// read
func readOps() {
	f, _ := os.Open("a.txt")
	defer f.Close()
	buf := make([]byte, 10)
	for {
		n, err := f.Read(buf)
		if err != nil {
			fmt.Printf("err: %v\n", err)
			break
		}
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
		sort.Slice(buf, func(i, j int) bool {
			return buf[i] < buf[j]
		})
		fmt.Printf("string(buf): %v\n", string(buf))
	}

	//从index开始read
	n, err := f.ReadAt(buf, 4)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}

	/*
		0 means relative to the origin of the file,
		1 means relative to the current offset,
		2 means relative to the end.
	*/
	f.Seek(3, 0)
	n, err = f.Read(buf)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {
		fmt.Printf("n: %v\n", n)
		fmt.Printf("string(buf): %v\n", string(buf))
	}
}

// read dir
func readDir(dir string, level int) {
	de, err := os.ReadDir(dir)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	for _, v := range de {
		for i := 0; i < level; i++ {
			fmt.Print("------")
		}
		if v.IsDir() {
			fmt.Println(v.Name() + "/")
			readDir(dir+v.Name()+"/", level+1)
		} else {
			fmt.Println(v.Name())
		}
	}
}

// write
func write() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.Write([]byte("hello golang"))
}

func writeString() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteString("hello World")
}

func writeAt() {
	f, err := os.OpenFile("a.txt", os.O_RDWR|os.O_TRUNC, 0755)
	defer f.Close()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	f.WriteAt([]byte("hello golang"), 3)
}

func main() {

	//文件操作
	{
		//createFile()
		//makeDir()
		//remove()
		//removeAll()
		//rename()
		//wd()
		// readFile()
		// writeFile()
		// readFile()
	}

	//文件读
	{
		//openCloseFile()
		//createFile2()
		//readOps()
		//readDir("../book1/", 0)
	}

	//文件写
	{
		//write()
		//writeString()
		writeAt()
	}
}
