package main

import "fmt"

// to run call this method on main.go
func makeSample() {

	// Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.
	//The make function allocates a zeroed array and returns a slice that refers to that array:
	a := make([]int, 5)
	printSliceMake("a", a)

	b := make([]int, 0, 5)
	printSliceMake("b", b)

	c := b[:2]
	printSliceMake("c", c)

	d := c[2:5]
	printSliceMake("d", d)
}

func printSliceMake(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}
