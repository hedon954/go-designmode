# 外观模式

现代的软件系统都非常复杂，尽管我们已经想尽一切方法将其“分而治之”，把一个系统划分为好几个较小的子系统了，但是仍然可能会存在这样的问题：子系统内有非常多的类，客户端往往需要和许多对象打交道之后才能完成想要完成的功能。

在我们的生活中医院就是这样的。一般的医院都会分为挂号、门诊、化验、收费、取药等。看病的病人要想治好自己的病（相当于一个客户端想要实现自己的功能）就要和医院的各个部门打交道。首先，病人需要挂号，然后门诊，如果医生要求化验的话，病人就要去化验，然后再回到门诊室，最后拿药，经过一系列复杂的过程后才能完成看病的过程。如下图所示：

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20230327104559177.png)

如果我们在医院设立一个接待员的话，病人只负责和接待员接触，由接待员负责与医院的各个部门打交道，如下图所示：

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20230327104626345.png)

医院设立的接待员的角色就是我们今天要介绍的外观模式，系统通过引入外观模式让需要调用多个子系统各自部分的功能接口以完成的需求，变为调用方只需要跟外观提供的统一功能进行交互即可。

![图片](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/640-20230327104639436.png)

## 1. 概念

外观模式（Facade Pattern）是一种结构型设计模式，它提供了一个简单的接口，用于访问一个复杂的子系统。它封装了子系统中的一组接口，并对外提供一个统一的接口，客户端只需要调用这个接口即可完成复杂的操作。外观模式的主要目的是降低子系统与客户端之间的耦合度，让客户端更加方便地使用子系统。

外观模式中包含以下角色：

- Facade（外观类）：外观类是外观模式的核心，它封装了子系统的一组接口，提供了一个统一的接口供客户端使用。外观类通常是单例的，并且知道哪些子系统类负责处理哪些请求。
- Subsystem Classes（子系统类）：子系统类是实现子系统功能的类，它们包含了子系统的业务逻辑和实现细节。子系统类对外提供了一组接口，供外观类或其他类使用。
- Client（客户端）：客户端是使用外观模式的类，它通过调用外观类提供的接口来访问子系统。

外观模式的优点包括：

- 简化接口：外观模式提供了一个简单的接口，使得客户端能够方便地访问子系统，而不需要了解子系统的复杂性和实现细节。
- 降低耦合度：外观模式将子系统与客户端之间的耦合度降低到最低，使得子系统的变化不会影响到客户端。
- 提高可维护性：由于外观模式将子系统的复杂性封装起来，使得系统更加简单，易于维护。
- 提高灵活性：由于外观模式将子系统的复杂性封装起来，使得系统更加灵活，易于扩展和修改。

外观模式虽然有很多优点，但是也存在一些缺点，包括：

- 不符合开闭原则：外观模式的设计中，如果需要添加新的子系统类或者修改子系统类的接口，就需要修改外观类的代码，这违背了开闭原则，使得系统的可维护性降低。
- 可能会增加系统的复杂性：如果外观类设计不当，可能会导致外观类本身变得非常复杂，增加系统的复杂性。

外观模式适用于以下情况：

- 当一个系统比较复杂，有多个子系统，而且子系统之间存在依赖关系时，可以考虑使用外观模式，将子系统的复杂性封装起来，提供一个简单的接口供客户端使用。
- 当客户端需要访问系统中的多个子系统时，可以考虑使用外观模式，提供一个统一的接口，使得客户端能够方便地访问多个子系统。
- 当系统需要进行重构时，可以考虑使用外观模式，将系统中的复杂逻辑封装起来，使得重构更加方便。

## 2. 理解

外观模式其实就是通过多些一个“汇总”的类来屏蔽后台多个系统的接口，可以针对客户端的需要给出统一的接口，让客户端可以更简单的进行调用，同时屏蔽各个子系统的细节。

## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/20210309120505651.png)

## 4. 实现

假设我们正在开发一个游戏服务器，这个服务器包含了多个子系统，包括用户系统、背包系统、任务系统、副本系统等。

```go
// UserSystem is the system for operations of user
type UserSystem struct{}

// Login checks user's auth
func (us *UserSystem) Login(username, password string) bool {
	return true
}

// Register is used to register a new user
func (us *UserSystem) Register(username, password string) bool {
	return true
}

// BackpackSystem is the system to handle user's backpack's item
type BackpackSystem struct{}

// GetItems returns specific item
func (bs *BackpackSystem) GetItems(userId int) []string {
	return nil
}

// AddItem adds a new item to user's backpack
func (bs *BackpackSystem) AddItem(userId int, itemId string) bool {
	return true
}

// TaskSystem is the system to handle user's tasks
type TaskSystem struct{}

// GetTasks return all task of user
func (ts *TaskSystem) GetTasks(userId int) []string {
	return nil
}

// FinishTask ends the specific task
func (ts *TaskSystem) FinishTask(userId int, taskId string) bool {
	return true
}

// InstanceSystem is the instance system
type InstanceSystem struct{}

// EnterInstance used to enter the instance
func (is *InstanceSystem) EnterInstance(userId int, instanceId string) bool {
	return true
}

// ExitInstance used the exit the instance
func (is *InstanceSystem) ExitInstance(userId int) bool {
	return true
}
```

