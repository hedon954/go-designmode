# 访问者模式

## 1. 概念

访问者模式（Visitor Pattern）是一种行为型设计模式，它允许你定义一些操作，而这些操作可以应用于不同类型的对象。这样你就可以在不改变这些对象的类的前提下，定义新的操作。访问者模式适用于那些数据结构相对稳定，但是需要在其上定义新的操作的场景。

在访问者模式中，通常有两类对象：元素（Element）和访问者（Visitor）。元素是被操作的对象，而访问者定义了要对这些元素执行的操作。元素会接受访问者作为参数，以便让访问者对自己进行操作。

以下是访问者模式的基本结构：

- 抽象访问者（AbstractVisitor）：定义了要对每个元素执行的操作。
- 具体访问者（ConcreteVisitor）：实现了抽象访问者定义的操作，以便于对元素进行处理。
- 抽象元素（AbstractElement）：定义了接受访问者的方法。
- 具体元素（ConcreteElement）：实现了抽象元素定义的接受访问者的方法，以便于访问者对自己进行操作。
- 对象结构（ObjectStructure）：维护了元素的集合，提供了可以接受访问者的方法。

## 2. 理解

访问者在 Element 数据结构比较稳定的时候，可以抽象出 AbstractElement 和 AbstractVisitor，这样就可以在不改变元素类结构的情况下，将算法封装到访问者中，并对元素对象进行遍历和访问，从而实现特定的操作。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/format,png.png)

## 4. 实现

以下是一个使用访问者模式的示例，该示例演示如何计算不同种类图形的面积。在该示例中，元素是不同种类的图形，而访问者是计算图形面积的操作。

### 4.1 定义抽象元素和访问者接口

```go
// Shape is the interface of all kinds of shape
type Shape interface {
	Accept(Visitor)
}

// Visitor is the interface of shape visitor, it provides methods to visit shapes
type Visitor interface {
	VisitForRectangle(*Rectangle)
	VisitForCircle(*Circle)
}
```

### 4.2 实现具体形状

```go
type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitForRectangle(r)
}

type Circle struct {
	radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitForCircle(c)
}
```

### 4.3 实现具体访问者

```go
// AreaVisitor is a visitor to visit shape's area
type AreaVisitor struct {
	area float64
}

func (a *AreaVisitor) VisitForRectangle(rectangle *Rectangle) {
	a.area += rectangle.width * rectangle.height
}

func (a *AreaVisitor) VisitForCircle(circle *Circle) {
	a.area += circle.radius * circle.radius * math.Pi
}
```

### 4.4 实现对象结构

```go
// ShapeCollection is a collection of shapes
type ShapeCollection struct {
	shapes []Shape
}

func (a *ShapeCollection) AddShapes(shapes ...Shape) {
	a.shapes = append(a.shapes, shapes...)
}

func (a *ShapeCollection) Accept(visitor Visitor) {
	for _, shape := range a.shapes {
		shape.Accept(visitor)
	}
}
```

### 4.5 使用

```go
func main() {
	// creates rectangles
	r1 := &Rectangle{width: 3, height: 4}
	r2 := &Rectangle{width: 4, height: 5}

	// creates circles
	c1 := &Circle{radius: 3}
	c2 := &Circle{radius: 4}

	// creates collection
	coll := &ShapeCollection{}
	coll.AddShapes(r1, r2, c1, c2)

	// creates area visitor
	visitor := &AreaVisitor{}

	// use
	coll.Accept(visitor)
	fmt.Println(visitor.area)
}
```

在这个例子中，我们创建了一个对象结构 ShapeCollection，其中包含了不同类型的图形元素（Rectangle 和 Circle）。我们也定义了一个具体访问者 AreaVisitor，用于计算图形的面积。最后，我们让对象结构接受访问者，并输出计算得到的面积。

这个例子展示了访问者模式的一个常见用例，即在一个稳定的对象结构中定义新的操作，同时避免修改元素类的情况。这种模式使得系统更加灵活，同时避免了类之间的紧耦合。

## 5. 场景

访问者模式适用于以下场景：

1. 对象结构比较稳定，但经常需要在此对象结构上定义新的操作。
2. 需要对一个对象结构中的对象进行很多不同的、且不相关的操作，而需要避免这些操作污染这些对象的类。
3. 需要通过对对象结构中的元素进行一些基本操作来组合出更复杂的操作。
4. 对象结构包含很多类型的对象，希望对这些对象实施一些依赖于其具体类型的操作。
5. 需要遍历一个复杂的对象结构，同时又不希望让这个遍历操作与对象结构本身耦合在一起。

常见的应用场景包括：

1. 编辑器：可以通过访问者模式实现对文档或者其他对象的多种编辑操作，如字体大小、颜色、加粗等。
2. 编译器：可以通过访问者模式实现语法树的遍历，对不同类型的语法节点进行不同的处理。
3. UI系统：可以通过访问者模式实现对UI元素的处理，如绘制、布局、事件处理等。
4. 数据库系统：可以通过访问者模式实现对数据库中的不同数据结构的处理，如表、字段、索引等。
5. 游戏开发：可以通过访问者模式实现游戏中的技能系统、AI系统、任务系统等。

## 参考

- ChatGPT