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
	// If the length is given as ..., Go will identify the length of the array based on the number
	// of elements that are initialized
	arr3 := [...]float64{3.14, 1.71828, 1.213}
	fmt.Println(arr3)
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

	s2 := []byte{'a', 'o', 'l'}
	// For loop read s1
	fmt.Println("slice")
	for i := 0; i < len(s2); i++ {
		fmt.Println(s2[i])
	}
}

func length_and_capacity_of_slice() []int {
	// define a empty slice, length = 0 capacity = 0
	slice := []int{}
	//define a nil slice
	var slice_nil []int

	orignal_slice := []int{10, 20, 30, 40, 50}

	newSlice := orignal_slice[1:3]
	appendedSlice := append(newSlice, 60)            // underlying array changed to [10,20,30,60,50]
	outCapacitySlice := append(newSlice, 60, 70, 80) // if new slice's length overpass its underlying array, a new underlying array will be allocated and its capacity will be doubled
	// In this case, capacity increased from 4 to 8
	fmt.Println("orignal slice is :", orignal_slice)
	fmt.Printf("orignal slice's length is %d capacity is %d \n", len(orignal_slice), cap(orignal_slice))
	fmt.Println("new slice is :", newSlice)
	fmt.Printf("new slice's length is %d capacity is %d \n", len(newSlice), cap(newSlice))
	fmt.Println("appended slice is:", appendedSlice)
	fmt.Printf("appended slice's length is %d capacity is %d \n", len(appendedSlice), cap(appendedSlice))
	fmt.Println("out-capacity slice is:", outCapacitySlice)
	fmt.Printf("out-capacity slice's length is %d capacity is %d \n", len(outCapacitySlice), cap(outCapacitySlice))
	fmt.Println(slice)
	return slice_nil

}

func passing_slice_between_functions(slice []int) []int {
	// define a slice with 1 million intergers.
	return slice // return its pointer,length and capacity, not its huge underlying array.

}
func main() {
	slice := make([]int, 1e6)
	define_an_array()
	define_an_slice()
	length_and_capacity_of_slice()
	passing_slice_between_functions(slice)
}
