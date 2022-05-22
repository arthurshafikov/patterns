package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arthurshafikov/patterns/prototype/problem/shapes"
)

func main() {
	rand.Seed(time.Now().Unix())

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
			// We could do that
			// circleCopy := *circle
			// But we need UniqueID field to stay unique and not to be copied.
			shapesCopy = append(shapesCopy, circleCopy)
		}
		if rectangle, ok := shape.(*shapes.Rectangle); ok {
			rectangleCopy := shapes.NewRectangle(rectangle.X, rectangle.Y, rectangle.Width, rectangle.Height)
			// imagine if there were way more properties to copy
			shapesCopy = append(shapesCopy, rectangleCopy)
		}
	}

	ListAll(shapesArr)
	/*
		&shapes.Circle{uniqueID:37, X:10, Y:4, Radius:5, opacity:0},
		&shapes.Rectangle{uniqueID:23, X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{uniqueID:24, X:10, Y:20, Radius:3, opacity:100},
		&shapes.Rectangle{uniqueID:82, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
	ListAll(shapesCopy) // opacity is not copied...
	/*
		&shapes.Circle{uniqueID:77, X:10, Y:4, Radius:5, opacity:0},
		&shapes.Rectangle{uniqueID:59, X:1, Y:1, Width:10, Height:25, opacity:0},
		&shapes.Circle{uniqueID:58, X:10, Y:20, Radius:3, opacity:0},
		&shapes.Rectangle{uniqueID:79, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/

	shapesArr[0].(*shapes.Circle).SetOpacity(22)
	ListAll(shapesArr)
	/*
		&shapes.Circle{uniqueID:37, X:10, Y:4, Radius:5, opacity:22},
		&shapes.Rectangle{uniqueID:23, X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{uniqueID:24, X:10, Y:20, Radius:3, opacity:100},
		&shapes.Rectangle{uniqueID:82, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
	ListAll(shapesCopy) // opacity has not changed in this slice, because circle is a copy, that's exactly what we want.
	/*
		&shapes.Circle{uniqueID:77, X:10, Y:4, Radius:5, opacity:0},
		&shapes.Rectangle{uniqueID:59, X:1, Y:1, Width:10, Height:25, opacity:0},
		&shapes.Circle{uniqueID:58, X:10, Y:20, Radius:3, opacity:0},
		&shapes.Rectangle{uniqueID:79, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
}

func ListAll(arr []any) {
	for _, v := range arr {
		fmt.Printf("%#v, ", v)
	}
	fmt.Println()
}