每个子系统都有自己的业务逻辑和实现细节，客户端需要访问多个子系统才能完成一次游戏操作，这时可以考虑使用外观模式来简化客户端的操作。我们可以定义一个游戏外观类 `GameFacade`，它封装了所有子系统的接口，客户端只需要调用 `GameFacade` 提供的接口就可以完成游戏操作。以下是示例代码：

```go
// GameFacade is the facade for the game system, client can communicate with the game server throught it
type GameFacade struct {
	userSystem     *UserSystem
	backpackSystem *BackpackSystem
	taskSystem     *TaskSystem
	instanceSystem *InstanceSystem
}

func NewGameFacade() *GameFacade {
	return &GameFacade{
		userSystem:     &UserSystem{},
		backpackSystem: &BackpackSystem{},
		taskSystem:     &TaskSystem{},
		instanceSystem: &InstanceSystem{},
	}
}

func (gf *GameFacade) Login(username, password string) bool {
	return gf.userSystem.Login(username, password)
}

func (gf *GameFacade) Register(username, password string) bool {
	return gf.userSystem.Register(username, password)
}

func (gf *GameFacade) GetItems(userId int) []string {
	return gf.backpackSystem.GetItems(userId)
}

func (gf *GameFacade) AddItem(userId int, itemId string) bool {
	return gf.backpackSystem.AddItem(userId, itemId)
}

func (gf *GameFacade) GetTasks(userId int) []string {
	return gf.taskSystem.GetTasks(userId)
}

func (gf *GameFacade) FinishTask(userId int, taskId string) bool {
	return gf.taskSystem.FinishTask(userId, taskId)
}

func (gf *GameFacade) EnterInstance(userId int, instanceId string) bool {
	return gf.instanceSystem.EnterInstance(userId, instanceId)
}

func (gf *GameFacade) ExitInstance(userId int) bool {
	return gf.instanceSystem.ExitInstance(userId)
}
```

在上面的示例中，我们定义了游戏服务器中的多个子系统，包括用户系统、背包系统、任务系统、副本系统等。然后，我们定义了一个游戏外观类 `GameFacade`，它封装了所有子系统的接口，客户端只需要调用 `GameFacade` 提供的接口就可以完成游戏操作。

更进一步，外观类中的一个方法可以封装子系统中的很多方法，屏蔽子系统的具体细节。通过使用外观模式，客户端可以方便地调用复杂的子系统，而不需要了解子系统的实现细节。

## 5. 场景

外观模式适合以下场景：

1. 需要简化复杂子系统的接口，提供更加简单易用的接口给客户端使用；
2. 需要将复杂的子系统解耦，避免客户端与子系统直接的耦合，提高代码的灵活性和可维护性；
3. 需要将子系统的变化隔离开来，避免对客户端产生影响，从而实现系统的可扩展性和可维护性；
4. 需要提供一个统一的接口来封装多个子系统，使得客户端可以更方便地使用多个子系统，减少客户端与子系统之间的交互次数。

在实际开发中，外观模式通常被用来封装复杂的第三方库、框架或者底层系统，提供更加简单的接口给客户端使用。例如，开发一个 Web 服务器时，可以使用外观模式封装多个子系统，包括网络层、数据存储层、业务逻辑层等，为客户端提供一个统一的接口来处理请求。

在游戏服务器开发中，外观模式也有很多应用场景，其中一些典型的场景包括：

1. 游戏引擎：游戏引擎是一个复杂的子系统，包括渲染引擎、物理引擎、音频引擎等多个子系统。使用外观模式可以将这些子系统封装起来，为游戏开发者提供一个统一的接口来使用游戏引擎。
2. 游戏存储：游戏服务器通常需要使用数据库或者其他数据存储系统来存储玩家数据、游戏配置等信息。使用外观模式可以将这些底层存储系统封装起来，为游戏开发者提供一个简单的接口来进行数据的读写操作。
3. 游戏逻辑：游戏服务器的核心部分是游戏逻辑，包括玩家行为的处理、游戏状态的管理等。使用外观模式可以将这些复杂的游戏逻辑封装起来，为游戏开发者提供一个简单的接口来实现游戏逻辑的处理。
4. 游戏服务：游戏服务器通常需要提供一些服务，如登录服务、匹配服务、战斗服务等。使用外观模式可以将这些服务封装起来，为游戏开发者提供一个统一的接口来访问这些服务。

Java 里的 Slf4j，它是一个抽象层，让用户对日志的操作统一由 Slf4j 跟用户去对接，用户用这个抽象层的 API 来写日志， 底层具体用什么日志工具实现用户完全不用关心，由 Slf4j 来对接 Log4j、LogBack 这些日志工具，这样就可以更方便地移植了。这个抽象层 Slf4j 就是Simple logging Facade For Java 的简称，从名字里我们也能看出来，它是一个外观模式的实践应用，由于普及度很高，很多讲解外观模式的教程里都会提及它，拿它的实现来给读者做分析。



## 参考

- ChatGPT
- [外观模式，一个每天都在用，却被多数人在面试中忽视的模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247497560&idx=1&sn=c3a8f8bc5e8d7748eb16d60a242369f2&chksm=fa8326cfcdf4afd91eab5dab7e556ae134c6b6a059913c3dbde0f7450db9faf8a74e5f1c9834&scene=178&cur_album_id=2531498848431669249#rd)