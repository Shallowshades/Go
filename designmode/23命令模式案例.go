package main

import "fmt"

/*
	练习：
	联想路边撸串烧烤场景， 有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员MM。根据命令模式，设计烤串场景。
*/

// 命令的接收者
type Cooker struct{}

func (c *Cooker) RoastChickenWings() {
	fmt.Println("The cooker is already to roast chicken wings...")
}

func (c *Cooker) RoastMutton() {
	fmt.Println("The cooker is already to roast mutton...")
}

// 抽象的命令
type Command interface {
	Roast()
}

// 具体的命令
type CommandRoastChickenWings struct {
	cooker *Cooker
}

func (cmd *CommandRoastChickenWings) Roast() {
	cmd.cooker.RoastChickenWings()
}

// 具体的命令
type CommandRoastMutton struct {
	cooker *Cooker
}

func (cmd *CommandRoastMutton) Roast() {
	cmd.cooker.RoastMutton()
}

// 命令的发出者 收集命令，并将命令的发出
type Waiter struct {
	cmdList []Command
}

// 收集命令的方法
func (w *Waiter) Collection(cmd Command) {
	w.cmdList = append(w.cmdList, cmd)
}

// 发出命令的方法
func (w *Waiter) Notify() {
	if w.cmdList == nil {
		return
	}

	for _, cmd := range w.cmdList {
		cmd.Roast()
	}
}

func main() {

	//大厨
	cooker := new(Cooker)

	//食客
	cmdRoastChickenWings := CommandRoastChickenWings{cooker}
	cmdRoastMutton := CommandRoastMutton{cooker}

	//服务员
	waiter := new(Waiter)
	waiter.Collection(&cmdRoastChickenWings)
	waiter.Collection(&cmdRoastMutton)

	//服务员告之大厨食客要吃些什么
	waiter.Notify()
}
