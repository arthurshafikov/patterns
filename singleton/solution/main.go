package main

import (
	"fmt"
	"sync"

	"github.com/arthurshafikov/patterns/singleton/solution/queue"
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
	// 9, 0, 1, 2, 3, 4, 5, 6, 7, 9, 0, 1, 2, 3, 4, 5, 6, 7, 8, 2, 8, 9, 0, 1, 5, 3, 4, 6, 7, 8,
}

func readerFunction() {
	queue := queue.GetInstance()

	elems := queue.Read()

	for _, elem := range elems {
		fmt.Printf("%v, ", elem)
	}
	fmt.Println()
}

func fillerFuntion() {
	queue := queue.GetInstance()

	for i := 0; i < 10; i++ {
		go func(i int) {
			queue.AddToQueue(i)
			wg.Done()
		}(i)
	}
}
