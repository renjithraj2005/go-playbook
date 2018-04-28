package main

import "fmt"

//https://blog.golang.org/defer-panic-and-recover
func defersample() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
