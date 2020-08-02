// Copyright © 2020 Mohit Jindal
// License: https://
package main

import (
	"fmt"
	"os"
	"time"
)

//// Echo prints its command-line arguments.
func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)

	fmt.Printf("%.2fs Elapsed Time\n", time.Since(start).Seconds())
}
