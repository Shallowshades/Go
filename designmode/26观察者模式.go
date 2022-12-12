package main

import "fmt"

//==============抽象层==============

// 抽象的观察者
type Listener interface {
	Status()
	OnTeacherComing()
}

// 抽象的被观察者 通知者
type Notifier interface {
	AddListener(listener Listener)
	RemoveListener(listener Listener)
	Notify()
}

//==============实现层==============

// 观察者 学生
type Student struct {
	Name     string
	Badthing string
}

func (s *Student) Status() {
	fmt.Println(s.Name, " start to ", s.Badthing)
}

func (s *Student) OnTeacherComing() {
	fmt.Println(s.Name, " stop to ", s.Badthing)
}

// 被观察者 班长
type Monitor struct {
	listenerList []Listener
}

func (m *Monitor) AddListener(listener Listener) {
	m.listenerList = append(m.listenerList, listener)
}

func (m *Monitor) RemoveListener(listener Listener) {
	for i, l := range m.listenerList {
		if listener == l {
			m.listenerList = append(m.listenerList[:i], m.listenerList[i+1:]...)
			return
		}
	}
}

func (m *Monitor) CheckStudentStatus() {
	for _, listener := range m.listenerList {
		listener.Status()
	}
}

func (m *Monitor) Notify() {
	for _, listener := range m.listenerList {
		listener.OnTeacherComing()
	}
}

//==============业务层==============

func main() {

	//观察者
	s1 := &Student{"Zhangsan", "Copy Homework"}
	s2 := &Student{"Lisi", "Play Game"}
	s3 := &Student{"Wangwu", "See and Point Wangwu Play Game"}

	//被观察者
	monitor := new(Monitor)
	//注册
	monitor.AddListener(s1)
	monitor.AddListener(s2)
	monitor.AddListener(s3)

	fmt.Println("When teacher isn't coming...")
	monitor.CheckStudentStatus()

	fmt.Println("When teacher is coming...")
	monitor.Notify() //通知

}
