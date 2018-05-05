package main

import "fmt"

func do(i interface{}) {
	//A type switch is a construct that permits several type assertions in series.

	//A type switch is like a regular switch statement, but the cases in a type switch specify types (not values),
	//and those values are compared against the type of the value held by the given interface value.
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

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

	//Type Switches
	do(21)
	do("hello")
	do(true)
}
