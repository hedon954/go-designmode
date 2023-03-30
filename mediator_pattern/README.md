# 中介者模式

## 1. 概念

中介者模式（Mediator Pattern）是一种行为型设计模式，它的主要作用是将对象之间的通信行为进行抽象和封装，从而使得对象之间的通信更加松散和灵活。

在中介者模式中，对象之间的通信不是直接发生的，而是通过一个中介者对象来进行协调和管理。中介者对象封装了对象之间的通信逻辑，将对象之间的交互行为从对象本身中解耦出来，使得对象只需要知道如何和中介者对象进行通信，而不需要知道如何和其他对象进行通信。这样就降低了对象之间的耦合度，使得对象更加灵活和易于复用。

在中介者模式中，通常会定义一个抽象中介者类和多个具体的同事类。抽象中介者类中定义了各个同事类之间的通信接口，而具体的中介者类则实现了这些接口，并负责协调和管理同事类之间的通信行为。每个同事类都需要持有一个中介者对象的引用，用于和中介者对象进行通信。

中介者模式的通常包含以下角色：

- Mediator（中介者）：定义了对象之间进行通信和协作的接口。
- ConcreteMediator（具体中介者）：实现了中介者接口，并负责协调各个对象之间的交互关系。
- Colleague（同事）：表示需要和其他对象进行交互的对象，持有一个指向中介者的引用。
- ConcreteColleague（具体同事）：实现了同事接口，并负责和其他对象进行交互。

优点：

1. 降低了对象之间的耦合度。将对象之间的通信逻辑封装在中介者对象中，使得对象之间不再直接依赖于其他对象，从而减少了对象之间的耦合度。
2. 使得对象更加灵活和易于复用。由于对象之间的通信逻辑被抽象和封装起来，因此对象变得更加灵活和易于复用。
3. 简化了对象之间的交互逻辑。中介者对象封装了对象之间的通信逻辑，使得对象之间的交互变得更加简单和清晰。

缺点：

1. 中介者对象可能会变得比较复杂。由于中介者对象需要管理多个同事类之间的通信行为，因此中介者对象可能会变得比较复杂。
2. 同事类之间的通信会变得间接。由于同事类之间的通信不再直接发生，而是通过中介者对象进行协调和管理，因此同事类之间的通信会变得间接。

中介者模式适用于以下情况：

1. 系统中的对象之间存在较为复杂的交互关系，导致系统难以维护和扩展。
2. 系统中的对象之间存在循环依赖关系，导致系统难以进行单元测试和调试。
3. 系统中的对象之间的通信行为难以复用和扩展，导致系统难以适应变化。

## 2. 理解

中介者模式就是对象与对象之前不直接进行交互，而是通过中介来进行交互以实现对象之间的解耦。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/v2-d3a8d03ae8fa16c375fa4354bde5edd8_720w.webp)

## 4. 实现

假设我们有一个简单的聊天系统，有多个用户可以在该系统中进行聊天。每个用户可以向其他用户发送消息，当一个用户发送消息时，其他所有用户都能够接收到这条消息。

这个聊天系统可以使用中介者模式来实现。中介者对象将负责维护所有用户之间的通信，并将接收到的消息广播给所有用户。每个用户只需要将消息发送给中介者对象，而不需要知道其他用户的存在，从而降低了系统的耦合度。

