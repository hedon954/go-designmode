# 观察者模式

观察者模式主要是用来实现事件驱动编程。事件驱动编程的应用还是挺广的，除了我们都知道的能够用来解耦：用户修改密码后，给用户发短信进行风险提示之类的典型场景，在微服务架构实现最终一致性、实现事件源（A + ES）这些都会用到。

## 1. 概念

观察者模式（Observer Pattern），定义对象间的一种 **一对多** 依赖关系，使得每当一个对象状态发生改变时，其相关依赖对象皆得到通知，依赖对象在收到通知后，可自行调用自身的处理程序，实现想要干的事情，比如更新自己的状态。

发布者对观察者唯一了解的是它实现了某个接口（观察者接口）。

这种松散耦合的设计最大限度地减少了对象之间的相互依赖，因此使我们能够构建灵活的系统。



## 2. 理解

观察者模式也经常被叫做 **发布 - 订阅（Publish/Subscribe）** 模式、上面说的定义对象间的一种一对多依赖关系：

- 一：指的是发布变更的主体对象
- 多：指的是订阅变更通知的订阅者对象

发布的状态变更信息会被包装到一个对象里，这个对象被称为 **事件**，事件一般用英语过去式的语态来命名，比如用户注册时，用户模块在用户创建好后发布一个事件 `UserCreated` 或者 `UserWasCreated` 都行，这样从名字上就能看出，这是一个已经发生过的事件。

事件发布给订阅者的过程，其实就是遍历一下已经注册的事件订阅者，逐个去调用订阅者实现的观察者接口方法，比如叫 `handleEvent` 之类的方法，这个方法的参数一般就是当前的事件对象。

至于很多人会好奇的，事件的处理是不是异步的？主要看我们的需求是什么，一般情况下是同步的，即发布事件后，触发事件的方法会阻塞等到全部订阅者返回后再继续，当然也可以让订阅者的处理异步执行，完全看我们的需求。

大部分场景下其实是同步执行的，单体架构会在一个数据库事务里持久化因为主体状态变更，而需要更改的所有实体类。

微服务架构下常见的做法是有一个事件存储，订阅者接到事件通知后，会把事件先存到事件存储里，这两步也需要在一个事务里完成才能保证最终一致性，后面会再有其他线程把事件从事件存储里搞到消息设施里，发给其他服务，从而在微服务架构下实现各个位于不同服务的实体间的最终一致性。

所以观察者模式，从程序效率上看，大多数情况下没啥提升，更多的是达到一种程序结构上的解耦，让代码不至于那么难维护。



## 3. 类图

![img](https://cdn.jsdelivr.net/gh/hedon954/mapStorage/img/0b17d3bd61844ebaab253163f6d6433f.png)





## 4. 最简单的观察者模式

- [demo1](demo/main.go)



## 5. Go 实现消息总线

- [eventbus](./eventbus/eventbus.go)



## 6. 总结

观察者模式在以下场景比较适用：如果一个对象状态改变要给其他对象通知，并且需要考虑到易用和低耦合，那么就可以使用观察者模式。

优点：

- 观察者和被观察者是抽象耦合的；
- 建立一套触发机制。

缺点：

- 如果一个被观察者对象有很多的直接和间接的观察者的话，将所有的观察者都通知到会花费很多时间；
- 如果在观察者和观察目标之间有循环依赖的话，观察目标会触发它们之间进行循环调用，可能导致系统崩溃；
- 观察者模式没有相应的机制让观察者知道所观察的目标对象是怎么发生变化的，而仅仅只是知道观察目标发生了变化；
- 如果顺序执行，某一观察者错误会导致系统卡壳，一般采用异步方式。



## 参考

- [拒绝 Go 代码臃肿，其实在这几块可以用下观察者模式](https://mp.weixin.qq.com/s?__biz=MzUzNTY5MzU2MA==&mid=2247495132&idx=1&sn=0c42ff03123e188c4de44df1b67ef4de&chksm=fa833c4bcdf4b55da316c82a7edaf33fbeca427116ab57caec1501993a8ef5905ce04f399de3&scene=178&cur_album_id=2531498848431669249#rd)
- [观察者模式（Go）](https://blog.csdn.net/a376240118/article/details/127025840)

