package queue

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const QUEUE_SIZE = 10;
const MAX_INT = 10;

// queue of integers (FIFO)
type Queue struct {
	items []int
}

func (q *Queue) size() int {
	return len(q.items)
}

func (q *Queue) enqueue(item int) {
	q.items = append(q.items, item)
	//fmt.Println("added:", item)
}

// remove the least recently added item
func (q *Queue) dequeue() int {
	if (q.isEmpty()) {
		fmt.Println("queue has no items to dequeue!")
		os.Exit(1)
	}
	removal := q.items[0]
	q.items = q.items[1:]
	return removal
}

func (q *Queue) isEmpty() bool {
	if (q.size() == 0) {
		return true
	}
	return false
}

func create() *Queue {
	q := new(Queue)
	for i := 0; i < QUEUE_SIZE; i++ {
		q.enqueue(rand.Intn(MAX_INT + 1))
	}
	return q
}

func Run() {
	rand.Seed(time.Now().UnixNano())
	q := create()
	fmt.Printf("created queue of size: %d\n", q.size())
	fmt.Println("removing items from queue, oldest first (FIFO)")
	for _, item := range q.items {
		fmt.Printf("%v\n", item)
	}

}
