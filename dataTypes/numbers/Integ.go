package main

import "fmt"

func integer_type() {
	var b byte = 255                                                 // byte type, valua ranges from 0 - 255 (2 ** 8 -1)
	var i int8 = 0                                                   // equivalent to byte type, In fact, 8 represent 8 bits
	i2 := 65535                                                      // shorthand assignment,Go will auto detect value type
	const unchangeable string = "This is a value you can not change" //constant value
	unchangeable = "try change it "
	fmt.Println(b, i, i2, unchangeable)
}
func main() {
	integer_type()
	fmt.Println("everythings is ok?")
}
