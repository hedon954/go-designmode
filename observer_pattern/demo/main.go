package main

import (
	"fmt"
)

func main() {
	sub := &SubjectImpl{}
	sub.Subscribe(&ObserverImpl1{})
	sub.Subscribe(&ObserverImpl2{})
	sub.Notify("hello")
}

// Subject 发布者
type Subject interface {
	Subscribe(observer Observer)
	Notify(msg string)
}

// Observer 订阅者
type Observer interface {
	Update(msg string)
}

// SubjectImpl 实现 Subject 接口
type SubjectImpl struct {
	observers []Observer
}

// Subscribe 添加观察者
func (s *SubjectImpl) Subscribe(observer Observer) {
	s.observers = append(s.observers, observer)
}

// Notify 发布通知
func (s *SubjectImpl) Notify(msg string) {
	for _, o := range s.observers {
		o.Update(msg)
	}
}

// ObserverImpl1 实现 Obsever 接口
type ObserverImpl1 struct{}

// Update 状态发生变化
func (o *ObserverImpl1) Update(msg string) {
	fmt.Printf("ObserverImpl1 updated: %s\n", msg)
}

// ObserverImpl2 实现 Obsever 接口
type ObserverImpl2 struct{}

// Update 状态发生变化
func (o *ObserverImpl2) Update(msg string) {
	fmt.Printf("ObserverImpl2 updated: %s\n", msg)
}
