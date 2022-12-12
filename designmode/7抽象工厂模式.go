package main

import "fmt"

// ===============抽象层===============

//产品等级结构

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象工厂
type AbstractFactory interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// ===============实现层===============

// 中国产品族
type ChinaApple struct{}

func (ca *ChinaApple) ShowApple() {
	fmt.Println("Chinese Apple...")
}

type ChinaBanana struct{}

func (cb *ChinaBanana) ShowBanana() {
	fmt.Println("Chinese Banana...")
}

type ChinaPear struct{}

func (cp *ChinaPear) ShowPear() {
	fmt.Println("Chinese Pear...")
}

type ChinaFactory struct{}

func (cf *ChinaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}

func (cf *ChinaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(ChinaBanana)
	return banana
}

func (cf *ChinaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(ChinaPear)
	return pear
}

// 日本产品族
type JapanApple struct{}

func (ja *JapanApple) ShowApple() {
	fmt.Println("Japan Apple...")
}

type JapanBanana struct{}

func (jb *JapanBanana) ShowBanana() {
	fmt.Println("Japan Banana...")
}

type JapanPear struct{}

func (jp *JapanPear) ShowPear() {
	fmt.Println("Japan Pear...")
}

type JapanFactory struct{}

func (cf *JapanFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(JapanApple)
	return apple
}

func (cf *JapanFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(JapanBanana)
	return banana
}

func (cf *JapanFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(JapanPear)
	return pear
}

// 美国产品族
type AmericaApple struct{}

func (ja *AmericaApple) ShowApple() {
	fmt.Println("America Apple...")
}

type AmericaBanana struct{}

func (jb *AmericaBanana) ShowBanana() {
	fmt.Println("America Banana...")
}

type AmericaPear struct{}

func (jp *AmericaPear) ShowPear() {
	fmt.Println("America Pear...")
}

type AmericaFactory struct{}

func (cf *AmericaFactory) CreateApple() AbstractApple {
	var apple AbstractApple
	apple = new(AmericaApple)
	return apple
}

func (cf *AmericaFactory) CreateBanana() AbstractBanana {
	var banana AbstractBanana
	banana = new(AmericaBanana)
	return banana
}

func (cf *AmericaFactory) CreatePear() AbstractPear {
	var pear AbstractPear
	pear = new(AmericaPear)
	return pear
}

// ===============逻辑层===============

func main() {

	var af AbstractFactory
	af = new(AmericaFactory)

	var aApple AbstractApple
	aApple = af.CreateApple()
	aApple.ShowApple()

	var aBanana AbstractBanana
	aBanana = af.CreateBanana()
	aBanana.ShowBanana()

	var aPear AbstractPear
	aPear = af.CreatePear()
	aPear.ShowPear()

	var cf AbstractFactory
	cf = new(ChinaFactory)

	var cApple AbstractApple
	cApple = cf.CreateApple()
	cApple.ShowApple()

	var cBanana AbstractBanana
	cBanana = cf.CreateBanana()
	cBanana.ShowBanana()

	// var cPear AbstractPear
	// cPear = cf.CreatePear()
	// cPear.ShowPear()

	cPear := cf.CreatePear()
	cPear.ShowPear()
}
