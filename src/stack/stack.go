package stack

import (
	"fmt"
	"math/rand"
	"time"
	. "item"
)

const (
 	STACK_SIZE = 10
	MAX_INT = 10
)

// stack of integers (LIFO)
type Stack struct {
	items []*ItemInt
}

func (s *Stack) push(item *ItemInt) {
	s.items = append(s.items, item)
}

// remove most recently added item
func (s *Stack) pop() *ItemInt {
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

func create() *Stack {
	s := new(Stack)
	for i := 0; i < STACK_SIZE; i++ {
		item := &ItemInt{rand.Intn(MAX_INT + 1)}
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
	a := &ItemInt{1}
	b := &ItemInt{2}
	c := &ItemInt{3}
	fmt.Printf("adding %v, %v and %v to the stack\n", a, b, c)
	s.push(a)
	s.push(b)
	s.push(c)
	fmt.Println("emptying stack:")
	for {
		popped := s.pop()
		if (popped == nil) {
			fmt.Println("stack has no items to pop!")
			break;
		} else {
			fmt.Printf("popped from stack: %v\n", s.pop())
		}
	}
}


// Djikstra's Two-Stack Algorithm for Expression Evaluation
/*func djikstraEvaluation() {
	fmt.Println("running Djikstra's Two-Stack Algorithm for Expression Evaluation")
	testCase := "5 - (4 - 2)"

}*/
