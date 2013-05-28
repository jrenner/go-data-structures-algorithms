package bag

import (
	"fmt"
	"math/rand"
	"time"
	"math"
)

const BAG_SIZE = 10
const MAX_INT = 100

// bag of integers
type Bag struct {
	items []interface{}
}

func (b *Bag) add(item interface{}) {
	b.items = append(b.items, item)
}

func (b *Bag) size() int {
	return len(b.items)
}

func (b *Bag) isEmpty() bool {
	if b.size() == 0 {
		return true
	}
	return false
}

func (b *Bag) String() string {
	out := "bag contents:"
	for i, item := range b.items {
		out += fmt.Sprintf("\n[%d]: %3v", i, item)
	}
	return out
}

func Run() {
	var sum float64
	var mean float64
	var stdDev float64
	rand.Seed(time.Now().UnixNano())

	bag := create()
	fmt.Println(bag)
	for _, item := range bag.items {
		sum += float64(item.(int))
	}
	mean = sum / float64(bag.size())

	sum = 0
	for _, item := range bag.items {
		sum += math.Pow(float64(item.(int)) - mean, 2)
	}
	stdDev = math.Sqrt(float64(sum) / float64(bag.size() - 1))

	fmt.Printf("mean: %g\n", mean)
	fmt.Printf("std dev: %g\n", stdDev)
}

func create() *Bag {
	bag := new(Bag)
	for i := 0; i < BAG_SIZE; i++ {
		item := rand.Intn(MAX_INT + 1)
		bag.add(item)
	}
	return bag
}




