// Printints demonstrates the use of bytes.Buffer to format a string.
package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Println(intsToString([]int{1, 2, 3, 4})) //"[1,2,3,4]"
}

func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteString("[")
	//buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}