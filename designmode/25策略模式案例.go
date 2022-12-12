package main

import "fmt"

/*
练习：
	商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景
*/

// 抽象的策略
type SellStrategy interface {
	GetPrice(price float64) float64
}

// 具体的策略
type StrategyA struct{}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("use strategy A discount 2...")
	return 0.8 * price
}

type StrategyB struct{}

func (sb *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("use strategy B 200 negetive 100...")
	if price >= 200 {
		return price - 100
	}
	return price
}

// 消费者 环境类
type Customer struct {
	strategy SellStrategy
}

// 设置消费策略
func (c *Customer) SetStrategy(s SellStrategy) {
	c.strategy = s
}

// 使用消费策略
func (c *Customer) Buy(price float64) float64 {
	return c.strategy.GetPrice(price)
}

// 商品 环境类
type Goods struct {
	Price    float64
	strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	g.strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("old price is ", g.Price)
	return g.strategy.GetPrice(g.Price)
}

func main() {

	{
		customer := new(Customer)
		customer.SetStrategy(&StrategyA{})
		cost := customer.Buy(250.00)
		fmt.Println("use A cost ", cost)

		customer.SetStrategy(&StrategyB{})
		cost = customer.Buy(250.00)
		fmt.Println("use B cost ", cost)
	}

	{
		nike := Goods{
			Price: 200.0,
		}
		//上午 商场执行A策略
		nike.SetStrategy(&StrategyA{})
		fmt.Println("buy the goods in the morning cost ", nike.SellPrice())

		//下午 商场执行B策略
		nike.SetStrategy(&StrategyB{})
		fmt.Println("buy the goods in the afternoon cost ", nike.SellPrice())
	}
}
