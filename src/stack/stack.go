package stack

import (
	"fmt"
	"math/rand"
	"time"
)

const (
 	STACK_SIZE = 5
	MAX_INT = 100
)

// stack of integers (LIFO)
type Stack struct {
	items []interface{}
}

func (s *Stack) push(item interface{}) {
	s.items = append(s.items, item)
}

// remove most recently added item
func (s *Stack) pop() interface{} {
	if s.isEmpty() {
	    return nil
	}
	removal := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items) - 1]
	return removal
}

func (s *Stack) isEmpty() bool {
	if (s.size() == 0) {
		return true
	}
	return false
}

func (s *Stack) size() int {
	return len(s.items)
}

func (s *Stack) String() string {
	out := "stack contents:"
	for i, item := range s.items {
		out += fmt.Sprintf("\n[%d]: %3v", i, item)
	}
	return out
}

func create() *Stack {
	s := new(Stack)
	for i := 0; i < STACK_SIZE; i++ {
		item := rand.Intn(MAX_INT + 1)
		s.push(item)
	}
	return s
}

func Run() {
	rand.Seed(time.Now().UnixNano())
	s := create()
	fmt.Println("size of stack:", s.size())
	halfway := s.size() / 2
	for i := 0; i < halfway; i++ {
		item := s.pop()
		fmt.Printf("popped from stack: %v\n", item)
	}
	a := 1
	b := 2
	c := 3
	fmt.Printf("adding %v, %v and %v to the stack\n", a, b, c)
	s.push(a)
	s.push(b)
	s.push(c)

	fmt.Println(s)
	fmt.Println("emptying stack:")
	for {
		popped := s.pop()
		if (popped == nil) {
			fmt.Println("stack has no items to pop!")
			break;
		} else {
			fmt.Printf("popped from stack: %v\n", popped)
		}
	}
}
