package main

import (
	"fmt"
)

type Goods struct {
	Kind string
	Fact bool
}

//============抽象层============

// 抽象的购物主题
type Shopping interface {
	Buy(goods *Goods)
}

// ============实现层============

// 具体的购物主题
type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("go Korea and buy ", goods.Kind)
}

type AmericaShopping struct{}

func (as *AmericaShopping) Buy(goods Goods) {
	fmt.Println("go America and buy ", goods.Kind)
}

type AfrikaShopping struct{}

func (as *AfrikaShopping) Buy(goods Goods) {
	fmt.Println("go Afrika and buy ", goods.Kind)
}

// 海外代理
type OverseaProxy struct {
	shopping Shopping //代理某个具体购物，这个用抽象
}

// 代理功能 辨别真伪
func (op *OverseaProxy) distinguish(goods *Goods) bool {
	fmt.Println("distinguish the ", goods.Kind)
	if goods.Fact == false {
		fmt.Println("It's fake, don't buy the ", goods.Kind)
	}
	return goods.Fact
}

// 代理功能 海关检验
func (op *OverseaProxy) check(goods *Goods) {
	fmt.Println("Custom inspect and go home with ", goods.Kind)
}

// 代理功能 购买
func (op *OverseaProxy) Buy(goods *Goods) {
	//验货
	if op.distinguish(goods) == true {
		//购买
		op.shopping.Buy(goods)
		//海关
		op.check(goods)
	}
}

// 创建代理，并且配置代理的对象
func NewProxy(shopping Shopping) Shopping {
	return &OverseaProxy{shopping}
}

// ============业务层============

func main() {

	g1 := Goods{
		Kind: "Keroa Facial mask",
		Fact: true,
	}

	g2 := Goods{
		Kind: "CET4 certificate",
		Fact: false,
	}

	//不使用代理
	var shopping Shopping
	shopping = new(KoreaShopping)

	if g1.Fact == true {
		fmt.Println("distinguish...")
		shopping.Buy(&g1)
		fmt.Println("custom inspect...")
	}

	//使用代理
	var proxy Shopping
	proxy = NewProxy(shopping)
	proxy.Buy(&g1)
	proxy.Buy(&g2)

}
