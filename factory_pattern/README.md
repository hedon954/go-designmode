# 工厂模式



## 1. 简单工厂

Go 语言没有构造函数一说，所以一般会定义 `NewXXX` 函数来初始化相关类。`NewXXX` 函数返回接口时就是简单工厂模式。

考虑一个简单的应用场景，这个应用场景里会提供很多语言的打印机，他们都源于一个 Printer 接口：

```go
// Printer 简单工厂要返回的接口类型
type Printer interface {
 Print(name string) string
}
```

程序通过简单工厂向客户端提供需要的语种的打印机：

```go
func NewPrinter(lang string) Printer {
 switch lang {
 case "cn":
  return new(CnPrinter)
 case "en":
  return new(EnPrinter)
 default:
  return new(CnPrinter)
 }
}
```

目前这个场景里我们先提供两个语种的打印机，他们都是 Printer 接口的具体实现类型：

```go
// CnPrinter 是 Printer 接口的实现，它说中文
type CnPrinter struct {}

func (*CnPrinter) Print(name string) string {
 return fmt.Sprintf("你好, %s", name)
}

// EnPrinter 是 Printer 接口的实现，它说中文
type EnPrinter struct {}

func (*EnPrinter) Print(name string) string {
 return fmt.Sprintf("Hello, %s", name)
}
```

这个场景下，工厂、产品接口、具体产品类的关系可以用下面这个图表示。

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20221220101120818.png)

客户端只需要告诉工厂想要哪个语种的打印机产品，工厂就会把产品给返回给客户端：

```go
printer := NewPrinter("en")
fmt.Println(printer.Print("Bob"))
```

总结下来，简单工厂模式主要包含3个角色。

- 简单工厂：是简单工厂模式的核心，负责实现创建所有实例的内部逻辑。工厂类的创建产品类的方法可以被外界直接调用，创建所需的产品对象。
- 抽象产品：是简单工厂创建的所有对象的抽象父类/接口，负责描述所有实例的行为。
- 具体产品：是简单工厂模式的创建目标。

简单工厂的优点是，简单，缺点嘛，如果具体产品扩产，就必须修改工厂内部，增加 Case，一旦产品过多就会导致简单工厂过于臃肿，为了解决这个问题，才有了下一级别的工厂模式--工厂方法。



## 2. 工厂方法

工厂方法模式（Factory Method Pattern）又叫作多态性工厂模式，指的是定义一个创建对象的接口，但由实现这个接口的工厂类来决定实例化哪个产品类，工厂方法把类的实例化推迟到子类中进行。

在工厂方法模式中，不再由单一的工厂类生产产品，而是由工厂类的子类实现具体产品的创建。因此，当增加一个产品时，只需增加一个相应的工厂类的子类, 以解决简单工厂生产太多产品时导致其内部代码臃肿（switch … case 分支过多）的问题。

下面举个简单的例子来理解工厂方法的设计思想，考虑有这样一个生产计算器的工厂，每类计算器产品都由一个子工厂负责生产。

```go
// OperatorFactory 工厂接口，由具体工厂类来实现
type OperatorFactory interface {
 Create() MathOperator
}

// MathOperator 实际产品实现的接口--表示数学运算器应该有哪些行为
type MathOperator interface {
 SetOperandA(int)
 SetOperandB(int)
 ComputeResult() int
}
```

接下来我们定义一个 Operator 的基类：

```go
// BaseOperator 是所有 Operator 的基类
// 封装公用方法，因为Go不支持继承，具体Operator类
// 只能组合它来实现类似继承的行为表现。
type BaseOperator struct {
 operandA, operandB int
}

func (o *BaseOperator) SetOperandA(operand int) {
 o.operandA = operand
}

func (o *BaseOperator) SetOperandB(operand int) {
 o.operandB = operand
}
```

现在我们假定程序可以生产两类计算器，加法计算器和乘法计算器：

```go
// PlusOperator 实际的产品类--加法运算器
type PlusOperator struct {
 *BaseOperator
}

// ComputeResult 计算并获取结果
func (p *PlusOperator) ComputeResult() int {
 return p.operandA + p.operandB
}

// MultiOperator 实际的产品类--乘法运算器
type MultiOperator struct {
 *BaseOperator
}
func (m *MultiOperator) ComputeResult() int {
 return m.operandA * m.operandB
}
```

