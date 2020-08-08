// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//      $ go build gopl.io/ch3/comma
//      $ ./comma 1 12 123 1234 1234567890
//      1
//      12
//      123
//      1,234
//      1,234,567,890
//
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		var sf, ss string //string first part, string second part
		s := os.Args[i]
		if dot := strings.LastIndex(s, "."); dot >= 0 {
			sf = s[:dot]
			ss = s[dot:]
			fmt.Printf(" %s\n", comma(sf)+ss)
		} else {
			fmt.Printf(" %s\n", comma(s))
		}
	}
}

func comma(s string) string {
	/*
		n := len(s)
		//if n <= 3 {
		//	return s
		//}
		//return comma(s[:n-3]) + "," + s[n-3:]
	*/
	var buf bytes.Buffer
	n := len(s)
	//for i := 0; i < n; i++ { //this is normal for loop
	for i, v := range s { //range return rune of string
		if (n-i)%3 == 0 {
			if i != 0 {
				buf.WriteByte(',')
			}
		}
		fmt.Fprintf(&buf, "%c", v)
		//buf.WriteString(s[i : i+1]) //this is with normal for loop
	}
	return buf.String()
}
