package main

import "fmt"

type vertex struct {
	X int
	Y int
}

func structsample() {
	fmt.Println(vertex{1, 2})
	//Struct Fields are accessed via dot
	v := vertex{1, 2}
	fmt.Println(v.X)

	//Pointers to structs

	p := &v
	//To access the field X of a struct when we have the struct pointer p we could write (*p).X.
	//However, that notation is cumbersome, so the language permits us instead
	//to write just p.X, without the explicit dereference.
	fmt.Println(p.X)

	//Struct Literals
	//v1 = Vertex{1, 2}  // has type Vertex
	//v2 = Vertex{X: 1}  // Y:0 is implicit
	//v3 = Vertex{}      // X:0 and Y:0
	//p  = &Vertex{1, 2} // has type *Vertex
}
