package eventbus

import (
	"fmt"
	"reflect"
	"sync"
)

// Bus 消息总线接口
type Bus interface {
	Subscribe(topic string, handler interface{}) error
	Publish(topic string, args ...interface{})
}

// AsyncEventBus 异步事件总线
type AsyncEventBus struct {
	handlers map[string][]reflect.Value
	lock     sync.Mutex
}

// NewAsyncEventBus 创建一个新的异步事件总线
func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: map[string][]reflect.Value{},
		lock:     sync.Mutex{},
	}
}

func (a *AsyncEventBus) Subscribe(topic string, handler interface{}) error {
	a.lock.Lock()
	defer a.lock.Unlock()

	v := reflect.ValueOf(handler)
	if v.Type().Kind() != reflect.Func {
		return fmt.Errorf("handler is not a function")
	}

	handler, ok := a.handlers[topic]
	if !ok {
		handler = []reflect.Value{}
	}
	handler = append(handler.([]reflect.Value), v)
	a.handlers[topic] = handler.([]reflect.Value)
	return nil
}

func (a *AsyncEventBus) Publish(topic string, args ...interface{}) {
	handlers, ok := a.handlers[topic]
	if !ok {
		fmt.Printf("not found handler in topoc: %s\n", topic)
		return
	}

	params := make([]reflect.Value, len(args))
	for i, arg := range args {
		params[i] = reflect.ValueOf(arg)
	}
	for _, handler := range handlers {
		go handler.Call(params)
	}
}
