package main

import "fmt"

func arraysample() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	//An array's length is part of its type, so arrays cannot be resized.. uncomment the below code to see what happen
	//a[2] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	//array slices

	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4] //a[low : high]
	fmt.Println(s)

}
