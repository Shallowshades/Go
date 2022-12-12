package main

import "fmt"

// 抽象的策略
type WeaponStrategy interface {
	UseWeapon()
}

// 具体的策略
type Ak47 struct{}

func (ak *Ak47) UseWeapon() {
	fmt.Println("use Ak47 to fight...")
}

// 具体的策略
type Knife struct{}

func (k *Knife) UseWeapon() {
	fmt.Println("use knife to fight...")
}

// 环境类
type Hero struct {
	strategy WeaponStrategy
}

// 设置一个策略
func (h *Hero) SetWeaponStrategy(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.UseWeapon()
}

func main() {

	hero := Hero{}
	hero.SetWeaponStrategy(new(Ak47))
	hero.Fight()

	hero.SetWeaponStrategy(new(Knife))
	hero.Fight()
}
