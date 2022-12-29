# 模板模式



## 1. 概念

模板模式又叫模版方法模式（Template Method Pattern），是指定义一个算法的骨架，并允许子类为一个或多个步骤提供实现。模板模式使得子类可以在不改变算法结构的情况下，重新定义算法的某些步骤，属于行为型设计模式。



## 2. 理解

**当要做一件事儿的时候，这件事儿的流程和步骤是固定好的，但是每一个步骤的具体实现方式是不一定的。这个时候就可以使用模板模式**。



## 3. 类图

![](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/20221229150635.png)



## 4. 实现

比如我们去银行柜台办理业务，存款、取款、购买理财等这些业务的流程中都会有：取号、排位等号、处理业务、服务评价这几个步骤，如果你是金葵花之类的 VIP 用户，有可能有专属窗口不用排队。

我们先定义该业务包含的步骤方法：

```go
// BankBusinessHandler 银行柜台办理业务各个流程
type BankBusinessHandler interface {
	TakeRowNumnber()
	WaitInHead()
	HandleBussiness()
	Commentate()
	checkVipIdentify() bool
}
```

再定义业务办理流程：

```go

type BankBusinessExecutor struct {
	handler BankBusinessHandler
}

// ExecuteBankBusiness 银行柜台办理业务全流程
func (b *BankBusinessExecutor) ExecuteBankBusiness() {
	b.handler.TakeRowNumnber()
	if !b.handler.checkVipIdentify() {
		b.handler.WaitInHead()
	}
	b.handler.HandleBussiness()
	b.handler.Commentate()
}

func NewBankBusinessExecutor(businessHandler BankBusinessHandler) *BankBusinessExecutor {
	return &BankBusinessExecutor{handler: businessHandler}
}
```

提供一个默认实现：

```go

// DefaultBusinessHandler 默认实现
type DefaultBusinessHandler struct{}

func (d *DefaultBusinessHandler) TakeRowNumnber() {
	fmt.Println("请拿好您的取件码：" + strconv.Itoa(rand.Intn(100)) +
		" ，注意排队情况，过号后顺延三个安排")
}

func (d *DefaultBusinessHandler) WaitInHead() {
	fmt.Println("排队等号中...")
	time.Sleep(5 * time.Second)
	fmt.Println("请去窗口xxx...")
}

func (d *DefaultBusinessHandler) Commentate() {
	fmt.Println("请对我的服务作出评价，满意请按0，满意请按0，(～￣▽￣)～")
}

func (d *DefaultBusinessHandler) checkVipIdentify() bool {
	return false
}
```

这里我们以“取款”业务为例子，使用 `组合` 将默认实现组合过来，这样就只需要再实现 `HandleBusiness()` 方法：

```go
// DepositBusinessHandler 取款
type DepositBusinessHandler struct {
	*DefaultBusinessHandler
	userVip bool
}

func (d *DepositBusinessHandler) HandleBussiness() {
	fmt.Println("处理业务中：取款....")
}
```

使用效果：

```go
func main() {
	dh := &DepositBusinessHandler{
		DefaultBusinessHandler: &DefaultBusinessHandler{},
		userVip:                false,
	}
	bbe := NewBankBusinessExecutor(dh)
	bbe.ExecuteBankBusiness()
}
```



## 5. 总结

优点：

1. 封装不变部分，扩展可变部分；
2. 提取公共代码，便于维护；
3. 行为由父类控制，子类实现。

缺点：

1. 每一个不同的实现都需要一个子类来实现，导致类的个数增加，使得系统更加庞大。



## 参考：

- [用Go学设计模式-提炼流程，减少重复开发就靠它了!](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247496813&idx=1&sn=fda31b59530deb7f2c05cff20bdc35c2&chksm=fa8325facdf4aceca2c7e4f0e38b416ea74f9b16231c06724b919ef6823de306fd3cf0c4303a&scene=178&cur_album_id=2531498848431669249#rd)
