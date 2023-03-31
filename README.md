# go-designmode
Some examples of design modes implemented with Go.

用 Go 实现 23 种设计模式

## 创建型模式

- [工厂方法模式](./factory_pattern)
- [抽象工厂模式](./factory_pattern)
- [单例模式](./singleton_pattern)
- [建造者模式](./builder_pattern)
- [原型模式](./singleton_pattern)



## 结构性模式

- [适配器模式](./adapter_pattern)
- [桥接模式](./bridge_pattern)
- [组合模式](./composite_pattern)
- [装饰者模式](./decorator_pattern)
- [外观模式](./facade_pattern)
- [享元模式](./flyweight_pattern)
- [代理模式](./proxy_pattern)



## 行为型模式

- [策略模式](./strategy_pattern)
- [模板方法模式](./template_pattern)
- [观察者模式](./observer_pattern)
- [迭代器模式](./iterator_pattern)
- [职责链模式](./responsibility_pattern)
- [命令模式](./command_pattern)
- [备忘录模式](./memento_pattern)
- [状态模式](./state_pattern)
- [访问者模式](./visitor_pattern)
- [中介者模式](./mediator_pattern)
- [解释器模式](./interpreter_pattern)

## 其他

- [Option 编程模式](./options_pattern)

## 概括

### 三个维度

1. 创建型模式 (Creational Patterns)：这些模式提供了不同种类的对象创建机制，使得一个系统在运行时可以选择其中的一个适当的创建方法来创建对象。
2. 结构型模式 (Structural Patterns)：这些模式描述如何将类或对象组合成更大的结构，以满足特定的需求。
3. 行为型模式 (Behavioral Patterns)：这些模式涉及到算法和对象间职责的分配，并描述了在对象之间的通信模式。

### 23 种

1. 单例模式 (Singleton Pattern)：确保一个类只有一个实例，并提供全局访问点来访问该实例。
2. 工厂模式 (Factory Pattern)：定义一个用于创建对象的接口，让子类决定实例化哪个类来创建对象。
3. 抽象工厂模式 (Abstract Factory Pattern)：提供一个创建一系列相关或相互依赖对象的接口，而无需指定它们具体的类。
4. 建造者模式 (Builder Pattern)：将一个复杂对象的构造与其表示分离，使得同样的构造过程可以创建不同的表示。
5. 原型模式 (Prototype Pattern)：通过复制现有的实例来创建新实例。
6. 适配器模式 (Adapter Pattern)：将一个类的接口转换成客户希望的另外一个接口。适配器模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。
7. 桥接模式 (Bridge Pattern)：将抽象部分与它的实现部分分离，使得它们都可以独立地变化。
8. 装饰器模式 (Decorator Pattern)：动态地给一个对象添加一些额外的职责。就增加功能而言，装饰器模式比生成子类更为灵活。
9. 组合模式 (Composite Pattern)：将对象组合成树形结构以表示“部分-整体”的层次结构。
10. 外观模式 (Facade Pattern)：为子系统中的一组接口提供一个一致的界面，该模式定义了一个高层接口，这个接口使得这一子系统更加容易使用。
11. 享元模式 (Flyweight Pattern)：运用共享技术有效地支持大量细粒度的对象。
12. 代理模式 (Proxy Pattern)：为其他对象提供一种代理以控制对这个对象的访问。
13. 责任链模式 (Chain of Responsibility Pattern)：使多个对象都有机会处理请求，从而避免请求的发送者和接收者之间的耦合关系。将这些对象连成一条链，并沿着这条链传递该请求，直到有一个对象处理它为止。
14. 命令模式 (Command Pattern)：将请求封装成对象，从而让你使用不同的请求、队列或者日志来参数化其它对象。命令模式也可以支持撤销操作。
15. 解释器模式 (Interpreter Pattern)：给定一个语言，定义它的文法的一种表示，并定义一个解释器，该解释器使用该表示来解释语言中的句子。
16. 迭代器模式 (Iterator Pattern)：提供一种方法顺序访问一个聚合对象中的各个元素，而又不暴露该对象的内部表示。
17. 中介者模式 (Mediator Pattern)：用一个中介对象封装一系列的对象交互。中介者使得各个对象之间不需要显式地相互引用，从而使其耦合松散，而且可以独立地改变它们之间的交互。
18. 备忘录模式 (Memento Pattern)：在不破坏封装性的前提下，捕获一个对象的内部状态，并在该对象之外保存这个状态。这样以后就可以将该对象恢复到原先保存的状态。
19. 观察者模式 (Observer Pattern)：定义了对象之间的一对多依赖关系，这样一来，当一个对象改变状态时，它的所有依赖者都会收到通知并自动更新。
20. 状态模式 (State Pattern)：允许对象在内部状态改变时改变它的行为，对象看起来似乎修改了它所属的类。
21. 策略模式 (Strategy Pattern)：定义一系列算法，把它们一个个封装起来，并使它们可以相互替换。本模式使得算法的变化可独立于使用它的客户端。
22. 模板方法模式 (Template Method Pattern)：定义一个操作中的算法的骨架，而将一些步骤延迟到子类中。模板方法使得子类可以不改变一个算法结构即可重定义该算法的某些特定步骤。
23. 访问者模式 (Visitor Pattern)：表示一个作用于某对象结构中的各元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。
