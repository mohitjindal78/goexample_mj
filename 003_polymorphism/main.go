//Sample program to show how to declate method and
//Go complier handles them
package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

//user defined type user
type user struct {
	name  string
	email string
}

//notify implements notifier interface with pointer receiver
func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n",
		u.name,
		u.email)
}

//user defined type admin
type admin struct {
	name  string
	email string
}

//notify implements notifier interface with pointer receiver
func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n",
		a.name,
		a.email)
}

//main is the entry point for the
func main() {

	//Create a user value and send notification
	dev := user{"Dev", "dev@gmail.com"}
	sendNotification(&dev)

	//Create an admin value and send notification
	riju := admin{"Riju", "riju@gmail.com"}
	sendNotification(&riju)
}

// sendNotification accepts values that implement the notifier
// interface and sends notifications.
func sendNotification(n notifier) {
	n.notify()
}
