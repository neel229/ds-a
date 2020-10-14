package main

import (
	"fmt"
)

type Queue struct {
	intSlice []int
}

func main() {
	var q *Queue = new(Queue)
	q.Enqueue(1)
	q.Enqueue(2)
	fmt.Println(q)
	q.Dequeue()
	fmt.Println(q)
}

// Enqueue adds an element to the back of the queue.
func (q *Queue) Enqueue(i int) {
	q.intSlice = append(q.intSlice, i)
}

// Dequeue returns the first item from the queue.
// And removes that item from the queue.
func (q *Queue) Dequeue() (int, error) {
	ret := q.intSlice[0]
	q.intSlice = q.intSlice[1:len(q.intSlice)]
	return ret, nil
}
