# 迭代器模式

## 1. 概念

迭代器模式（iterator pattern）是一种行为设计模式，它提供了一种访问集合对象的方法，而不需要暴露集合对象的内部实现。迭代器模式可以将迭代操作与集合对象本身分离开来，使得代码更加简单和易于维护。

在迭代器模式中，我们通常会定义两个核心组件：

1. 迭代器（Iterator）：负责定义访问和遍历元素的接口，并跟踪集合的当前位置。迭代器提供了一些常用的方法，例如 `hasNext`（判断是否还有下一个元素）和 `next`（返回下一个元素）。
2. 集合（Collection）：负责提供创建迭代器的接口，以便客户端能够访问集合中的元素。集合通常是一个抽象类或接口，它定义了一些与集合对象相关的操作，例如 `add`（添加元素）和 `remove`（删除元素）。

在实现迭代器模式时，我们通常需要考虑以下几个方面：

1. 迭代器如何知道何时停止遍历集合对象？我们可以在迭代器中实现一个 `hasNext` 方法，用来判断是否还有下一个元素。当遍历到最后一个元素时，该方法会返回 `false`，从而停止遍历。
2. 如何访问集合对象中的元素？我们可以在迭代器中实现一个 `next` 方法，用来返回当前位置的元素。每次调用 `next` 方法时，迭代器会将指针移动到下一个位置。
3. 如何保证迭代器的安全性？迭代器的设计应该尽可能地简单和安全。为了防止迭代器的越界访问，我们可以在迭代器中实现一些边界检查，例如在 `hasNext` 方法中检查当前位置是否小于集合大小。此外，我们还可以实现一些保护性的方法，例如 `remove` 方法，用来在迭代过程中安全地删除元素。

Java 语言的 Collection、Map 类族中提供的各种迭代器是对该模式的典型应用，代码实现写的很优秀，值得借鉴学习。

## 2. 理解

迭代器模式通过提供一个统一的遍历集合的方式，让使用者无需关系集合的底层结构，还可以通过统一的方式来访问不同结构的集合。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/1946490-20200226092609706-1444971982.jpg)

## 4. 实现

假设我们有一个商品列表，我们需要实现一个功能，可以遍历该列表并输出每个商品的信息。这时候，我们可以使用迭代器模式来实现。

### 4.1 定义迭代器接口

```go
type Iterator interface {
	HasNext() bool
	Next() interface{}
}
```

### 4.2 定义商品列表

```go
type Product struct {
	Name  string
	Price float64
}

type ProductList struct {
	Products []*Product
}
```

### 4.3 实现商品迭代器

```go
// ProductIterator is the iterator for product list
type ProductIterator struct {
	Index    int
	Products []*Product
}

func (p *ProductIterator) HasNext() bool {
	return p.Index < len(p.Products)
}

func (p *ProductIterator) Next() interface{} {
	defer func() {
		p.Index++
	}()
	return p.Products[p.Index]
}
```

### 4.4 实现创建迭代器方法

在 `ProductList` 结构体中，我们提供了一个 `CreateIterator()` 方法来创建一个 `ProductIterator` 实例，用来遍历商品列表。

```go
// CreateIterator returns an iterator for ProductList
func (pl *ProductList) CreateIterator() Iterator {
	return &ProductIterator{
		Index:    0,
		Products: pl.Products,
	}
}
```

### 4.5 使用

```go
func main() {
	productList := ProductList{
		Products: []*Product{
			{Name: "iPhone", Price: 9.9},
			{Name: "iPad", Price: 8.8},
			{Name: "Mac", Price: 6.6},
		},
	}
	iterator := productList.CreateIterator()
	for iterator.HasNext() {
		product := iterator.Next().(*Product)
		fmt.Println(product)
	}
}
```

## 5. 场景

1. 遍历集合：迭代器模式最常见的应用场景就是遍历集合。在很多编程语言中，像数组、列表、字典等集合都提供了迭代器接口，可以通过迭代器来遍历集合中的元素。
2. 封装复杂数据结构的遍历算法：如果一个数据结构很复杂，遍历它的算法也很复杂，这时可以使用迭代器模式来将遍历算法封装起来，使得使用者可以简单地遍历数据结构。
3. 无需暴露集合的内部结构：通过使用迭代器模式，可以隐藏集合的内部结构，从而避免使用者直接访问集合的内部结构。这样可以提高集合的安全性和可维护性。
4. 实现延迟计算：有些数据结构在迭代时需要进行复杂的计算，这时可以使用迭代器模式实现延迟计算，只有当需要遍历到某个元素时才进行计算，从而提高程序的性能。
5. 支持多种遍历方式：有些数据结构支持多种遍历方式，比如前序遍历、后序遍历、层次遍历等，这时可以使用迭代器模式来实现不同的迭代器，从而支持多种遍历方式。

## 参考

- ChatGPT
- [Go设计模式-迭代器到底是不是一个多此一举的模式？](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497742&idx=1&sn=12d38e4d727962233dad0833bb0966b7&chksm=fa832999cdf4a08fac7ef3d74399e18fd28e078d009c8abd98f76b5a77de433da9042cb31457&scene=178&cur_album_id=2531498848431669249#rd)