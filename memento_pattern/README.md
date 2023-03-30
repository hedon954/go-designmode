# 备忘录模式



## 1. 概念

备忘录模式是一种行为型设计模式，它允许在不暴露对象实现细节的情况下捕获和恢复其内部状态。它通过在不破坏封装性原则的前提下，让对象在不同时间保存自身的状态，并在需要时恢复状态。

备忘录模式通常包含三个角色：

- 发起者（Originator）：被保存和恢复状态的对象。
- 备忘录（Memento）：保存发起者状态的对象。
- 管理者（Caretaker）：负责协调发起者和备忘录之间的交互，以便发起者能够在不破坏封装性原则的前提下保存和恢复状态。

在备忘录模式中，发起者对象会创建备忘录对象，以保存其当前状态。备忘录对象通常只保存一些必要的状态信息，以避免浪费过多内存空间。当需要恢复发起者对象的状态时，发起者对象会通过传入备忘录对象来恢复其状态。

优点：

- 它能够在不破坏封装性原则的前提下，让对象在不同时间保存自身的状态，并在需要时恢复状态。这使得备忘录模式能够更加灵活地管理对象状态，从而提高代码的可维护性和可读性。

缺点：

- 可能会增加代码的复杂度。备忘录模式需要为每个需要被保存和恢复状态的对象都创建备忘录对象，并将其保存在管理者对象中。这可能会增加代码的复杂度和维护难度，从而降低代码的可读性和可维护性。



## 2. 理解

备忘录模式通过记录对象的状态信息，可以帮助我们保存和恢复对象的状态，也让我们能够方便地进行撤销操作。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/u=3393662197,1950252189&fm=253&fmt=auto&app=138&f=JPG.jpeg)

## 4. 实现

假设我们有一个游戏，它包含一个主角和一些敌人。主角的状态包括当前位置和当前生命值，而敌人的状态仅包括当前位置。现在，我们希望能够在游戏中实现撤销操作，即在主角移动或受到攻击后能够恢复到之前的状态。

这时，我们可以使用备忘录模式来管理主角的状态。具体来说，我们可以创建一个 `Memento` 类来保存主角的状态，然后再创建一个 `Caretaker` 类来管理备忘录对象。每当主角的状态发生改变时，我们都可以创建一个新的备忘录对象，并将其保存到 `Caretaker` 对象中。当需要撤销操作时，我们只需要从 `Caretaker` 中获取最新的备忘录对象，并使用它来恢复主角的状态即可。

```go
type Player struct {
	x, y int // position
	hp   int // health point
}

// Memento is the memento object, use to hold the player status
type Memento struct {
	Player
}

// Originator is the originator object, use to save or restore the player status
type Originator struct {
	Player
}

func (o *Originator) MoveTo(x, y int) {
	o.x, o.y = x, y
	fmt.Printf("player moves to (%d, %d)\n", x, y)
}

func (o *Originator) TakeDamage(damage int) {
	o.hp -= damage
	fmt.Printf("player takes %d damage, now the hp is %d\n", damage, o.hp)
}

// CreateMemento creates a memento according to current player's status
func (o *Originator) CreateMemento() *Memento {
	return &Memento{
		Player{
			x:  o.x,
			y:  o.y,
			hp: o.hp,
		},
	}
}

// RestoreMemento restores player's status from memento
func (o *Originator) RestoreMemento(m *Memento) {
	o.x, o.y, o.hp = m.x, m.y, m.hp
}

// CareTaker is the manager object
type CareTaker struct {
	mementos []*Memento
}

func (ct *CareTaker) AddMementos(ms ...*Memento) {
	ct.mementos = append(ct.mementos, ms...)
}

func (ct *CareTaker) GetLastMemento() *Memento {
	n := len(ct.mementos)
	if n == 0 {
		return nil
	}
	return ct.mementos[n-1]
}

func main() {
	// create a player
	p := Player{0, 0, 100}

	// create originator and manager
	originator := Originator{
		Player: p,
	}
	careTaker := &CareTaker{}

	// player move and save status
	originator.MoveTo(1, 1)
	careTaker.AddMementos(originator.CreateMemento())

	// player is damaged and save status
	originator.TakeDamage(20)

	// restore status
	m := careTaker.GetLastMemento()
	originator.RestoreMemento(m)
	fmt.Printf("player restores to (%d,%d), and hp is %d\n", originator.x, originator.y, originator.hp)
}
```



## 5. 场景

1. 当需要保存对象状态的历史记录时，可以使用备忘录模式来保存历史状态，方便进行回溯和恢复操作。
2. 当需要提供撤销操作时，可以使用备忘录模式来保存对象状态，以便在撤销操作时恢复之前的状态。
3. 当需要在不破坏对象封装性的情况下保存对象状态时，可以使用备忘录模式来实现状态的保存和恢复。
4. 当需要实现“只读快照”功能时，可以使用备忘录模式来保存对象的状态，并返回一个只读的快照，以保证数据的安全性。
5. 当需要实现事务的回滚功能时，可以使用备忘录模式来保存事务执行前的状态，并在回滚操作时恢复之前的状态。



## 参考

- ChatGPT