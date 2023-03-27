# 享元模式

## 1. 概念

享元模式（flyweight pattern）是一种可以优化性能和节省内存的设计模式，它适用于需要频繁创建许多相似对象的场景。在享元模式中，通过共享内部状态，将多个对象合并为一个共享对象，以避免创建多个重复对象，从而减少内存占用。同时，通过在对象内部维护外部状态，保证对象在不同的上下文环境下能够正确地展现自己。

享元模式的核心是将对象分为内部状态和外部状态两部分。内部状态是不变的，相同的内部状态对应的对象可以被共享使用，而外部状态是变化的，需要作为参数传递给对象。通过这种方式，我们可以共享对象的内部状态，避免创建过多的对象，从而优化程序的性能和内存占用。

享元模式一般由下面几个组成：

1. 抽象享元角色（Flyweight）：定义了享元对象的接口和公共方法，用于管理外部状态。
2. 具体享元角色（ConcreteFlyweight）：实现抽象享元角色的接口，维护内部状态，并对外部状态进行处理。
3. 享元工厂角色（FlyweightFactory）：管理并创建享元对象，并根据内部状态进行缓存和共享。
4. 客户端角色（Client）：使用享元对象的角色，在使用时需要传递外部状态。

优点：

1. 减少内存占用：通过共享内部状态，避免创建重复对象，减少内存占用。
2. 提高程序性能：减少对象创建和销毁的次数，提高程序性能。
3. 支持大量细粒度对象的重用：在需要创建大量细粒度对象的场景中，可以显著减少对象的创建数量，从而提高程序的性能。

缺点：

1. 实现复杂度高：需要将对象分为内部状态和外部状态两部分，并对外部状态进行处理，实现复杂度较高。
2. 适用场景受限：仅适用于需要创建大量相似对象的场景，如果对象类型不同或者对象数量较少，则不适用。

## 2. 理解

享元模式其实就类似于单例模式，只不过它通过定义接口，然后通过享元工厂类来维护实现该接口的每一个实例对象，不管是对接口的不同实现，还是对接口同一实现但拥有不同的内部数值，都可以只维护一份。这样就达到优化程序性能和减少内存占用的效果。故为“享元”。

前提是内部状态是不可变的。

> 更进一步，我们还可以记录每个对象的最近使用时间，再起一个协程定时去探测对象的最近使用时间，如果对象太久没使用，我们就可以释放它，进一步减少内存占用。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/u=3638920120,2039450650&fm=253&fmt=auto&app=138&f=JPEG-20230327110844702.jpeg)

## 4. 实现

假设我们需要实现一个多人游戏服务器，玩家在游戏中可以选择不同的角色进行游戏。每个角色都有不同的属性和技能，但是有些属性和技能是多个角色共享的，例如移动速度、攻击力等。如果直接创建大量的角色对象，将会占用大量的内存，导致程序性能下降。

因此，我们可以使用享元模式来管理角色对象。具体实现如下：

### 4.1 抽象享元角色

首先，我们定义一个抽象的角色接口 Role，包含了一些公共的属性和方法：

```go
type Role interface {
    GetName() string
    GetSpeed() int
    GetPower() int
    SetSpeed(speed int)
    SetPower(power int)
    UseSkill()
}
```

### 4.2 具体享元角色

```go
// RoleImpl is the concrete flyweight instance
type RoleImpl struct {
	name  string // external status
	speed int    // internal status
	power int    // internal status
}

func (r *RoleImpl) GetName() string {
	return r.name
}

func (r *RoleImpl) GetSpeed() int {
	return r.speed
}

func (r *RoleImpl) GetPower() int {
	return r.power
}

func (r *RoleImpl) SetSpeed(speed int) {
	r.speed = speed
}

func (r *RoleImpl) SetPower(power int) {
	r.power = power
}

func (r *RoleImpl) UseSkill() {
	fmt.Printf("%s uses skill\n", r.name)
}
```

### 4.3 享元工厂角色

```go
// RoleFactory is the flyweight factory, it holds all the flyweight objects in memory
type RoleFactory struct {
	roles sync.Map
}

func (rf *RoleFactory) GetRole(name string) Role {

	// if exists, return it
	if role, ok := rf.roles.Load(name); ok {
		return role.(Role)
	}

	// first time to invoke, create it
	var newRole Role
	switch name {
	case "warrior":
		newRole = &RoleImpl{
			name:  "warrior",
			speed: 5,
			power: 10,
		}
	case "mage":
		newRole = &RoleImpl{
			name:  "mage",
			speed: 3,
			power: 15,
		}
	case "archer":
		newRole = &RoleImpl{
			name:  "archer",
			speed: 8,
			power: 8,
		}
	}

	// use loadOrStore to determine if any other goroutine has been created the role
	actual, loaded := rf.roles.LoadOrStore(name, newRole)
	if loaded {
		return actual.(Role)
	}
	return newRole
}
```

## 5. 场景

享元模式可以应用于需要创建大量对象的场景，其中许多对象的状态是相同或类似的，这些状态可以被共享来减少内存占用和提高性能。

具体来说，以下是一些常见的应用场景：

1. 文字处理器：文字处理器可以使用享元模式来缓存已经创建的字体、字号等对象。这些对象可以在文档中多次使用，所以缓存它们可以减少内存的占用和创建对象的时间。
2. 图像处理器：图像处理器可以使用享元模式来缓存已经创建的图像对象。这些对象可以在多个地方重复使用，所以缓存它们可以减少内存的占用和创建对象的时间。
3. 游戏开发：在游戏中，有很多元素需要创建多次，如游戏角色、敌人、子弹、道具、音效等。在这种情况下，可以使用享元模式来共享这些元素的相同或相似的属性，如纹理、大小、速度、颜色等，以减少内存的占用和提高性能。
4. 网络编程：在网络编程中，有时需要创建大量的连接对象，这些对象的状态是相似的，如IP地址、端口号等。可以使用享元模式来共享这些相同的状态，以减少内存的占用和提高性能。
5. 数据库访问：在数据库访问中，有很多查询语句需要创建多次。可以使用享元模式来共享已经创建的查询语句对象，以减少内存的占用和提高性能。

总之，享元模式可以应用于需要创建大量对象的场景，并且这些对象的状态是相同或类似的，这些状态可以被共享来减少内存占用和提高性能。

## 参考

- ChatGPT
- [Go设计模式--享元模式，节省内存的好帮手](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497613&idx=1&sn=428a4ffa977421b2b2c78f36585b7c62&chksm=fa83261acdf4af0c014190c5982b1b5af34018d69e12c3baf5a1f44b680056901103d2afd2c0&scene=178&cur_album_id=2531498848431669249#rd)