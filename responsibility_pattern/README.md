# 责任链模式



## 1. 概念

责任链模式（Chain of responsibility）是一种行为型设计模式。使用这个模式，我们能为请求创建一条由多个处理器组成的链路，每个处理器各自负责自己的职责，相互之间没有耦合，完成自己任务后请求对象即传递到链路的下一个处理器进行处理。

职责链在很多流行框架里都有被用到，像中间件、拦截器等框架组件都是应用的这种设计模式，这两个组件大家应该用的比较多。在做Web 接口开发的时候，像记录访问日志、解析 Token、格式化接口响应的统一结构这些类似的项目公共操都是在中间件、拦截器里完成的，这样就能让这些基础操作与接口的业务逻辑进行解耦。

## 2. 理解

责任链的核心就是将一件事情分解成多个独立的责任事件来执行，这样**能够让我们无痛地扩展业务流程的步骤**。

例如某电商的第一版购物下单逻辑可能是这样的：

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20230322140934598.png)

现在它想在**生成订单**这一步之前加入**优惠券**的相关逻辑。如果之前的逻辑耦合在一起，那么我们就需要在原有逻辑上加 if else 分支，然后需要把整个流程全部再测一遍。并且有了上面的经验我们也应该知道这个流程以后肯定还会扩展，比如再给你加上社群砍一刀、拼单这些功能，以后每次在订单生成流程中加入步骤都得修改已经写好的代码，怕不怕？

如果使用责任链模式，则我们就可以非常方便地再加入两个步骤，然后嵌入在责任链中，也可以单独测试。

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20230322141226472.png)



## 3. 类图

实现责任链模式的对象最起码需要包含如下特性：

实现责任链模式的对象最起码需要包含如下特性：

- 成员属性

- - `nextHandler`: 下一个等待被调用的对象实例

- 成员方法

- - `SetNext`: 把下一个对象的实例绑定到当前对象的`nextHandler`属性上；
  - `Do`: 当前对象业务逻辑入口，他是每个处理对象实现自己逻辑的地方；
  - `Execute`: 负责职责链上请求的处理和传递；它会调用当前对象的`Do`，`nextHandler`不为空则调用`nextHandler.Do`；

类图如下所示：

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20230322141333235.png)

不过事实上 `SetNext` 和 `Execute` 都是通用的，所以我们可以写个抽象类来实现  `SetNext` 和 `Execute`，这样每个独立步骤就只需要再实现 `Do` 就可以了。

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20230322141525410.png)

## 4. 实现

我们以一个看病的例子来实现责任链模式。具体流程如下：**挂号—>诊室看病—>收费处缴费—>药房拿药**。

### 4.1 定义责任接口

```go
type Handler interface {
	Execute(*patient) error
	SetNext(Handler) Handler
	Do(*patient) error
}
```

### 4.2 公共实现

`SetNext` 和 `Execute` 都是一样的，提供一个公共实现，具体的责任对象直接组合这个公共实现，然后单独实现 `Do` 就可以了。

```go
type Next struct {
	nextHandler Handler
}

func (n *Next) SetNext(handler Handler) Handler {
	n.nextHandler = handler
	return handler
}

func (n *Next) Execute(patient *patient) error {
	if n.nextHandler != nil {
		if err := n.nextHandler.Do(patient); err != nil {
			return err
		}
		return n.nextHandler.Execute(patient)
	}
	return nil
}
```

### 4.3 场景责任具体实现

先针对场景定义一个责任结构体：

```go
// patient defines the process of seeing a patient
type patient struct {
	Name              string
	ReceptionDone     bool
	DockerCheckUpDone bool
	MedicineDone      bool
	PaymentDone       bool
}
```

实现具体步骤：

```go
type Start struct {
	Next
}

func (s *Start) Do(p *patient) error {
	// Start dose nothing, just used as the first handler to transfer the request to Next Handler
	return nil
}

type Reception struct {
	Next
}

func (r *Reception) Do(p *patient) error {
	if p.ReceptionDone {
		return nil
	}
	fmt.Println("Reception...")
	p.ReceptionDone = true
	return nil
}

type DockerCheck struct {
	Next
}

func (d *DockerCheck) Do(p *patient) error {
	if p.DockerCheckUpDone {
		return nil
	}
	fmt.Println("docker check...")
	p.DockerCheckUpDone = true
	return nil
}

type Payment struct {
	Next
}

func (p *Payment) Do(p2 *patient) error {
	if p2.PaymentDone {
		return nil
	}
	fmt.Println("payment...")
	p2.PaymentDone = true
	return nil
}

type Medicine struct {
	Next
}

func (m Medicine) Do(p *patient) error {
	if p.MedicineDone {
		return nil
	}
	fmt.Println("medicine...")
	p.MedicineDone = true
	return nil
}
```

