package pattern

import "fmt"

// Посетитель (visitor) — паттерн поведения объектов. Описывает операцию, выполняемую с каждым объектом
//  из некоторой структуры. Позволяет определить новую операцию, не изменяя классы этих объектов.

// Паттерн посетитель применяется:
// Когда классам необходимо добавить одинаковый набор операций без изменения этих классов.
// Когда часто добавляются новые операции к классам, при этом общая структура классов стабильна
//  и практически не изменяется.
// Когда имеется много объектов разнородных классов с разными интерфейсами, и требуется выполнить
// ряд операций над каждым из этих объектов.

// Приемущества паттерна посетитель:
// Упрощает добавление операций, работающих со сложными структурами объектов.
// Объединяет родственные операции в одном классе.
// Посетитель может накапливать состояние при обходе структуры элементов.

// Недостатки паттерна посетитель:
// Паттерн не оправдан, если иерархия элементов часто меняется.
// Может привести к нарушению инкапсуляции элементов.

type Visitor interface {
	visitForSquare(*Square)
	visitForCircle(*Circle)
	visitForRectangle(*Rectangle)
}

// Элемент для посещения.
type shape interface {
	GetType() string
	Accept(Visitor)
}

// Конкретный элемент для посещения.
type Square struct {
	side int
}

func (s *Square) GetType() string {
	return "Square"
}
func (s *Square) Accept(visitor Visitor) {
	visitor.visitForSquare(s)
}

// Конкретный элемент для посещения.
type Circle struct {
	radius int
}

func (c *Circle) GetType() string {
	return "Circle"
}
func (c *Circle) Accept(visitor Visitor) {
	visitor.visitForCircle(c)
}

// Конкретный элемент для посещения.
type Rectangle struct {
	length int
	width  int
}

func (r *Rectangle) GetType() string {
	return "Rectangle"
}
func (r *Rectangle) Accept(visitor Visitor) {
	visitor.visitForRectangle(r)
}

// Конкретный посетитель.
type AreaCalculator struct{}

func (a *AreaCalculator) visitForSquare(square *Square) {
	area := square.side * square.side
	fmt.Printf("Площадь квадрата со сторонами %d: %d\n", square.side, area)
}

func (a *AreaCalculator) visitForCircle(circle *Circle) {
	area := 3.14 * (float64(circle.radius) * float64(circle.radius))
	fmt.Printf("Площадь круга с радиусом %d: %f\n", circle.radius, area)
}

func (a *AreaCalculator) visitForRectangle(rectangle *Rectangle) {
	area := rectangle.length * rectangle.width
	fmt.Printf("Площадь прямоугольника со сторонами %d и %d: %d\n", rectangle.length, rectangle.width, area)
}
