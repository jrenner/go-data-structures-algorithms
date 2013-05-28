package queue

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	QUEUE_SIZE = 5;
	MAX_INT = 100;
)

type Queue struct {
	items []interface{}
}


func (q *Queue) size() int {
	return len(q.items)
}

func (q *Queue) enqueue(item interface{}) {
	q.items = append(q.items, item)
}

// remove the least recently added item
func (q *Queue) dequeue() interface{} {
	if (q.isEmpty()) {
		fmt.Println("queue has no items to dequeue!")
		os.Exit(1)
	}
	removal := q.items[0]
	q.items = q.items[1:]
	return &removal
}

func (q *Queue) isEmpty() bool {
	if (q.size() == 0) {
		return true
	}
	return false
}

func (q *Queue) String() string {
	out := "queue contents:"
	for i, item := range q.items {
		out += fmt.Sprintf("\n[%d]: %3v", i, item)
	}
	return out
}

func create() *Queue {
	q := new(Queue)
	for i := 0; i < QUEUE_SIZE; i++ {
		item :=rand.Intn(MAX_INT + 1)
		q.enqueue(item)
	}
	return q
}

func Run() {
	rand.Seed(time.Now().UnixNano())
	q := create()
	fmt.Println(q)
	fmt.Println("removing items from queue, oldest first (FIFO)")
	for _, item := range q.items {
		fmt.Printf("deqeued: %v\n", item)
	}

}
