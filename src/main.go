package main

import (
	"bag"
	"fmt"
	"os"
	"queue"
	"stack"
)

// TODO make all data structures deal with a struct instead of ints
// i.e.
/*type Item struct {
	cargo int
}
func (i *Item) getCargo() int {...
*/

type module interface {
	Run()
}

var (
	moduleName string
	validNames = []string{
		"bag",
		"queue",
		"stack",
	}
)

func isModuleNameValid(name string) bool {
	for _, validName := range validNames {
		if name == validName {
			return true
		}
	}
	return false
}

func printValidModuleNames() {
	for _, validName := range validNames {
		fmt.Println(validName)
	}
}

func init() {
	if len(os.Args) < 2 {
		fmt.Println("please supply module name as first argument.")
		os.Exit(1)
	}
	moduleName = os.Args[1]
	if !isModuleNameValid(moduleName) {
		fmt.Printf("modulename '%s' is not valid, valid names are:\n", moduleName)
		printValidModuleNames()
		os.Exit(1)
	}
}

func runModule(name string) {
	switch name {
	case "bag":
		bag.Run()
	case "queue":
		queue.Run()
	case "stack":
		stack.Run()
	}
}

func main() {
	fmt.Println("trying to run module: " + moduleName)
	runModule(moduleName)
}
