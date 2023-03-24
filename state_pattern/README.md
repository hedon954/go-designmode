# 状态模式

## 1. 概念

状态模式是一种行为型设计模式，它允许一个对象在其内部状态改变时改变它的行为。该模式将一个对象的状态分离成不同的类，使得每个状态的行为都可以独立地进行修改和扩展，同时也降低了对象的复杂性。

在状态模式中，定义了一个抽象状态类和一系列具体状态类，每个具体状态类代表对象的一种状态，并定义了在该状态下对象的行为。对象通过维护当前状态来决定它的行为。在运行时，可以动态地改变对象的状态，从而改变它的行为。

状态模式主要由三个角色构成。

- Context（环境类）：环境类又称为上下文类，它定义客户端需要的接口，内部维护一个当前状态实例，并负责具体状态的切换。
- State（抽象状态）：定义状态下的行为，可以有一个或多个行为。
- ConcreteState（具体状态）：每一个具体状态类对应环境的一个具体状态，不同的具体状态类其行为有所不同。

状态模式的优点包括：

1. 将状态转移逻辑封装在具体状态类中，使得状态的转移更加明确、易于理解和修改。
2. 在增加新的状态类时，不需要修改原有代码，只需要增加新的状态类即可。
3. 使得对象状态的转移更加有条理，易于维护和扩展。

状态模式的缺点包括：

1. 增加了系统中的类和对象数量，增加了系统的复杂性。
2. 在状态较多时，可能会导致类的数量剧增，使得代码难以维护。

状态模式适用于以下情况：

1. 当一个对象的行为取决于其内部状态时。
2. 当一个对象的行为在运行时需要根据其状态改变时。
3. 当有多个状态且每个状态下的行为不同，且这些状态的转移比较复杂时。

## 2. 理解

状态模式就是将一个对象的每个状态单独处理，然后通过一个统一的入口去设置和调用特定状态的具体行为。

## 3. 类图

![image-20230324114738756](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/image-20230324114738756.png)

## 4. 实现

假设有一个游戏中的角色对象，它有三种状态：正常状态、受伤状态和死亡状态，每种状态下对象的行为不同，具体实现如下：

```go
type State interface {
	Update(role *Role)
}

type Role struct {
	state State
}

// SetState sets the role's state
func (r *Role) SetState(state State) {
	r.state = state
}

// Update updates the role's state
func (r *Role) Update() {
	r.state.Update(r)
}

type NormalState struct{}

func (s *NormalState) Update(role *Role) {
	fmt.Println("role is in normal state")
}

// 受伤状态
type InjuredState struct{}

func (s *InjuredState) Update(role *Role) {
	// 受伤状态下的行为
	fmt.Println("role is in injured state")
}

type DeadState struct{}

func (s *DeadState) Update(role *Role) {
	// 死亡状态下的行为
	fmt.Println("role is in dead state")
}

func main() {
	role := &Role{}
	normalState := &NormalState{}
	injuredState := &InjuredState{}
	deadState := &DeadState{}

	role.SetState(normalState)
	role.Update()

	role.SetState(injuredState)
	role.Update()

	role.SetState(deadState)
	role.Update()
}
```

在上面的例子中，我们定义了一个角色对象和三种状态：正常状态、受伤状态和死亡状态。角色对象通过 SetState 方法设置当前状态，通过 Update 方法执行当前状态下的行为。在实际的开发中，可以根据具体需求对状态和状态行为进行扩展和修改。

## 5. 场景

状态模式是一种非常常见的设计模式，它适用于以下场景：

1. 对象的行为取决于其状态，并且对象的状态可能在运行时发生变化。
2. 在对象的状态发生变化时，对象的行为也会随之发生变化。
3. 对象状态的转换比较复杂，或者包含大量的条件分支语句。
4. 为了避免大量的条件分支语句，导致代码难以维护和扩展，需要将状态转换的代码分离出来。

一些具体的应用场景包括：

1. 订单状态：订单在不同的状态下，其对应的行为也是不同的，例如：待支付状态、已支付状态、已发货状态等。
2. 网络连接状态：网络连接在不同的状态下，其对应的行为也是不同的，例如：未连接状态、连接中状态、已连接状态、已断开状态等。
3. 交通信号灯状态：交通信号灯在不同的状态下，其对应的行为也是不同的，例如：红灯状态、黄灯状态、绿灯状态等。
4. 游戏对象状态：游戏对象在不同的状态下，其对应的行为也是不同的，例如：正常状态、受伤状态、死亡状态等。

## 参考

- ChatGPT
- [Go设计模式实战--用状态模式实现系统工作流和状态机](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497490&idx=1&sn=cc4cab583d6691e27434ddab90da6f12&chksm=fa832685cdf4af93403876990caecbcdeff032f685a5ce93d1eaf82b0f2f30c439e817b0ede3&scene=178&cur_album_id=2531498848431669249#rd)