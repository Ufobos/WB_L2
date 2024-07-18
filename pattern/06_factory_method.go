package pattern

import "fmt"

// Фабричный метод (Factory Method) — это порождающий паттерн проектирования, который решает проблему создания
// различных продуктов, без указания конкретных классов продуктов.

// Паттерн фабричный метод применяется:
// Когда заранее неизвестно, объекты каких типов необходимо создавать.
// Когда система должна быть независимой от процесса создания новых объектов и расширяемой: в нее можно
// легко вводить новые классы, объекты которых система должна создавать.
// Когда создание новых объектов необходимо делегировать из базового класса классам наследникам.

// Преимущества паттерна фабричный метод:
// Избавляет класс от привязки к конкретным классам продуктов.
// Выделяет код производства продуктов в одно место, упрощая поддержку кода.
// Упрощает добавление новых продуктов в программу.

// Недостатки паттерна фабричный метод:
// Для каждого нового продукта необходимо создавать свой класс создателя.

// Интерфейс продукта
type Product interface {
	someMethod()
}

// Конкретный продукт
type ConcreteProduct struct {
	name string
}

func (c *ConcreteProduct) someMethod() {
	fmt.Println(c.name)
}

// Подкласс конкретного продукта
type ConcreteProductA struct {
	ConcreteProduct
}

func NewConcreteProductA(name string) *ConcreteProductA {
	concreteProductA := new(ConcreteProductA)
	concreteProductA.name = name
	return concreteProductA
}

// Подкласс конкретного продукта
type ConcreteProductB struct {
	ConcreteProduct
}

func NewConcreteProductB(name string) *ConcreteProductB {
	concreteProductB := new(ConcreteProductB)
	concreteProductB.name = name
	return concreteProductB
}

// Фабрика
func GetProduct(productType string, name string) (Product, error) {
	switch productType {
	case "ConcreteProductA":
		return NewConcreteProductA(name), nil
	case "ConcreteProductB":
		return NewConcreteProductB(name), nil
	default:
		return nil, fmt.Errorf("не правильный тип продукта")
	}
}
