package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	//使用递归无法确定何时关闭信道
	st := make([]*tree.Tree, 0) //模拟栈遍历tree
	for len(st) > 0 || t != nil {
		for t != nil {
			st = append(st, t)
			t = t.Left
		}
		ch <- st[len(st)-1].Value
		t = st[len(st)-1].Right
		st = st[:len(st)-1]
	}
	close(ch)
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	c1, c2 := make(chan int), make(chan int)
	go Walk(t1, c1)
	go Walk(t2, c2)
	for {
		v1, ok1 := <-c1
		v2, ok2 := <-c2
		if ok1 && ok2 { //信通同时存在
			if v1 != v2 {
				return false
			}
		} else if ok1 || ok2 { //节点数不同
			return false
		} else { //节点数相同
			break
		}
	}
	return true
}

func main() {
	/*
		练习：等价二叉查找树
		不同二叉树的叶节点上可以保存相同的值序列。

		使用 Go 的并发和信道来编写一个简单的解法。

		本例使用了 tree 包，它定义了类型：

		type Tree struct {
		    Left  *Tree
		    Value int
		    Right *Tree
		}
	*/
	{
		t := tree.New(1)
		fmt.Println(t)
		ch := make(chan int)
		go Walk(t, ch)
		for {
			val, ok := <-ch
			if ok == false {
				break
			}
			fmt.Println(val)
		}

		if Same(tree.New(1), tree.New(1)) {
			fmt.Println("Ok!")
		} else {
			fmt.Println("Fault!")
		}
		fmt.Println(Same(tree.New(1), tree.New(1)))
		fmt.Println(Same(tree.New(1), tree.New(2)))
	}
}
