package main

import (
	"fmt"
)

func main() {
	xs := []float64{3.2, 3.6, 7.7, 8.6, 0.9}
	defer fmt.Println(average(xs)) // delay to execute this function until next one was finished
	fmt.Println(multi_args_function(1, 2, 3, 4))
}

func average(xs []float64) float64 {
	// panic(" Not Implemented")

	total := 0.0
	for _, v := range xs {
		total += v
	}

	return total / float64(len(xs))
}

func multi_args_function(ma ...int) int {
	sum := 0

	for _, v := range ma {
		sum += v
	}
	return sum
}
