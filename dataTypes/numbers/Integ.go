package main

import "fmt"

func integer_type() {
	var b byte = 255 // byte type, valua ranges from 0 - 255(2 ** 8 -1)
	fmt.Println(b)
}
func main() {
	integer_type()
	fmt.Println("everythings is ok?")
}
