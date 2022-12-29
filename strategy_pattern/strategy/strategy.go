package main

import (
	"fmt"
)

// PayBehavior 定义行为
type PayBehavior interface {
	OrderPay(px *PayCtx)
}

// PayCtx 上下文
type PayCtx struct {
	payBehavior PayBehavior // 持有策略
	payParams   map[string]interface{}
}

// setPayBehavior 设置支付模式，即选择策略
func (px *PayCtx) setPayBehavior(p PayBehavior) {
	px.payBehavior = p
}

// Pay 执行策略
func (px *PayCtx) Pay() {
	px.payBehavior.OrderPay(px)
}

func NewPayCtx(p PayBehavior) *PayCtx {
	return &PayCtx{
		payBehavior: p,
		payParams: map[string]interface{}{
			"appid":  1,
			"appkey": 2,
		},
	}
}

// WxPay 策略：微信支付
type WxPay struct{}

func (wp *WxPay) OrderPay(px *PayCtx) {
	fmt.Printf("Wx支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用Wx支付进行支付")
}

// ThirdPay 策略：第三方支付
type ThirdPay struct{}

func (*ThirdPay) OrderPay(px *PayCtx) {
	fmt.Printf("三方支付加工支付请求 %v\n", px.payParams)
	fmt.Println("正在使用三方支付进行支付")
}

func main() {
	wp := &WxPay{}
	px := NewPayCtx(wp)
	px.Pay()

	tp := &ThirdPay{}
	px.setPayBehavior(tp) // 切换策略
	px.Pay()
}
