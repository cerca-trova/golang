package main

import "fmt"

// goroutines

func loop(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}
}

func main() {
	go loop(0)
	var input string
	fmt.Scanln(&input)
}
