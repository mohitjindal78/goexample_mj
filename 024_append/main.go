// Append illustrates the behavior of the built-in append function.
package main

import "fmt"

func appendSlice(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		//There is room to grow extend slice
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	copy(z[len(x):], y)
	return z
}

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		//There is room to grow extend slice
		z = x[:zlen]
	} else {
		// There is insufficient space.  Allocate a new array.
		// Grow by doubling, for amortized linear complexity.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x) // a built-in function; see text
	}
	z[len(x)] = y
	return z
}

func main() {
	var x, y []int
	for i := 0; i <= 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d \tcap=%d\t\t%v\n", i, cap(y), y)
		x = y
	}
	y = []int{11, 12, 13}
	x = appendSlice(x, y...)
	fmt.Println(x)
	fmt.Println(appendSlice(x, 14, 15, 16))
}
