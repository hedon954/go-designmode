# 策略模式



## 1. 概念

策略模式：策略模式是一种行为型模式，它将对象和行为分开，将行为定义为 `一个行为接口` 和 `具体行为的实现`。策略模式最大的特点是行为的变化，行为之间可以相互替换。每个 if 判断都可以理解为就是一个策略。本模式使得算法可独立于使用它的用户而变化。

策略模式包含如下角色：

- `Strategy`: 抽象策略类：策略是一个接口，该接口定义若干个算法标识，即定义了若干个抽象方法；
- `Context`: 环境类 /上下文类
  - 上下文是 `依赖` 于接口的类，即上下文包含 `用策略(接口)声明的变量`；
  - 上下文 `提供一个方法`，持有一个策略类的引用，最终给客户端调用。该方法委托策略变量调用具体策略所实现的策略接口中的方法（实现接口的类重写策略(接口）中的方法，来完成具体功能）
- `ConcreteStrategy`: 具体策略类，即实现策略接口的类，实现了具体的策略实现。



## 2. 理解

策略模式利用一个叫做 `Context` 的东西，可以通过它来设置执行某项任务的策略，这样就可以在运行时动态切换策略。



## 3. 类图

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20221229193337989.png)



## 4. 实现

我们以一个支付为例子，支付的策略，有很多种，如微信支付、第三方支付。

我们先定义支付行为：

```go
// PayBehavior 定义行为
type PayBehavior interface {
	OrderPay(px *PayCtx)
}
```

然后定义策略上下文：

```go
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
```

实现具体的策略，即让每种策略都去实现 `PayBehavior` 接口：

```go
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
```

这样在使用的时候，就可以利用 `PayCtx` 来动态切换支付策略了：

```go
func main() {
	wp := &WxPay{}
	px := NewPayCtx(wp)
	px.Pay()

	tp := &ThirdPay{}
	px.setPayBehavior(tp) // 切换策略
	px.Pay()
}
```





## 5. 总结

优点：

- 策略模式提供了对“开闭原则”的完美支持，用户可以在不 修改原有系统的基础上选择算法或行为，也可以灵活地增加新的算法或行为；
- 策略模式提供了管理相关的算法族的办法。
- 策略模式提供了可以替换继承关系的办法。
- 使用策略模式可以避免使用多重条件转移语句。

缺点：

- 客户端必须知道所有的策略类，并自行决定使用哪一个策略类。
- 策略模式将造成产生很多策略类，可以通过使用享元模式在一定程度上减少对象的数量。

适用场景：

- 如果在一个系统里面有许多类，它们之间的区别仅在于它们的行为，那么使用策略模式可以动态地让一个对象在许多行 为中选择一种行为。
- 一个系统需要动态地在几种算法中选择一种。
- 如果一个对象有很多的行为，如果不用恰当的模式，这些行为就只好使用多重的条件选择语句来实现。
- 不希望客户端知道复杂的、与算法相关的数据结构，在具体策略类中封装算法和相关的数据结构，提高算法的保密性与安全性。

> 电商网站支付方式，一般分为银联、微信、支付宝，可以采用策略模式。
>
> 电商网站活动方式，一般分为满减送、限时折扣、包邮活动，拼团等可以采用策略模式。



## 参考

- [Go 程序里 if else 分支太多？试着用策略模式治理一下吧](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496925&idx=1&sn=6535ed0a8b5ad98d56c48f568c0720ae&chksm=fa83254acdf4ac5c64ea7df4ead46fca891520a52491160e1de8ea19410dd60a6d198f5fdd65&scene=178&cur_album_id=2531498848431669249#rd)
- [Java设计模式——策略模式](https://huaweicloud.csdn.net/63a56fd3b878a545459471f8.html?spm=1001.2101.3001.6661.1&utm_medium=distribute.pc_relevant_t0.none-task-blog-2~default~CTRLIST~activity-1-121032404-blog-123997917.pc_relevant_3mothn_strategy_and_data_recovery&depth_1-utm_source=distribute.pc_relevant_t0.none-task-blog-2~default~CTRLIST~activity-1-121032404-blog-123997917.pc_relevant_3mothn_strategy_and_data_recovery&utm_relevant_index=1)
