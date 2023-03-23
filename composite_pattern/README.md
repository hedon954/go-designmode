# 组合模式

## 1. 概念

组合模式（Composite Pattern）又叫作部分-整体（Part-Whole）模式，它的宗旨是通过将单个对象（叶子节点）和组合对象（树枝节点）用相同的接口进行表示，使得客户对单个对象和组合对象的使用具有一致性，属于结构型设计模式。

组合模式由以下几个角色构成：

- 组件 （Component）： 组件是一个接口，描述了树中单个对象和组合对象都要实现的的操作。
- 叶节点 （Leaf） ：即单个对象节点，是树的基本结构， 它不包含子节点，因此也就无法将工作指派给下去，叶节点最终会完成大部分的实际工作。
- 组合对象 （Composite）：是包含叶节点或其他组合对象等子项目的符合对象。 组合对象不知道其子项目所属的具体类， 它只通过通用的组件接口与其子项目交互。
- 客户端 （Client）： 通过组件接口与所有项目交互。 因此， 客户端能以相同方式与树状结构中的简单或复杂对象进行交互。

## 2. 理解

组合模式的使用要求业务场景中的实体必须能够表示成树形结构才行，由组合模式将一组对象组织成树形结构，客户端（代码的使用者）可以将单个对象和组合对象都看做树中的节点，以统一处理逻辑，并且利用树形结构的特点，将对树、子树的处理转化成叶节点的递归处理，依次简化代码实现。

通过上边的描述我们可以马上想到文件系统、公司组织架构这些有层级结构的事物的操作会更适合应用组合模式。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/092145017171491.png)

## 4. 实现

在游戏服务器开发中，组合模式可以用于处理玩家的行为和动作。例如，我们可以创建一个名为 `Action` 的接口，它定义了执行操作所需的方法。然后，我们可以创建一个名为 `PlayerAction` 的结构体，它实现了 `Action` 接口，并具有一些属性，例如玩家ID、行动时间戳等。我们还可以创建一个名为 `CompositeAction` 的结构体，它实现了 `Action` 接口，并包含多个 `PlayerAction` 实例。

### 4.1 定义组件

```go
// Action is the composite component
type Action interface {

	// Execute does the action
	Execute()
}
```

### 4.2 实现叶子结点

```go
// PlayerAction is the composite leaf, it implements the Action interface
type PlayerAction struct {
	playerId   int
	actionTime time.Time
}

func (p *PlayerAction) Execute() {
	fmt.Printf("Player %d performed actions at %v\n", p.playerId, p.actionTime)
}
```

### 4.3 实现组合

```go
// CompositeAction is the composite, it is the collection of Action
type CompositeAction struct {
	actions []Action
}

func (c *CompositeAction) Execute() {
	// run the composited actions
	for _, action := range c.actions {
		action.Execute()
	}
}

// AddAction adds an action to the actions collection
func (c *CompositeAction) AddAction(action Action) {
	c.actions = append(c.actions, action)
}
```

### 4.4 使用

```go
func main() {
  
  // the first sub composite
	action1 := &CompositeAction{}
	action11 := &PlayerAction{
		playerId:   11,
		actionTime: time.Now(),
	}
	action12 := &PlayerAction{
		playerId:   12,
		actionTime: time.Now(),
	}
	action1.AddAction(action11)
	action1.AddAction(action12)

  // the second sub composite
	action2 := &CompositeAction{}
	action21 := &PlayerAction{
		playerId:   21,
		actionTime: time.Now(),
	}
	action22 := &PlayerAction{
		playerId:   22,
		actionTime: time.Now(),
	}
	action2.AddAction(action21)
	action2.AddAction(action22)

  // the final composite
	composite := &CompositeAction{}
	composite.AddAction(action1)
	composite.AddAction(action2)

	composite.Execute()
}
```

## 5. 场景

1. 等等。在这些情况下，每个节点都可以是一个叶子节点，也可以是一个父节点，它包含了多个子节点，从而形成了一种树形结构。
2. 处理递归结构：组合模式可以用于处理递归结构，例如表达式树、XML 文档、语法分析树等等。在这些情况下，每个节点都可以是一个叶子节点，也可以是一个父节点，它包含了多个子节点，从而形成了一种递归结构。
3. 处理对象集合：组合模式可以用于处理对象集合，例如图形界面中的控件集合、游戏中的场景对象集合等等。在这些情况下，每个节点都代表一个对象，它们可以按照一定的方式进行组合，从而形成一个对象集合。

总的来说，组合模式适用于一些具有层次结构的场景，它可以将复杂的结构分解成一个个简单的组件，从而方便对整个结构进行操作和管理。

## 6. 对比

### 6.1 组合 vs 装饰器

- **组合模式**：为叶子对象和组合对象提供了统一的接口，叶子对象分担组合对象要做的工作。其实组合对象就是派了下活儿，等下面的干完后，它再给上层调用者返（汇）回（报），类似于公司里的那些组合。
- **装饰器模式**：装饰器属于大哥带小弟的类型，核心的活儿是小弟干的（小弟就是被装饰的对象）但是各位大哥会帮你做好干活儿之外的事儿，比如公司你在公司里的 Mentor、项目经理、领导们干的事儿就是给在给你做增强，你可以把他们理解成是你的装饰器。





## 参考

- ChatGPT
- [Go 设计模式｜组合，一个对数据结构算法和职场都有提升的设计模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497331&idx=1&sn=cb8a154b7e6b2913cbd4db89dfe27e80&chksm=fa8327e4cdf4aef21935d4c524f26975bcd38e5e0cfb4145c751e5cca6b7e36c58306d647923&scene=178&cur_album_id=2531498848431669249#rd)