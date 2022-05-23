package queue

import "sync"

type Queue struct {
	mu sync.Mutex

	elems []int
}

func NewQueue() *Queue {
	return &Queue{}
}

func (d *Queue) AddToQueue(newElem int) bool {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.elems = append(d.elems, newElem)

	return true
}

func (d *Queue) Read() []int {
	return d.elems
}
