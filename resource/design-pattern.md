# Design Pattern

Design patterns are basic elements of reusable object oriented software

## Creational Patterns

Creational design patterns abstract the instantiation process.They help make a system independent of how its objects are created,composed, and represented. A class creational pattern uses inheritance to vary the class that's instantiated, whereas an object creational pattern will delegate instantiation to another object.

创建型设计模式抽象了实例化过程。它们帮助一个系统独立于如何创建、组合和表示它的那些对象。一个类创建型模式使用继承改变被实例化的类，而一个对象创建型模式将实例化委托给另一个对象。

Creational patterns become important as systems evolve to depend more on object composition than class inheritance. As that happens,emphasis shifts away from hard-coding a fixed set of behaviors toward defining a smaller set of fundamental behaviors that can be composed into any number of more complex ones. Thus creating objects with particular behaviors requires more than simply instantiating a class.

随着系统演化得越来越依赖于对象组合而不是类继承，创建型模式变得更为重要。当这种情况发生时，重心从对一组固定行为的硬编码(hard-coding)转移为定义一个较小的基本行为集，这些行为可以被组合成任意数目的更复杂的行为。这样创建有特定行为的对象要求的不仅仅是实例化一个类。

There are two recurring themes in these patterns. First, they all encapsulate knowledge about which concrete classes the system uses.Second, they hide how instances of these classes are created and put together. All the system at large knows about the objects is their interfaces as defined by abstract classes. Consequently, the creational patterns give you a lot of flexibility in what gets created, who creates it, how it gets created, and when. They let you configure a system with "product" objects that vary widely in structure and functionality. Configuration can be static (that is, specified at compile-time) or dynamic(at run-time).

在这些模式中有两个不断出现的主旋律。第一，它们都将关于该系统使用哪些具体的类的信息封装起来。第二，它们隐藏了这些类的实例是如何被创建和放在一起的。整个系统关于这些对象所知道的是由抽象类所定义的接口。因此，创建型模式在什么被创建、谁创建它、它是怎样被创建的，以及何时创建等方面给予你很大的灵活性。它们允许你用结构和功能差别很大的“产品”对象配置一个系统。配置可以是静态的（即在编译时指定）​，也可以是动态的（在运行时指定）​。

### Abstract Factory

Provide an interface for creating families of related or dependent objects without specifying their concrete classes.

提供一个接口以创建一系列相关或相互依赖的对象，而无须指定它们具体的类

![Abstract Factory structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240902092849.png)

### Builder

Separate the construction of a complex object from its representation so that the same construction process can create different representations.

将一个复杂对象的构建与它的表示分离，使得同样的构建过程可以创建不同的表示；

![Builder structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240903205359.png)

### Factory Method

Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation so subclasses.

定义一个用于创建对象的接口，让子类决定实例化哪一个类。Factory Method使一个类的实例化延迟到其子类

![Factory Method structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240904084242.png)

### Prototype

Specify the kinds of objects to create using a prototypical instance, and create new objects by copying this prototype.

用原型实例指定创建对象的种类，并且通过拷贝这些原型创建新的对象。

![Prototype structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240904090625.png)

### Singleton

Ensure a class only has one instance, and provide a global point of access to it.

保证一个类仅有一个实例，并提供一个访问它的全局访问点。

![Singleton structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240904092510.png)

## Structural Patterns

Structural patterns are concerned with how classes and objects are composed to form larger structures.

结构型模式涉及如何组合类和对象以获得更大的结构。

### Adapter

Convert the interface of a class into another interface client expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.

将一个类的接口转换成客户希望的另外一个接口。Adapter 模式使得原本由于接口不兼容而不能一起工作的那些类可以一起工作。

Class Adapter structure:

![Class Adapter structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240905211455.png)

Object Adapter structure:

![Object Adapter structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240905211617.png)

### Bridge

Decouple the abstraction from its implementation so that the two can vary independently.

将抽象部分与它的实现部分分离，使它们可以独立地变化。

![Bridge structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240905214602.png)

### Composite

Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.

将对象组合成树形结构以表示“部分–整体”的层次结构。Composite 使得用户对单个对象和组合对象的使用具有一致性。

![Composite structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240905220053.png)

### Decorator

Attach addition responsibilities to an object dynamically. Decorators provide a flexible alternative to subclassing for extending functionality.

### Facade

Provide a unified interface to a set of interfaces in a subsystem. Facade defines a high-level interface that make the subsystem easier to use.

### Flyweight

Use sharing to support large numbers of fine-grained object efficiently.

### Proxy

Provide a surrogate or placeholder for another object to control access to it.

## Behavioral Patterns

Behavioral patterns are concerned with algorithms and the assignment of responsibilities between objects. Behavioral patterns describe not just patterns of objects or classes but also the patterns of communication between them.

### Chain of Responsibility

Avoid coupling the sender of a request to its receiver by giving more than one object a chance to handle the request. Chain the receiving objects and pass the request along the chain until an object handle it.

### Command

Encapsulate a request as an object, thereby letting you parameterize clients with different requests, queue or log requests, and support undoable operations.

### Interpreter

Given a language, define a representation for its grammar along with an interpreter that uses the representation to interpret sentence in th language.

### Iterator

Provide a way to access the elements of an aggregated object sequentially without expose its underlying representation.

### Mediator

Define an object that encapsulate how a set of objects interact. Mediator promotes loose coupling by keeping objects from referring to each other explicitly, and it let you vary their interaction independently.

### Memento

Without violating encapsulation, capture and externalize an object's internal state so that the object can be restored to this state later.

### Observer

Define a one-to-many dependency between objects so that when one object changes state, all its dependents are notified and updated automatically.

### State

Allow an object to alter its behavior when its internal state changes. The object will appear to change its class.

### Strategy

Define a family of algorithms, encapsulate each one, and make them interchangeable. Strategy let the algorithm vary independently from clients that use it.

### Template Method

Define the skeleton of an algorithm in an operation, deferring some steps to subclasses. Template Method lets subclasses redefine certain steps of an algorithm without changing the algorithm's structure.

### Visitor

Represent an operation to be performed on the elements of an object structure. Visitor let you define a new operation without changing the classes of the elements on which it operates.
