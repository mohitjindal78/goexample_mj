// Copyright © 2020 Mohit Jindal
// License: https://
package main

import (
	"fmt"
	"os"
)

// Echo prints its command-line arguments.
func main() {
	var s, sep string
	fmt.Println(os.Args[0], os.Args[1])

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