```go
// Mediator is the chat room mediator
type Mediator interface {
	sendMessage(user User, message string)
}

// User is a user of the chat room,
// it holds the mediator and communicates with it
type User struct {
	name     string
	mediator Mediator
}

// NewUser creates a new user
func NewUser(name string, mediator Mediator) *User {
	return &User{
		name:     name,
		mediator: mediator,
	}
}

// SendMessage sends a message
// it just needs to send msg to the mediator
// needs not to cara about how many users there are
func (u *User) SendMessage(msg string) {
	u.mediator.sendMessage(*u, msg)
}

// ChatRoom is the chat room, it implements the Mediator
type ChatRoom struct {
	users []*User
}

func (c *ChatRoom) sendMessage(user User, message string) {
	for _, u := range c.users {
		if u.name != user.name {
			fmt.Printf("%s recevied the message [%s] from %s\n", u.name, message, user.name)
		}
	}
}

func (c *ChatRoom) AddUsers(users ...*User) {
	c.users = append(c.users, users...)
}

func (c *ChatRoom) RemoveUsers(names ...string) {
	for _, name := range names {
		for i, u := range c.users {
			if u.name == name {
				c.users = append(c.users[:i], c.users[i+1:]...)
				break
			}
		}
	}
}

func main() {
	room := &ChatRoom{}
	u1 := &User{"Alice", room}
	u2 := &User{"Bob", room}
	u3 := &User{"Charlie", room}
	room.AddUsers(u1, u2, u3)

	u1.SendMessage("hello")
	// Bob recevied the message [hello] from Alice
	// Charlie recevied the message [hello] from Alice

	u2.SendMessage("hi")
	// Alice recevied the message [hi] from Bob
	// Charlie recevied the message [hi] from Bob

	u3.SendMessage("here")
	// Alice recevied the message [here] from Charlie
	// Bob recevied the message [here] from Charlie
}
```

ChatRoom 结构体作为中介者，实现了 Mediator 接口。在 ChatRoom 中，维护了一个 users 字段，用于存储所有加入聊天室的用户。当用户发送消息时，ChatRoom 会将消息广播给所有其他用户。

具体来说，ChatRoom 结构体实现了 Mediator 接口中的 Send 方法。当用户发送消息时，ChatRoom 会调用 Send 方法，将消息作为参数传递给 Send 方法。Send 方法会遍历所有用户，并将消息发送给除当前用户以外的其他所有用户。这样，所有用户都可以收到这条消息，完成了消息的中转和传递的过程。

此外，ChatRoom 还提供了 AddUser 方法和 RemoveUser 方法，用于向聊天室添加或删除用户。在用户加入或离开聊天室时，ChatRoom 会更新 users 字段，保证聊天室中的用户列表是最新的。

如果不使用中介者模式，那么每个用户对象都需要知道其他所有用户对象的存在，并且需要维护一个列表来保存所有用户。当一个用户向聊天室中发送消息时，需要将消息发送给所有其他用户。这样的设计会使得用户对象之间的耦合度非常高，而且增加新的用户对象或者删除已有的用户对象时，都需要修改所有其他用户对象的代码，非常不便于维护。

另外，这种设计方式也会使得用户对象的职责过于复杂，不利于单一职责原则的实现，代码的可扩展性和可维护性也会受到影响。

因此，在这种场景下使用中介者模式可以有效地降低对象之间的耦合度，使得代码更加易于维护和扩展。

## 5. 场景

中介者模式适用于以下场景：

1. 对象之间的交互比较复杂：当对象之间的交互变得复杂时，中介者可以起到简化交互的作用，避免对象之间出现复杂的耦合关系，使得系统更加灵活、可维护。
2. 对象之间的交互需要扩展性：当系统中需要新增、删除、修改对象之间的交互时，中介者可以起到统一调度的作用，可以方便地进行修改和扩展，而不需要修改每个对象之间的交互逻辑，提高了系统的可扩展性。
3. 有大量对象之间的交互：当系统中存在大量对象之间的交互时，中介者可以起到减少交互次数、降低系统复杂度的作用。如果不使用中介者模式，每个对象都需要与其他所有对象进行交互，会造成大量的交互次数和复杂度，而中介者可以起到集中调度、降低交互次数的作用。
4. 需要对对象之间的交互进行控制：当系统需要对对象之间的交互进行控制时，中介者可以起到控制作用。中介者可以限制对象之间的交互，或者加入一些额外的逻辑，以满足系统的需求。

总的来说，中介者模式适用于对象之间的交互比较复杂、需要扩展性、存在大量对象之间的交互、需要对交互进行控制等场景。

## 参考

- ChatGPT