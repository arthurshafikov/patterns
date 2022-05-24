package queue

import (
	"sync"
)

// struct is private.
type queue struct {
	mu sync.Mutex

	elems []int
}

var queueInstance *queue

func GetInstance() *queue {
	if queueInstance == nil {
		queueInstance = newQueue()
	}

	return queueInstance
}

// constructor is private.
func newQueue() *queue {
	return &queue{}
}

func (d *queue) AddToQueue(newElem int) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.elems = append(d.elems, newElem)

	return true
}

func (d *queue) Read() []int {
	return d.elems
}
