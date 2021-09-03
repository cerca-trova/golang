package main

import "fmt"

func loop_for() {
	start := 0
	for i := start; i < 15; i++ {
		fmt.Println(i)
	}
}

func if_and_else() {
	start := 0
	for i := start; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i, "even")
		} else if i%2 != 0 {
			fmt.Println(i, "odd")
		}
	}
}
func main() {
	loop_for()
	if_and_else()
}
