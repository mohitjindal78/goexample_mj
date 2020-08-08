// Comma prints its argument numbers with a comma at each power of 1000.
//
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args)-1 == 2 {
		ok := isAnagram(os.Args[1], os.Args[2])
		if ok {
			fmt.Printf("%s and %s are Anagram of each other\n", os.Args[1], os.Args[2])
		} else {
			fmt.Printf("%s and %s are not Anagram of each other\n", os.Args[1], os.Args[2])
		}
	} else {
		fmt.Println("Need only two arguments")
	}
}

func isAnagram(s1 string, s2 string) bool {
	ok := false
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			fmt.Printf("%c %[1]v %c %[2]v\n", v1, v2)
			if v1 == v2 {
				ok = true
				break
			} else {
				ok = false
			}
		}
		if !ok {
			return ok
		}
	}
	return ok
}
