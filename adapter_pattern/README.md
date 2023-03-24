# 适配器模式

## 1. 概念

适配器模式（Adapter Pattern）是一种结构型设计模式，它能够将一个类的接口转换成另一个类的接口，使得原本由于接口不兼容而无法协作的类能够在一起工作。

在实际应用中，我们常常会遇到需要使用现有的类，但是这些类的接口与我们所需的接口不一致的情况。适配器模式可以解决这一问题，使得原本不兼容的类可以协同工作。

适配器模式一般包含以下几个角色：

- 目标接口（Target Interface）：客户端所需要的接口，适配器需要将原始类的接口转换成目标接口。
- 原始类（Adaptee）：需要被适配的类，其接口与目标接口不兼容。
- 适配器（Adapter）：实现目标接口，包含一个原始类的实例，通过调用原始类的方法，将其接口转换成目标接口。

适配器模式一般分为两种类型：类适配器和对象适配器。

1. 类适配器

   在类适配器中，适配器继承自原始的类，并实现客户端所需要的接口。通过这种方式，适配器可以将客户端所需要的接口转换成原始类的接口。

   ```java
   class Adaptee {
       public void specificRequest() {
           // 具体的功能实现
       }
   }
   
   interface Target {
       void request();
   }
   
   class Adapter extends Adaptee implements Target {
       public void request() {
           specificRequest();
       }
   }
   ```

   在上面的代码中，Adaptee 是原始的类，它提供了一个不兼容的接口 specificRequest。Adapter 继承自 Adaptee，并实现了客户端所需要的接口 Target。在 Adapter 的 request 方法中，调用了 Adaptee 的 specificRequest 方法，从而完成了接口的转换。

2. 对象适配器

   在对象适配器中，适配器包含一个原始的类的实例，并实现客户端所需要的接口。通过这种方式，适配器可以将客户端所需要的接口转换成原始类的接口。

   ```java
   class Adaptee {
       public void specificRequest() {
           // 具体的功能实现
       }
   }
   
   interface Target {
       void request();
   }
   
   class Adapter implements Target {
       private Adaptee adaptee;
   
       public Adapter(Adaptee adaptee) {
           this.adaptee = adaptee;
       }
   
       public void request() {
           adaptee.specificRequest();
       }
   }
   ```

   在上面的代码中，Adaptee 是原始的类，它提供了一个不兼容的接口 specificRequest。Adapter 包含一个 Adaptee 的实例，并实现了客户端所需要的接口 Target。在 Adapter 的 request 方法中，调用了 Adaptee 的 specificRequest。

## 2. 理解

适配器就是中间加一层转换，将不同的接口连接在一起，类似于日常生活中使用的转接口。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/011620181723824.png)

## 4. 实现

假设我们正在开发一个游戏服务器，其中有一个场景管理器（Scene Manager）负责管理游戏中所有的场景，而每个场景都有一个不同的类型，如普通场景（Normal Scene）、Boss 场景（Boss Scene）、PVP 场景（PVP Scene）等等。现在我们需要在不同类型的场景之间进行切换，但是每种场景都有自己不同的接口和实现，这就需要使用适配器模式来进行适配。

下面是一个简单的示例代码，演示了如何使用适配器模式来实现场景的切换：

```go
// 场景接口
type Scene interface {
    Enter()
}

// 普通场景
type NormalScene struct {}

func (ns *NormalScene) Enter() {
    fmt.Println("Enter Normal Scene")
}

// Boss场景
type BossScene struct {}

func (bs *BossScene) Start() {
    fmt.Println("Enter Boss Scene")
}

// Boss场景适配器
type BossSceneAdapter struct {
    BossScene *BossScene
}

func (bsa *BossSceneAdapter) Enter() {
    bsa.BossScene.Start()
}

// 场景管理器
type SceneManager struct {
    CurrentScene Scene
}

func (sm *SceneManager) ChangeScene(s Scene) {
    sm.CurrentScene = s
    sm.CurrentScene.Enter()
}

// 使用示例
func main() {
    sm := &SceneManager{}
    ns := &NormalScene{}
    sm.ChangeScene(ns)

    bs := &BossScene{}
    bsa := &BossSceneAdapter{BossScene: bs}
    sm.ChangeScene(bsa)
}
```

在上面的示例代码中，我们定义了Scene接口来表示场景，NormalScene和BossScene分别实现了Scene接口。然后，我们定义了一个BossSceneAdapter来适配BossScene，使其也能够实现Scene接口。最后，我们在SceneManager中使用Scene接口来管理场景，并使用适配器将BossScene适配成Scene接口。

通过使用适配器模式，我们可以将不同类型的场景适配成统一的接口，从而实现了场景之间的切换。这样，即使我们添加了新的场景类型，也可以很容易地将其适配成Scene接口，而无需修改SceneManager的代码。

## 5. 场景

适配器模式是一种很常见的设计模式，它的应用场景非常广泛，下面列举几个典型的应用场景：

1. 对现有类库或第三方类库进行封装

   在开发过程中，我们可能会使用一些现有的类库或第三方类库，但是它们的接口与我们的代码不兼容，这时候就可以使用适配器模式来对这些类进行封装，使得它们可以被我们的代码所使用。

2. 在系统升级中保持兼容性

   当我们对系统进行升级时，很可能会出现接口的变化，而某些旧的系统可能无法适应新接口的变化，这时候就可以使用适配器模式来保持系统的兼容性。

3. 对不兼容的数据进行转换

   在实际应用中，数据的格式往往是多种多样的，有时候我们需要将一种数据格式转换成另一种数据格式，这时候就可以使用适配器模式来进行数据转换。

4. 对不同的类进行适配

   在实际应用中，我们可能会遇到需要将不同的类适配成一个类的情况，这时候就可以使用适配器模式来将这些不同的类适配成一个类，以便于进行统一的处理。

5. 在框架设计中使用

   在框架设计中，适配器模式也是一种非常常见的设计模式，它可以将框架中的不同组件适配成一个统一的接口，以方便开发者进行开发。例如，在Spring框架中就有很多适配器，用来将不同的组件适配成统一的接口。



## 参考

- ChatGPT
- [Go 设计模式｜项目依赖耦合度太高？可以用适配器做下优化](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497405&idx=1&sn=425262da90e6812599960616480718d6&chksm=fa83272acdf4ae3c302b52af82166d5ab5262d3368d4190008057f5168a1f44d445bf482c83d&scene=178&cur_album_id=2531498848431669249#rd)