package main

import "fmt"

/**
* A struct literal denotes a newly allocated struct value by listing the values of its fields.
* You can list just a subset of fields by using the Name: syntax.
* (And the order of named fields is irrelevant.)
* The special prefix & returns a pointer to the struct value.
 */

var (
	v1 = vertex{1, 2}  // has type Vertex
	v2 = vertex{X: 1}  // Y:0 is implicit
	v3 = vertex{}      // X:0 and Y:0
	p  = &vertex{1, 2} // has type *Vertex
)

func structLiterals() {
	fmt.Println(v1, p, v2, v3)
}
