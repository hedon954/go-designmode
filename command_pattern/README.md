# 命令模式





## 1. 概念

命令模式（Command Pattern）是一种行为设计模式，它使得可以将请求封装为一个独立的对象，从而可以将请求的发送者与接收者解耦。通过命令对象，可以将请求的参数化、延迟执行或者将多个请求组合成一个复合命令，从而提高系统的灵活性和可扩展性。

在命令模式中，通常包含以下角色：

- 命令接口(Command Interface)：定义了命令对象的基本接口，包括执行、撤销等操作；
- 具体命令(Concrete Command)：实现了命令接口，并且包含了具体的执行逻辑；
- 命令发送者(Command Sender)：负责创建和发送命令对象；
- 命令接收者(Command Receiver)：负责接收和执行命令对象。

## 2. 理解

命令模式的核心思想是将请求发送者和请求接收者解耦，这样可以让不同的请求发送者发送不同的请求，并且请求接收者也可以按照自己的方式进行处理，从而使整个系统更加灵活和可扩展。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/231247427711364.png)

## 4. 实现

### 4.1 定义命令接口

```go
type Command interface {
	Execute()
}
```

### 4.2 定义 player

```go
type Player struct {
	x      int
	y      int
	health int
}

func (p *Player) Move(direction string) {
	switch direction {
	case "up":
		p.y++
	case "down":
		p.y--
	case "left":
		p.x--
	case "right":
		p.x++
	}
}

func (p *Player) Attack(target *Player) {
	target.health -= 10
}
```

### 4.3 实现两个具体的游戏命令

```go
type MoveCommand struct {
	player    *Player
	direction string
}

func (c *MoveCommand) Execute() {
	c.player.Move(c.direction)
}

type AttackCommand struct {
	palyer *Player
	target *Player
}

func (c *AttackCommand) Execute() {
	c.palyer.Attack(c.target)
}
```

### 4.4 定义游戏服务器

```go
type GameServer struct {
	commands []Command
}

func (s *GameServer) AddCommands(commands ...Command) {
	s.commands = append(s.commands, commands...)
}

func (s *GameServer) ProcessCommands() {
	for _, c := range s.commands {
		c.Execute()
	}
}
```

### 4.5 使用

```go
func main() {
	server := GameServer{}

	player1 := &Player{}
	player2 := &Player{}

	moveCommand1 := &MoveCommand{player1, "right"}
	attackCommand1 := &AttackCommand{player1, player2}

	moveCommand2 := &MoveCommand{player2, "up"}
	attackCommand2 := &AttackCommand{player2, player1}

	server.AddCommands(moveCommand1, attackCommand1, moveCommand2, attackCommand2)

	server.ProcessCommands()

	fmt.Println(player1)
	fmt.Println(player2)
}
```

在这个示例中，我们创建了两个玩家并为它们创建了移动和攻击命令。然后将这些命令添加到游戏服务器中，并执行它们。这样，玩家就会朝着指定的方向移动并攻击另一个玩家。



## 5. 场景

命令模式的主要目的是将请求封装成对象，从而使得不同的请求能够被参数化，支持撤销和重做等操作。以下是一些常见的使用场景：

1. 用户界面操作

   命令模式可以用于实现用户界面中的各种操作，例如撤销、重做、剪切、复制、粘贴等。在这种情况下，命令对象可以封装具体的用户操作，然后将其添加到一个命令队列中，等待执行。

2. 任务调度

   命令模式可以用于实现任务调度系统，例如在一个分布式系统中，任务可以由多个节点执行。在这种情况下，命令对象可以封装一个任务，然后将其添加到一个任务队列中，等待执行。

3. 事务处理

   命令模式可以用于实现事务处理系统，例如在数据库操作中，每个操作都可以封装成一个命令对象。在这种情况下，命令对象可以支持撤销和重做操作，并且可以保证每个操作都被正确执行。

4. 日志记录

   命令模式可以用于实现日志记录系统，例如在一个复杂的系统中，需要记录每个用户操作以便进行调试和排错。在这种情况下，命令对象可以封装每个用户操作，并将其添加到一个日志队列中，等待记录。

5. 消息队列

   命令模式可以用于实现消息队列系统，例如在一个分布式系统中，节点之间需要通过消息进行通信。在这种情况下，命令对象可以封装一条消息，并将其添加到一个消息队列中，等待传递。

总之，命令模式可以用于任何需要将请求封装成对象的场景，从而提供更加灵活和可扩展的设计方案。



## 参考

- ChatGPT