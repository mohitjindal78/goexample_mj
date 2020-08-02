// Dup2 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
//The version of dup read the entire input into memory in one big gulp, split it into lines all at
//once, then process the lines.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)

	for _, filename := range os.Args[1:] {
		//ReadFile returns a byte slice that must be converted into a string so it can be split by
		//strings.Split
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}

		//Note: Ignoring potential errors from input.Err()
		for line, n := range counts {
			if n > 1 {
				fmt.Printf("%s\t%d\t%s\n", filename, n, line)
			}
		}

		//Clearing earlier map counts
		for k := range counts {
			delete(counts, k)
		}
	}
}