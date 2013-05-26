package stack

import (
	"fmt"
)

const (
 	STACK_SIZE = 20
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

func (s *Stack) pop() {

}
