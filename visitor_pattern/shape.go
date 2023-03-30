package main

import (
	"fmt"
	"math"
)

// Shape is the interface of all kinds of shape
type Shape interface {
	Accept(Visitor)
}

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) Accept(visitor Visitor) {
	visitor.VisitForRectangle(r)
}

type Circle struct {
	radius float64
}

func (c *Circle) Accept(visitor Visitor) {
	visitor.VisitForCircle(c)
}

// Visitor is the interface of shape visitor, it provides methods to visit shapes
type Visitor interface {
	VisitForRectangle(*Rectangle)
	VisitForCircle(*Circle)
}

// AreaVisitor is a visitor to visit shape's area
type AreaVisitor struct {
	area float64
}

func (a *AreaVisitor) VisitForRectangle(rectangle *Rectangle) {
	a.area += rectangle.width * rectangle.height
}

func (a *AreaVisitor) VisitForCircle(circle *Circle) {
	a.area += circle.radius * circle.radius * math.Pi
}

// ShapeCollection is a collection of shapes
type ShapeCollection struct {
	shapes []Shape
}

func (a *ShapeCollection) AddShapes(shapes ...Shape) {
	a.shapes = append(a.shapes, shapes...)
}

func (a *ShapeCollection) Accept(visitor Visitor) {
	for _, shape := range a.shapes {
		shape.Accept(visitor)
	}
}

func main() {
	// creates rectangles
	r1 := &Rectangle{width: 3, height: 4}
	r2 := &Rectangle{width: 4, height: 5}

	// creates circles
	c1 := &Circle{radius: 3}
	c2 := &Circle{radius: 4}

	// creates collection
	coll := &ShapeCollection{}
	coll.AddShapes(r1, r2, c1, c2)

	// creates area visitor
	visitor := &AreaVisitor{}

	// use
	coll.Accept(visitor)
	fmt.Println(visitor.area)
}
