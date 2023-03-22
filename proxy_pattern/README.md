# 代理模式

## 1. 概念

代理模式是一种结构型设计模式，它允许通过代理对象控制对某个对象的访问。在代理模式中，代理对象充当了被代理对象的中间人，通过代理对象来访问被代理对象，并在必要时加入额外的逻辑。

代理模式通常被用来解决以下问题：

1. 远程代理：在分布式系统中，可以使用代理模式来实现远程服务调用。客户端通过代理对象访问远程主机上的对象，而不必了解对象所在的具体位置。
2. 虚拟代理：在创建开销较大的对象时，可以使用代理模式来实现延迟加载。在需要时才会真正创建对象，而在此之前会使用一个轻量级的代理对象代替。
3. 安全代理：在需要进行安全检查的场景下，可以使用代理模式来实现权限控制。代理对象可以在访问被代理对象之前进行安全检查，以确保客户端有权限访问被代理对象。
4. 缓存代理：在需要缓存数据的场景下，可以使用代理模式来实现缓存。代理对象可以在访问被代理对象之前，检查缓存是否存在需要的数据，如果存在，则返回缓存中的数据，否则调用被代理对象获取数据并缓存。
5. 日志记录代理：在需要记录日志的场景下，可以使用代理模式来实现日志记录。代理对象可以在访问被代理对象之前，记录访问日志。
6. 智能引用代理：在需要实现引用计数的场景下，可以使用代理模式来实现引用计数。代理对象可以记录被引用对象的引用次数，并在引用次数为0时，自动销毁被引用对象。

代理模式由以下角色组成：

1. 抽象主题（Subject）：定义被代理对象和代理对象的共同接口，客户端只与抽象主题进行交互。
2. 具体主题（RealSubject）：实现抽象主题接口，是真正的被代理对象。
3. 代理（Proxy）：实现抽象主题接口，并维护一个指向具体主题的引用，通过代理对象来控制对具体主题的访问。

代理模式的优点包括：

1. 代理模式使得客户端不必直接与被代理对象交互，从而降低了耦合度。
2. 代理模式可以在不改变被代理对象的情况下，实现额外的逻辑，例如缓存、安全检查等。
3. 代理模式可以实现远程访问，使得客户端可以访问远程主机上的对象，从而具有更大的灵活性和扩展性。
4. 代理模式可以实现虚拟代理，延迟对象的创建时间，从而提高了系统的性能。

代理模式的缺点包括：

1. 代理模式增加了代码的复杂度，增加了系统的开销。
2. 代理模式可能会导致系统的响应时间变慢，特别是在访问远程主机上的对象时。



## 2. 理解

代理模式其实就是对接口的完美应用。代理对象通过跟服务对象实现同样的接口，做到调用方无感知的方法增强功能。通过代理对象，我们可以对服务对象的某些方法很方便地进行增强。



## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/u=643016613,3753330583&fm=253&fmt=auto&app=138&f=JPEG-20230322152104639.jpeg)



## 4. 实现

下面以一个游戏服务器开发为例子来演示一下代理模式的具体应用。

### 4.1 定义接口

```go
type GameServer interface {
	Connect() error        // connect to the game server
	Disconnect() error     // disconnect from the game server
	Send(msg string) error // send message to the game server
}
```

### 4.2 服务对象

```go
// GameServerImpl is a implements of the GameServer
type GameServerImpl struct {
	address string
}

func (gs *GameServerImpl) Connect() error {
	// build the connection with game server
	return nil
}

func (gs *GameServerImpl) Disconnect() error {
	// cut the connection with game server
	return nil
}

func (gs *GameServerImpl) Send(message string) error {
	// send message to the game server
	return nil
}
```

### 4.3 代理对象

```go
// GameServerProxy is a proxy of GameServerImpl
type GameServerProxy struct {
	gameServer GameServer // service object
	username   string
	password   string
}

func (gsp *GameServerProxy) Connect() error {
	// check auth
	if gsp.username != "admin" || gsp.password != "password" {
		return fmt.Errorf("authentication failed")
	}
	if err := gsp.gameServer.Connect(); err != nil {
		return err
	}
	return nil
}

func (gsp *GameServerProxy) Disconnect() error {
	// clear cache
	gsp.username = ""
	gsp.password = ""
	if err := gsp.gameServer.Disconnect(); err != nil {
		return err
	}
	return nil
}

func (gsp *GameServerProxy) Send(message string) error {
	// log
	log.Printf(gsp.username, message)
	if err := gsp.gameServer.Send(message); err != nil {
		return err
	}
	return nil
}
```

### 4.4 使用

```go
func main() {
	gameServer := &GameServerImpl{address: "localhost:12345"}
	gameServerProxy := &GameServerProxy{
		gameServer: gameServer,
		username:   "admin",
		password:   "password",
	}
	if err := gameServerProxy.Connect(); err != nil {
		fmt.Printf("error connecting to game server: %v", err)
		return
	}
	defer gameServerProxy.Disconnect()

	if err := gameServerProxy.Send("hello"); err != nil {
		fmt.Printf("error sending message to game server: %v", err)
		return
	}
}
```






## 参考

- ChatGPT