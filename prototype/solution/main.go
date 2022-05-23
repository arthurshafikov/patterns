package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/arthurshafikov/patterns/prototype/solution/shapes"
)

func main() {
	rand.Seed(time.Now().Unix())

	shapesArr := []shapes.Shape{
		shapes.NewCircle(10, 4, 5),
		shapes.NewRectangle(1, 1, 10, 25),
		shapes.NewCircle(10, 20, 3),
		shapes.NewRectangle(9, 2, 24, 12),
	}
	shapesArr[1].SetOpacity(50)
	shapesArr[2].SetOpacity(100)

	// we need to create a copy of every shape
	var shapesCopy []shapes.Shape

	for _, shape := range shapesArr {
		shapesCopy = append(shapesCopy, shape.Clone()) // pretty, isn't it?
	}

	ListAll(shapesArr)
	/*
		&shapes.Circle{uniqueID:178, X:10, Y:4, Radius:5, opacity:0},
		&shapes.Rectangle{uniqueID:280, X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{uniqueID:122, X:10, Y:20, Radius:3, opacity:100},
		&shapes.Rectangle{uniqueID:901, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
	ListAll(shapesCopy)
	/*
		&shapes.Circle{uniqueID:211, X:10, Y:4, Radius:5, opacity:0},
		&shapes.Rectangle{uniqueID:417, X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{uniqueID:0, X:10, Y:20, Radius:3, opacity:100},
		&shapes.Rectangle{uniqueID:213, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
	shapesArr[0].SetOpacity(22)
	ListAll(shapesArr)
	/*
		&shapes.Circle{uniqueID:178, X:10, Y:4, Radius:5, opacity:22},
		&shapes.Rectangle{uniqueID:280, X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{uniqueID:122, X:10, Y:20, Radius:3, opacity:100},
		&shapes.Rectangle{uniqueID:901, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
	ListAll(shapesCopy) // opacity has not changed in this slice, because circle is a copy, that's exactly what we want.
	/*
		&shapes.Circle{uniqueID:211, X:10, Y:4, Radius:5, opacity:0},
		&shapes.Rectangle{uniqueID:417, X:1, Y:1, Width:10, Height:25, opacity:50},
		&shapes.Circle{uniqueID:0, X:10, Y:20, Radius:3, opacity:100},
		&shapes.Rectangle{uniqueID:213, X:9, Y:2, Width:24, Height:12, opacity:0}
	*/
}

func ListAll(arr []shapes.Shape) {
	for _, v := range arr {
		fmt.Printf("%#v, ", v)
	}
	fmt.Println()
}
