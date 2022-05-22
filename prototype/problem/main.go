package main

import (
	"fmt"

	"github.com/arthurshafikov/patterns/prototype/problem/shapes"
)

func main() {
	shapesArr := []any{
		shapes.NewCircle(10, 4, 5),
		shapes.NewRectangle(1, 1, 10, 25),
		shapes.NewCircle(10, 20, 3),
		shapes.NewRectangle(9, 2, 24, 12),
	}
	shapesArr[1].(*shapes.Rectangle).SetOpacity(50)
	shapesArr[2].(*shapes.Circle).SetOpacity(100) // these opacity values cannot be copied in main package at all.

	// we need to create a copy of every shape
	var shapesCopy []any

	for _, shape := range shapesArr {
		if circle, ok := shape.(*shapes.Circle); ok {
			circleCopy := shapes.NewCircle(circle.X, circle.Y, circle.Radius)
			// imagine if there were way more properties to copy
			shapesCopy = append(shapesCopy, circleCopy)
		}
		if rectangle, ok := shape.(*shapes.Rectangle); ok {
			rectangleCopy := shapes.NewRectangle(rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height)
			shapesCopy = append(shapesCopy, rectangleCopy)
		}
	}

	ListAll(shapesArr)
	/*
		&shapes.Circle{X:10, Y:4, Radius:5, opacity:0}, &shapes.Rectangle{X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{X:10, Y:20, Radius:3, opacity:100}, &shapes.Rectangle{X:9, Y:2, Width:24, Height:12, opacity:0},
	*/
	ListAll(shapesCopy)
	/*
		&shapes.Circle{X:10, Y:4, Radius:5, opacity:0}, &shapes.Rectangle{X:1, Y:1, Width:10, Height:25, opacity:0},
		&shapes.Circle{X:10, Y:20, Radius:3, opacity:0}, &shapes.Rectangle{X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
	shapesArr[0].(*shapes.Circle).SetOpacity(22)
	ListAll(shapesArr)
	/*
		&shapes.Circle{X:10, Y:4, Radius:5, opacity:22}, &shapes.Rectangle{X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{X:10, Y:20, Radius:3, opacity:100}, &shapes.Rectangle{X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
	ListAll(shapesCopy) // opacity has not changed in this slice, because circle is a copy, that's exactly what we want.
	/*
		&shapes.Circle{X:10, Y:4, Radius:5, opacity:0}, &shapes.Rectangle{X:1, Y:1, Width:10, Height:25, opacity:0},
		&shapes.Circle{X:10, Y:20, Radius:3, opacity:0}, &shapes.Rectangle{X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
}

func ListAll(arr []any) {
	for _, v := range arr {
		fmt.Printf("%#v, ", v)
	}
	fmt.Println()
}
