package main

import "fmt"

func zero(xPtr *int) { // xPtr refer to memory address, like 0x0000001c
	*xPtr = 0 // *xPtr refer to the value in which specific memory address stored, Here is 0
}

func main() {
	x := 5
	zero(&x)
	fmt.Println(x)
}
