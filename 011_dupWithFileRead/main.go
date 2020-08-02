// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
//The versions of dup above operate in a “streaming” mode in which input is read and broken into
//lines as needed, so in principle these programs can handle an arbitrary amount of input.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	//Note: Ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	//Bufio scanner reads input and breaks it into lines or words
	input := bufio.NewScanner(f)
	//Each call to input.Scan() reads the next line and removes the newline character from the end;
	//the result can be retrieved by calling input.Text().
	//The Scan function returns true if there is a line and false when there is no more input
	for input.Scan() {
		counts[input.Text()]++
	}
}
