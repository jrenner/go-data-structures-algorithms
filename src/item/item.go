package item

import (
	"fmt"
)

// Integer Item
type ItemInt struct {
	Cargo int
}

func (item ItemInt) String() string {
	return fmt.Sprintf("%v", item.Cargo)
}

// String item
type ItemString struct {
	Cargo int
}

func (item ItemString) String() string {
	return fmt.Sprintf("%v", item.Cargo)
}
