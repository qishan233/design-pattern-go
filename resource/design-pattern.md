# Design Pattern

Design patterns are basic elements of reusable object oriented software

## Creational Patterns

How to create something

### Abstract Factory

Provide an interface for creating families of related or dependent objects without specifying their concrete classes.

![abstract factory structure](https://raw.githubusercontent.com/qishan233/images/main/design-pattern/20240902092849.png)

### Builder

Separate the construction of a complex object from its representation so that the same construction process can create different representations.

### Factory Method

Define an interface for creating an object, but let subclasses decide which class to instantiate. Factory Method lets a class defer instantiation so subclasses.

### Prototype

Specify the kinds of objects to create using a prototypical instance, and create new objects by copying this prototype.

### Singleton

Ensure a class only has one instance, and provide a global point of access to it.

## Structural Patterns

Structural patterns are concerned with how classes and objects are composed to form larger structures.

### Adapter

Convert the interface of a class into another interface client expect. Adapter lets classes work together that couldn't otherwise because of incompatible interfaces.

### Bridge

Decouple the abstraction from its implementation so that the two can vary independently.

### Composite

Compose objects into tree structures to represent part-whole hierarchies. Composite lets clients treat individual objects and compositions of objects uniformly.

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
