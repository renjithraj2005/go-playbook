package main

import "fmt"

func typeAssertions() {
	var i interface{} = "hello"

	//A type assertion provides access to an interface value's underlying concrete value.
	s := i.(string)
	fmt.Println(s)
	//This statement asserts that the interface value i holds the concrete type String and assigns the underlying String value to the variable s.

	//If i does not hold a String, the statement will trigger a panic.
	//To test whether an interface value holds a specific type, a type assertion can return two values:
	//the underlying value and a boolean value that reports whether the assertion succeeded.
	s, ok := i.(string)
	fmt.Println(s, ok)

	//If i holds a string, then t will be the underlying value and ok will be true.

	//If not, ok will be false and t will be the "" value of type String, and no panic occurs.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	//uncomment the below line to see what happens
	//f = i.(float64) // panic
	//fmt.Println(f)
}
