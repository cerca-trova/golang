package main

import "fmt"

func define_an_array() {
	var arr [5]int // elements' value not pre-defined
	arr[4] = 100
	fmt.Println(arr) // result would be [0 0 0 0 100], if elements' value are not pre-defined, It would be 0
	
	// arr2 will construct an array with pre-defined elements' value
	arr2 := [4]string{"chen", "cai", "hu", "fang"}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
}

// Define a slice
func define_an_slice() {
	arr3 := [5]float64{1, 2, 3, 4, 6}
	s := make([]float64, 5)

	// Another wat in defing a slice
	s1 := arr3[0:] // equivalent to arr3[:] or arr3[0:len(arr3)]

	// slice have two basic ways of appending an new element into a existed slice.
	// they are COPY
	s = append(s, 3.14, 1.71828)

	// and APPEND, Besides, Array don't have such attribute, they are fixed and unchangeable
	s1Copy := make([]float64, len(s1))
	copy(s1Copy, s1)

	// For loop read s1
	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func main() {
	// define_an_array()
	define_an_slice()
}
