package main

//An interface type is defined as a set of method signatures.

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func interfaceTest() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := VertexInterface{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//There is an error in the below code
	//Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex
	//	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type VertexInterface struct {
	X, Y float64
}

func (v *VertexInterface) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
