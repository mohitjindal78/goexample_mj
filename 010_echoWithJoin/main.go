// Copyright © 2020 Mohit Jindal
// License: https://
package main

import (
	"fmt"
	"os"
	"strings"
)

//// Echo prints its command-line arguments.
func main() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
