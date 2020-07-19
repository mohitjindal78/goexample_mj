//Sample program to show how to declate method and
//Go complier handles them
package main

import (
	"fmt"
)

//user defined trpe user
type user struct {
	name  string
	email string
}

//notify implements a method with value receiver
func (u user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

//changeEmail implements a method with pointer receiver
func (u *user) changeEmail(email string) {
	u.email = email
}

//main is the entry point for the
func main() {
	fmt.Println("Lets begin")

	// Values of type user can be used to call methods
	// declared with a value receiver.
	fmt.Println("Values of type user can be used to call methods declared with a value receiver.")
	dev := user{"Dev", "dev@gmail.com"}
	dev.notify()

	// Pointers of type user can also be used to call methods
	// declared with a value receiver.
	fmt.Println("Pointers of type user can also be used to call methods declared with a value receiver.")
	lisa := &user{"Lisa", "lisa@gmail.com"}
	lisa.notify()

	// Values of type user can also be used to call methods
	// declared with a pointer receiver.
	fmt.Println("Values of type user can also be used to call methods declared with a pointer receiver.")
	dev.changeEmail("dev1@gmail.com")
	dev.notify()

	// Pointers of type user can be used to call methods
	// declared with a pointer receiver.
	fmt.Println("Pointers of type user can be used to call methods declared with a pointer receiver.")
	lisa.changeEmail("lisa1@gmail.com")
	lisa.notify()
}
