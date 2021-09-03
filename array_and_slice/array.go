package main

import "fmt"

func define_an_array() {
	var arr [5]int // elements' value not pre-defined
	arr[4] = 100
	fmt.Println(arr) // result would be [0 0 0 0 100], if elements' value not pre-defined, default is 0
	// arr2 will construct an array with pre-defined elements' value

	arr2 := [4]string{"chen", "cai", "hu", "fang"}
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}
}

func main() {
	define_an_array()
}
