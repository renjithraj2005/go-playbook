package main

import "fmt"

//https://blog.golang.org/defer-panic-and-recover
func defersample() {

	//A defer statement defers the execution of a function until the surrounding function returns.

	//The deferred call's arguments are evaluated immediately,
	//but the function call is not executed until the surrounding function returns.
	defer fmt.Println("world")

	fmt.Println("hello")
}