要创建这两种计算器，那么在工厂方法模式中，我们需要存在两个子类工厂，由它们来创建计算机，如下：

```go

// PlusOperatorFactory 是 PlusOperator 加法运算器的工厂类
type PlusOperatorFactory struct{}

func (pf *PlusOperatorFactory) Create() MathOperator {
 return &PlusOperator{
  BaseOperator: &BaseOperator{},
 }
}

// MultiOperatorFactory 是 MultiOperator 产品的工厂
type MultiOperatorFactory struct {}

func (mf *MultiOperatorFactory) Create() MathOperator{
 return &MultiOperator{
  BaseOperator: &BaseOperator{},
 }
}
```

>注意：这里为了理解，例子都很简单，真实场景下每个子类工厂创建产品实例的时候是可以放进去复杂逻辑的，不是简单的 New 一下。

这个场景下，工厂、产品接口、具体产品类的关系可以用下面这个图表示。

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20221220101107265.png)

简单使用方式如下：

```go
func main() {
 var factory OperatorFactory
 var mathOp MathOperator
 factory = &PlusOperatorFactory{}
 mathOp = factory.Create()
 mathOp.SetOperandB(3)
 mathOp.SetOperandA(2)
 fmt.Printf("Plus operation reuslt: %d\n", mathOp.ComputeResult())

 factory= &MultiOperatorFactory{}
 mathOp = factory.Create()
 mathOp.SetOperandB(3)
 mathOp.SetOperandA(2)
 fmt.Printf("Multiple operation reuslt: %d\n", mathOp.ComputeResult())
}
```

工厂方法模式的优点

- 灵活性增强，对于新产品的创建，只需多写一个相应的工厂类。
- 典型的解耦框架。高层模块只需要知道产品的抽象类，无须关心其他实现类，满足迪米特法则、依赖倒置原则和里氏替换原则。

工厂方法模式的缺点

- 类的个数容易过多，增加复杂度。
- 增加了系统的抽象性和理解难度。
- 只能生产一种产品，此弊端可使用抽象工厂模式解决。

无论是简单工厂还是工厂方法都只能生产一种产品，如果工厂需要创建生态里的多个产品，就需要更进一步，使用第三级的工厂模式--抽象工厂。



## 3. 抽象工厂

抽象工厂模式：用于创建一系列相关的或者相互依赖的对象。

为了更清晰地理解工厂方法模式和抽象工厂模式的区别，我们举一个品牌产品生态的例子。

比如智能家居领域多家公司，现在有华为和小米，他们的工厂除了生产我们熟知的手机外，还会生产电视、空调这种家电设备。

假如我们有幸作为他们工厂智能化管理软件的供应商，在软件系统里要对工厂进行抽象，这个时候就不能再用工厂方法这种设计模式了。因为工厂方法只能用来生产一种产品。

我们先看一下使用类图表示的这个抽象工厂抽象多品牌--多产品的形态。

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20221220101306350.png)

下面我们用代码简单实现一个抽象工厂，这个工厂能生成智能电视和空调，当然产品的功能在代码里比较简单，就是输出一条相应的信息。

我们先抽象一个工厂，它需要有生产视频和空调的能力：

```go
// AbstractFactory 抽象工厂类
// 目前有两个实际工厂类一个是华为的工厂，一个是小米的工厂
// 他们用来实际生产自家的产品设备
type AbstractFactory interface {
	CreateTelevision() ITelevision
	CreateAirConditioner() IAirConditioner
}

type ITelevision interface {
	Watch()
}

type IAirConditioner interface {
	SetTemperature(int)
}
```

现在有华为和小米两家厂商，他们有各自的产品：

```go

// HuaWeiTV 华为电视
type HuaWeiTV struct{}

func (ht *HuaWeiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

// HuaWeiAirConditioner 华为空调
type HuaWeiAirConditioner struct{}

func (ha *HuaWeiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("HuaWei AirConditioner set temperature to %d ℃\n", temp)
}

// MiTV 小米电视
type MiTV struct{}

func (mt *MiTV) Watch() {
	fmt.Println("Watch HuaWei TV")
}

// MiAirConditioner 小米空调
type MiAirConditioner struct{}

func (ma *MiAirConditioner) SetTemperature(temp int) {
	fmt.Printf("Mi AirConditioner set temperature to %d ℃\n", temp)
}
```

