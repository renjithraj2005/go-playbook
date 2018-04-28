package main

import "fmt"

// to run call this method on main.go
func sliceLiterals() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	//A slice literal is like an array literal without the length.
	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)

	//Slice defaults
	//When slicing, you may omit the high or low bounds to use their defaults instead.

	//The default is zero for the low bound and the length of the slice for the high bound.
	//For the array var a [10]int

	//these slice expressions are equivalent:

	//a[0:10]
	//a[:10]
	//a[0:]
	//a[:]
}
