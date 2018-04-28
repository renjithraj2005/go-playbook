package main

import "fmt"

func emptyinterface() {
	var i interface{}
	describeempty(i)

	i = 42
	describeempty(i)

	i = "hello"
	describeempty(i)
}

func describeempty(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
