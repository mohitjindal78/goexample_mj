// Package treesort provides insertion sort using an unbalanced binary tree.
package main

import (
	"fmt"
	"sort"
)

type tree struct {
	value int
	left  *tree
	right *tree
}

func main() {
	data := make([]int, 50)
	data = []int{10, 1, 21}
	/*for i := range data {
		data[i] = rand.Int() % 50
		fmt.Printf("%d\t%d\n", i, data[i])
	}*/
	sortt(data)
	if !sort.IntsAreSorted(data) {
		fmt.Printf("not sorted: %v\n", data)
	} else {
		fmt.Printf("sorted: %v\n", data)
	}
}

func sortt(data []int) {
	var root *tree
	for _, v := range data {
		root = add(root, v)
	}
	appendValues(data[:0], root)
}

func add(t *tree, value int) *tree {
	if t == nil {
		// Equivalent to return &tree{value: value}.
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func appendValues(values []int, t *tree) []int {
	//fmt.Println(values)
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}
