
package linked_list

import (
	"fmt"
	"math/rand"
)

const (
	LIST_SIZE = 5
	MAX_INT = 100
)

type Node struct {
	Item interface{}
	Next *Node
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Item)
}

func iterateLinkedList(start *Node) *Node{
	i := 0
	current := start
	for {
		fmt.Printf("[%d]: %v, next: %v", i, current, current.Next)
		if i > 0 && current == start {
			fmt.Printf("    ----    looped around to start node, stopping loop.\n")
			break
		}
		if current.Next == nil {
			fmt.Println("    ----    no more nodes in linked list")
			break
		}
		i++
		current = current.Next
		fmt.Println()
	}
	return current // return our stopping point
}

func Run() {
	first := create()
	current := first
	fmt.Println("traversing linked list starting from first node:")
	current = iterateLinkedList(first)
	fmt.Println("making list circular and looping twice:")
	// next of the last node is now the first node
	current.Next = first
	for j := 0; j < 2; j++ {
		iterateLinkedList(first)
	}


}

func createNode() *Node {
	node := &Node{Item: rand.Intn(MAX_INT + 1), Next: nil}
	return node
}

// returns first node
func create() *Node {
	var current *Node
	first := createNode()
	current = first
	for i := 0; i < LIST_SIZE; i++ {
		current.Next = createNode()
		current = current.Next
	}
	return first
}

