package main

import (
	"fmt"
	"sync"

	"github.com/arthurshafikov/patterns/singleton/problem/queue"
)

var wg sync.WaitGroup

func main() {
	readerFunction()

	wg.Add(30)

	fillerFuntion()
	fillerFuntion()
	fillerFuntion()

	wg.Wait()

	readerFunction()
}

func readerFunction() {
	queue := queue.NewQueue()
	// queue.elems slice is always empty, because there are different instances of Queue inside this programm.

	elems := queue.Read()

	for _, elem := range elems {
		fmt.Printf("%v, ", elem)
	}
	fmt.Println()
}

func fillerFuntion() {
	queue := queue.NewQueue()

	for i := 0; i < 10; i++ {
		go func(i int) {
			queue.AddToQueue(i)
			wg.Done()
		}(i)
	}
}
