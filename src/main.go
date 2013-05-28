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
	runAll     = false
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
	if len(os.Args) < 2 || os.Args[1] == "all" {
		runAll = true
	} else {
		moduleName = os.Args[1]
		if !isModuleNameValid(moduleName) {
			fmt.Printf("modulename '%s' is not valid, valid names are:\n", moduleName)
			printValidModuleNames()
			os.Exit(0)
		}
	}
}

func runModule(name string) {
	fmt.Println("running module: " + name)
	printSeparator()
	switch name {
	case "bag":
		bag.Run()
	case "queue":
		queue.Run()
	case "stack":
		stack.Run()
	}
}

func printSeparator() {
	fmt.Println("------------------------------")
}

// for running modules one after another
func runSeriesModule(name string) {
	runModule(name)
	fmt.Print("press any key to continue with next module: ")
	inputBuff := make([]byte, 256)
	for {
		n, _ := os.Stdin.Read(inputBuff)
		if n > 0 {
			break
			fmt.Print("\n")
		}
	}
	printSeparator()
}

func runAllModules() {
	fmt.Println("running all modules:")
	printSeparator()
	runSeriesModule("bag")
	runSeriesModule("queue")
	runModule("stack")
}

func main() {
	if runAll {
		runAllModules()
	} else {
		runModule(moduleName)
	}
}
