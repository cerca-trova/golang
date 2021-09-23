package main

import "fmt"

// user defines a user in the program.
type user struct {
	name  string
	email string
}

// notify implements a method with a value receiver.
func (u user) notify() {
	fmt.Printf("Sending User email to %s<%s> \n", u.name, u.email)
}

// changeEmail implements a method with pointer receiver.
func (u *user) changeEmail(email string) {
	u.email = email
}

// main function as application entry point
func main() {
	// define a self-defined data structure
	bill := user{"Bill", "bill@email.com"}
	bill.notify()

	lisa := &user{"lisa", "lisa@email.com"}
	lisa.notify()

	bill.changeEmail("chenran@nawaa.com")
	bill.notify()

	lisa.changeEmail("xux@nawaa.com")
	lisa.notify()
}