### 4.4 使用

```go
func main() {
	start := Start{}
	p := &patient{Name: "abc"}

	// set the chains
	start.SetNext(&Reception{}).
		SetNext(&DockerCheck{}).
		SetNext(&Payment{}).
		SetNext(&Medicine{})

	// execute
	if err := start.Execute(p); err != nil {
		fmt.Println("error:", err)
		return
	}
	fmt.Println("success")
}
```



## 5. 场景

责任链模式（Chain of Responsibility Pattern）通常用于处理请求和避免请求发送者和接收者之间的耦合，使多个对象都有机会处理请求，从而避免请求发送者与接收者之间的直接耦合关系。具体来说，当请求需要经过多个对象处理时，可以使用责任链模式。

以下是一些可能的应用场景：

1. 订单处理：在电商系统中，订单需要依次通过多个处理环节，如库存检查、价格计算、支付处理、配送等。可以使用责任链模式，将每个处理环节封装为一个对象，由这些对象依次处理订单，直到订单处理完成。
2. 异常处理：在程序运行过程中，可能会遇到多种异常情况，需要依次处理。可以使用责任链模式，将每种异常情况封装为一个对象，由这些对象依次处理异常，直到异常处理完成。
3. 日志记录：在系统中，可能需要记录多种日志信息，如操作日志、错误日志、调试日志等。可以使用责任链模式，将每种日志信息封装为一个对象，由这些对象依次处理日志信息，直到日志记录完成。
4. 请求过滤：在 Web 应用程序中，可能需要对请求进行过滤，如 XSS 过滤、CSRF 防御、登录验证等。可以使用责任链模式，将每种请求过滤封装为一个对象，由这些对象依次处理请求，直到请求处理完成。
5. 消息传递：在分布式系统中，可能需要将消息传递给多个节点，由这些节点依次处理消息。可以使用责任链模式，将每个节点封装为一个对象，由这些对象依次处理消息，直到消息处理完成。

在游戏服务器开发中，责任链模式也有很多具体的应用场景，以下是一些常见的场景：

1. 消息处理：游戏服务器需要处理玩家发送的各种消息，如登录消息、聊天消息、道具使用消息等。可以使用责任链模式，将每种消息类型封装为一个对象，由这些对象依次处理消息，直到消息处理完成。
2. 角色属性计算：游戏服务器需要根据角色的属性值计算各种游戏逻辑，如战斗力计算、技能伤害计算、装备属性加成等。可以使用责任链模式，将每种属性计算封装为一个对象，由这些对象依次处理角色属性值，直到角色属性计算完成。
3. 技能触发：游戏服务器需要根据玩家的操作触发各种技能效果，如攻击技能、加血技能、群体攻击技能等。可以使用责任链模式，将每种技能效果封装为一个对象，由这些对象依次处理技能触发，直到技能效果生效。
4. AI 行为树：游戏服务器需要为 NPC 设计各种 AI 行为，如巡逻、攻击、逃跑等。可以使用责任链模式，将每种行为封装为一个对象，由这些对象依次处理 NPC 的行为，直到 NPC 行为决策完成。
5. 道具处理：游戏服务器需要处理各种道具的使用效果，如药品回复生命、装备增加属性等。可以使用责任链模式，将每种道具效果封装为一个对象，由这些对象依次处理道具的使用效果，直到道具使用完成。



## 6. 总结

总之，责任链模式可以用于处理各种复杂的逻辑和流程，将其分解为多个处理环节，使系统更加灵活和可扩展。



## 参考

- [你也是业务开发？提前用这个设计模式预防产品加需求吧](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497039&idx=1&sn=ee5ef2ca2a378e9836564da0f2eae485&chksm=fa8324d8cdf4adce942debfe07b76f656bc3963ec9a70192d0195a9e2b2df9e15589e4630438&cur_album_id=2531498848431669249&scene=189#wechat_redirect)
- ChatGPT
