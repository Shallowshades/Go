package main

import (
	"fmt"
)

/*
	江湖有一名无事不知，无话不说的大嘴巴，“江湖百晓生”，任何江湖中发生的事件都会被百晓生知晓，且进行广播。
	先江湖中有两个帮派，分别为：
		丐帮：黄蓉、洪七公、乔峰。
		明教：张无忌、灭绝师太、金毛狮王。
	现在需要用观察者模式模拟如下场景：
		事件一：丐帮的黄蓉把明教的张无忌揍了，这次武林事件被百晓生知晓，并且进行广播。
         		主动打人方的帮派收到消息要拍手叫好。
         		被打的帮派收到消息应该报酬，如：灭绝师太得知消息进行报仇，将丐帮黄蓉揍了。触发事件二。
		事件二：明教的灭绝师太把丐帮的黄蓉揍了，这次武林事件被百姓生知晓，并且进行广播。
*/

const (
	PGaiBang  string = "丐帮"
	PMingJiao string = "明教"
)

// ============抽象层============

// 消息
type Event struct {
	Notifier Notifier
	Positive Listener
	Negetive Listener
	Message  string
}

// 观察者
type Listener interface {
	OnFriendBeFight(event *Event)
	GetName() string
	GetParty() string
	Title() string
}

// 通知者
type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify(event *Event)
}

//============实现层============

// 具体的观察者
type Hero struct {
	Name  string
	Party string
}

func (h *Hero) GetName() string {
	return h.Name
}

func (h *Hero) GetParty() string {
	return h.Party
}

func (h *Hero) Title() string {
	return fmt.Sprintf("[%s]%s", h.Party, h.Name)
}

func (h *Hero) Fight(someone Listener, baixiao Notifier) {
	//生成事件
	event := &Event{
		Notifier: baixiao,
		Positive: h,
		Negetive: someone,
		Message:  fmt.Sprintf("%s 打了 %s...", h.Title(), someone.Title()),
	}

	//此处也是循环的关键点
	//百晓生广播消息
	baixiao.Notify(event)
}

func (h *Hero) OnFriendBeFight(event *Event) {
	//当事人
	if h.Name == event.Positive.GetName() || h.Name == event.Negetive.GetName() {
		return
	}

	//本帮派 主动者
	if h.Party == event.Positive.GetParty() {
		fmt.Println(h.Title(), "得知消息，拍手叫好...")
		return
	}

	//本帮派 被动者
	if h.Party == event.Negetive.GetParty() {
		fmt.Println(h.Title(), "得知消息,发起反击...")

		//此处，在通知者通知观察者时，若观察者又调用了通知者，将进入循环调用
		//h.Fight(event.Positive, event.Notifier)
		return
	}

}

// 百晓生 通知者
type BaiXiao struct {
	heroList []Listener
}

func (b *BaiXiao) AddListener(listener Listener) {
	b.heroList = append(b.heroList, listener)
}

func (b *BaiXiao) RemoveListener(listener Listener) {
	for i, l := range b.heroList {
		if listener == l {
			b.heroList = append(b.heroList[:i], b.heroList[i+1:]...)
			return
		}
	}
}

func (b *BaiXiao) Notify(event *Event) {
	fmt.Println("[世界消息] 百晓生广播了消息：", event.Message)
	for _, l := range b.heroList {
		l.OnFriendBeFight(event)
	}
}

//============业务层============

func main() {

	h1 := &Hero{"黄蓉", PGaiBang}
	h2 := &Hero{"洪七公", PGaiBang}
	h3 := &Hero{"乔峰", PGaiBang}
	h4 := &Hero{"张无忌", PMingJiao}
	h5 := &Hero{"灭绝师太", PMingJiao}
	h6 := &Hero{"金毛狮王", PMingJiao}

	baixiao := &BaiXiao{}

	baixiao.AddListener(h1)
	baixiao.AddListener(h2)
	baixiao.AddListener(h3)
	baixiao.AddListener(h4)
	baixiao.AddListener(h5)
	baixiao.AddListener(h6)

	fmt.Println("武林一片平静......")
	//本帮派打了本帮派，本帮派的拍手叫好 ？？！！
	//h1.Fight(h3, baixiao)

	h1.Fight(h4, baixiao)
}
