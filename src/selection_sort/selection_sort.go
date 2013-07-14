package selection_sort

import (
	"fmt"
	"math/rand"
	"time"
)

var unsorted_array []int
var unsorted_array_exists = false
var array []int
var size int = 5000

func find_lowest_element() int {
	lowest_value := -1
	lowest_index := -1
	for i, v := range array {
		if v < lowest_value || lowest_value == -1 {
			lowest_value = v
			lowest_index = i
		}
	}
	return lowest_index
}

func array_print() {
	for i, v := range array {
		fmt.Printf("[%d] %d\n", i, v)
	}
}

func create_array() {
	if !unsorted_array_exists {
		unsorted_array = make([]int, size)
		for i := 0; i < size; i++ {
			r := rand.Intn(1000)
			unsorted_array[i] = r
		}
		unsorted_array_exists = true
	}
	array = make([]int, size)
	for i := 0; i < size; i++ {
		array[i] = unsorted_array[i]
	}
}

func verify_sort() {
	for i := 0; i < len(array) - 2; i++ {
		if array[i] > array[i + 1] {
			fmt.Printf("array not sorted! [%d]: %d, [%d]: %d\n", i, array[i], i+1, array[i+1])
			panic("SORT ERROR")
		}
	}
}

func Run() {
	num_cycles := 10
	size = 20000
	recorded_times := make([]int64, num_cycles)
	cycle := 0
	for ;cycle < num_cycles; cycle++ {
		start := time.Now()
		create_array()
		for j := 0; j < size - 1; j++ {
			iMin := j
			for i := j + 1; i < size; i++ {
				if array[i] < array[iMin] {
					iMin = i
				}
			}
			if iMin != j {
				tmp := array[iMin]
				array[iMin] = array[j]
				array[j] = tmp
			}
		}
		elapsed := time.Since(start)
		recorded_times[cycle] = elapsed.Nanoseconds() / 1000 // microseconds
		verify_sort()
		fmt.Printf("elapsed time: %v\n", elapsed)
	}
	var total int64 = 0
	for _, v := range recorded_times {
		total += v
	}
	average := float64(total / int64(num_cycles)) / 1000 // milliseconds
	fmt.Printf("average elapsed time: %v ms\n", average)


}
