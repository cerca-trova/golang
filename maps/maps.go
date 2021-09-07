package main

import "fmt"

// How to define a map

func define_a_map() {
	// var m map[string]int, This sentence define a map but not initialize
	m := make(map[string]int) // Define and initialize
	m["first"] = 1

	fmt.Println("After initializtion, m['first'] is :", m["first"])

	// Map have some embedded method
	// 1. DELETE
	delete(m, "first") // After deletion, m would be {"first": 0}, 1-->0
	fmt.Println("After deletion, m['first'] would be :", m["first"])

	// Shorter way to create a map
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}

	v1, v2 := elements["H"]
	fmt.Println(v1, v2)

	if v1, v2 := elements["H"]; v2 {
		fmt.Println(v1, v2)
	}
}

func main() {
	define_a_map()
}
