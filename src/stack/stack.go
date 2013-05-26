package stack

import (
	"fmt"
	"os"
	"math/rand"
	"time"
)

const (
 	STACK_SIZE = 10
	MAX_INT = 10
)

// stack of integers (LIFO)
type Stack struct {
	items []int
}

func (s *Stack) push(item int) {
	s.items = append(s.items, item)
	fmt.Println("added: ", item)
}

// remove most recently added item
func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Printf("stack has no items to pop!\n")
		os.Exit(1)
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

func create() *Stack {
	s := new(Stack)
	for i := 0; i < STACK_SIZE; i++ {
		s.push(rand.Intn(MAX_INT + 1))
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
	fmt.Println("emptying stack:")
	for {
		fmt.Printf("popped from stack: %v\n", s.pop())
	}
}
