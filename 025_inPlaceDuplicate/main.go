package main

import "fmt"

func main() {
	s := []string{"a", "b", "b", "", "a", "a"}
	fmt.Println(s)
	s = nonDuplicate(s)
	fmt.Println(s)
}

func nonDuplicate(s []string) []string {
	for i := range s {
		if i < len(s)-1 {
			if s[i] == s[i+1] {
				copy(s[i:], s[i+1:])
				s = s[:len(s)-1]
			}
		}
		fmt.Println(i, s, len(s))
	}
	return s
}
