package main

import "fmt"

// 适配目标 抽象的技能
type Attack interface {
	Fight()
}

// 具体的技能
type Sword struct{}

func (s *Sword) Fight() {
	fmt.Println("use sword to attack...")
}

type Hero struct {
	Name   string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println(h.Name, " use skill...")
	h.attack.Fight()
}

// 适配者
type PowerOff struct{}

func (p *PowerOff) ShutDown() {
	fmt.Println("PC is about to shut down...")
}

// 适配器 是适配目标的子类,包含一个适配者的成员属性
type Adapter struct {
	powerOff *PowerOff
}

func (a *Adapter) Fight() {
	a.powerOff.ShutDown()
}

func NewAdapter(p *PowerOff) *Adapter {
	return &Adapter{p}
}

func main() {

	Zack := Hero{"Zack", new(Sword)}
	Zack.Skill()

	Cloud := Hero{"Cloud", NewAdapter(new(PowerOff))}
	Cloud.Skill()
}