为了生产这些产品，它们需要创建各自的工厂，这些工厂都需要去实现抽象工厂：

```go

// HuaweiFactory 华为工厂，实现 AbstractFactory
type HuaWeiFactory struct{}

func (hf *HuaWeiFactory) CreateTelevision() ITelevision {
	return &HuaWeiTV{}
}
func (hf *HuaWeiFactory) CreateAirConditioner() IAirConditioner {
	return &HuaWeiAirConditioner{}
}

// MiFactory 小米工厂，实现 AbstractFactory
type MiFactory struct{}

func (mf *MiFactory) CreateTelevision() ITelevision {
	return &MiTV{}
}
func (mf *MiFactory) CreateAirConditioner() IAirConditioner {
	return &MiAirConditioner{}
}
```

这样抽象工厂就实现了，测试如下：

```go
func main() {
	var factory AbstractFactory
	var tv ITelevision
	var air IAirConditioner

	factory = &HuaWeiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(25)

	factory = &MiFactory{}
	tv = factory.CreateTelevision()
	air = factory.CreateAirConditioner()
	tv.Watch()
	air.SetTemperature(26)
}
```

同样抽象工厂也具备工厂方法把产品的创建推迟到工厂子类去做的特性，假如未来加入了 VIVO 的产品，我们就可以通过再创建 VIVO 工厂子类来扩展。

对于抽象工厂我们可以总结以下几点：

- 当系统所提供的工厂所需生产的具体产品并不是一个简单的对象，而是多个位于不同产品等级结构中属于不同类型的具体产品时需要使用抽象工厂模式。
- 抽象工厂模式是所有形式的工厂模式中最为抽象和最具一般性的一种形态。
- 抽象工厂模式与工厂方法模式最大的区别在于，工厂方法模式针对的是一个产品等级结构，而抽象工厂模式则需要面对多个产品等级结构，一个工厂等级结构可以负责多个不同产品等级结构中的产品对象的创建 。
- 当一个工厂等级结构可以创建出分属于不同产品等级结构的一个产品族中的所有对象时，抽象工厂模式比工厂方法模式更为简单、有效率。

抽象工厂模式的优点

- 当需要产品族时，抽象工厂可以保证客户端始终只使用同一个产品的产品族。
- 抽象工厂增强了程序的可扩展性，对于新产品族的增加，只需实现一个新的具体工厂即可，不需要对已有代码进行修改，符合开闭原则。

抽象工厂模式的缺点

- 规定了所有可能被创建的产品集合，产品族中扩展新的产品困难，需要修改抽象工厂的接口。
- 增加了系统的抽象性和理解难度。


## 对比

- 简单工厂模式：工厂类是整个模式的关键所在，包含了必要的逻辑判断，能够外界给定的信息， 决定究竟创建哪个具体类的对象。
- 工厂方法模式：是对简单工厂方法模式的一个抽象，抽离出了一个 Factory 接口，这个接口不负责具体产品的生产，而只是指定一些规范，具体的生产工作由其子类去完成。这个模式中，工厂类和产品类往往是一一对应的，完全解决了简单工厂模式中违背“开闭原则”的问题，实现了可扩展。
- 抽象工厂模式：的特点是存在多个抽象产品类，每个抽象产品类可以派生出多个具体产品类，工厂提供多种方法，去生产"系列"产品。

简单工厂模式适用于工厂类需要创建的对象比较少的情况，客户只需要传入具体的参数，就可以忽略工厂的生产细节，去获取想要的对象；
工厂方法模式，主要是针对单一产品结构的情景；
抽象工厂模式则是针对多级产品结构（系列产品）的一种工厂模式。



## 参考

- [工厂模式有三个Level，你能用Go写到第几层？](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247495992&idx=1&sn=591faf1acfbd5f5aa7f0dbc95506f85c&chksm=fa8320afcdf4a9b965a768e34dff675e754de7e0c9ad95ae3134be1fa4b0990c93c6e8b3162c&scene=178&cur_album_id=2531498848431669249#rd)
