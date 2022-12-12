package main

import "fmt"

// 抽象主题
type BeautyGirl interface {
	MakeEyesWithMan()
	HappyWithMan()
}

// 具体主题 潘
type Pan struct{}

func (p *Pan) MakeEyesWithMan() {
	fmt.Println("Make Eyes With Man")
}

func (p *Pan) HappyWithMan() {
	fmt.Println("Happy With Man")
}

// 代理人 王
type Wang struct {
	girl BeautyGirl
}

func NewProxy(girl BeautyGirl) BeautyGirl {
	return &Wang{girl}
}

// 不能对外暴露？？？ 破坏interface？
func (p *Wang) providePlace() {
	fmt.Println("provide place to girl")
}

func (p *Wang) MakeEyesWithMan() {
	p.girl.MakeEyesWithMan()
}

func (p *Wang) HappyWithMan() {
	p.providePlace()
	p.girl.HappyWithMan()
}

/*
	王婆年轻时候也是一个美女：代理也是抽象主题的一个具体主题
	王婆手里有美女的资源潘：代理又组合了一个具体对象
	王婆有自己的手段：代理除了重写了抽象接口的方法，还有自己的方法
*/

func main() {

	//西门找王婆作为中间人和潘交流
	wang := NewProxy(new(Pan))
	wang.MakeEyesWithMan()
	wang.HappyWithMan()

}
