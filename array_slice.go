package main

import "fmt"

func arraySlice() {
	num := [6]int{1, 2, 3, 4, 5, 6}

	var slice []int = num[1:4] //a[low : high]
	fmt.Println(slice)

	//Slices are like references to arrays
	//A slice does not store any data, it just describes a section of an underlying array.

	//Changing the elements of a slice modifies the corresponding elements of its underlying array.

	//Other slices that share the same underlying array will see those changes.
	//When slicing, you may omit the high or low bounds to use their defaults instead.

	//The default is zero for the low bound and the length of the slice for the high bound.
}
