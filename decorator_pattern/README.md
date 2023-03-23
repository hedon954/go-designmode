# 装饰器模式

给对象添加新行为最简单直观的办法就是扩展本体对象，通过继承的方式达到目的。但是使用继承不可避免地有如下两个弊端：

1. 继承是静态的，在编译期间就已经确定，无法在运行时改变对象的行为。
2. 子类只能有一个父类，当需要添加的新功能太多时，容易导致类的数量剧增。

而使用装饰器模式，我们通过将现有对象放置在实现了相同一套接口的包装器对象中来动态地向现有对象添加新行为。在包装器中进行我们代码的扩展，有助于重用功能并且不会修改现有对象的代码，符合“开闭原则”。

## 1. 概念

装饰器模式（Decorator Pattern）也叫作包装器模式（Wrapper Pattern），是一种结构型设计模式，它允许在运行时动态地将行为添加到一个对象中，而无需修改其源代码。它通过包装一个已有的类，来增加类的功能和行为，从而扩展了该类的功能。

装饰器模式的核心思想是将对象的功能分离，使其单独修改，而不需要改变对象的结构。该模式由三个主要组件组成：

- 抽象组件（Component）：定义了一个抽象接口，以规定具体组件和装饰器的行为。
- 具体组件（Concrete Component）：实现抽象组件接口的类，它定义了基本的行为。
- 装饰器（Decorator）：实现了抽象组件接口，并在其中包含一个指向具体组件的引用，它增加了具体组件的功能。

在装饰器模式中，装饰器类和具体组件类共享一个相同的接口，这使得它们可以互相替换。装饰器可以用于增加、删除或修改具体组件的行为，而不会影响到其他的组件。因此，装饰器模式通常用于动态地扩展对象的功能。



## 2. 理解

装饰器模式就是代理模式的一种特殊应用，实际上就是用对象 B 去包含 对象 A，在 B 调用 A 的基础上，B 可以做一些增强。但是由于 A 和 B 实现了同样的接口，所以可以做到外界无感知。



## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/cb52ef587f1ca6542e54905cb00537b9.png)



## 4. 实现

### 4.1 定义抽象组件

```go
// Player is the Component
type Player interface {
	ChooseClass(class string)
	CastSkill() string
}
```

### 4.2 实现具体组件

```go
// BasicPlayer is the concrete who implements the Player interface
type BasicPlayer struct {
	Class string
}

func (c *BasicPlayer) ChooseClass(class string) {
	c.Class = class
}

func (c *BasicPlayer) CastSkill() string {
	return "basic class"
}
```

### 4.3 实现装饰器

```go
// FirePlayer is the decorator of Player
type FirePlayer struct {
   player Player
}

func (f *FirePlayer) ChooseClass(class string) {
   f.player.ChooseClass(class)
}

func (f *FirePlayer) CastSkill() string {
   return f.player.CastSkill() + ", fireball"
}
```

### 4.4 使用

```go
func main() {
	player := &BasicPlayer{}

	// firePlayer decorates the basicPlayer
	firePlayer := &FirePlayer{player: player}
	firePlayer.ChooseClass("mage")
	skill := firePlayer.CastSkill()
	fmt.Println(skill)
}
```

### 4.5 更进一步

当然，你可以更进一步进行多层装饰器，也可以随意对任意 Player 进行装饰顺序的组合。



## 5. 对比

### 5.1 装饰器 vs 代理模式

- 装饰器模式就是代理模式的一个特殊应用。
- 装饰器模式强调自身功能的扩展。
- 代理模式强调对代理过程的控制。

### 5.2 装饰器 vs 职责链模式

装饰器和职责链在行为上看都是多个单元进行组合完成逻辑处理，但是装饰器注重给某样东西添加扩展，最终会得到一个产品。而职责链更强调分步骤完成某个流程，更像是一个任务链表，而且与装饰器模式不同的是，职责链可以随时终止。



## 参考

- ChatGPT
- [Go学设计模式--装饰器和职责链，哪个模式实现中间件更科学](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497282&idx=1&sn=0de76856e8649967bd3979cb122383fd&chksm=fa8327d5cdf4aec375fcae915f4eba2960b8f07f91b445f137319c81ea5a9ac4349f44d204d4&scene=178&cur_album_id=2531498848431669249#rd)

