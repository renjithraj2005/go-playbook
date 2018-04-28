package main

import (
	"fmt"
	"math"
)

type vertexDemo struct {
	X, Y float64
}

//Go does not have classes. However, you can define methods on types.
func (v vertexDemo) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v vertexDemo) Sum() float64 {
	return v.X + v.Y
}

//Pointer Receiver

//You can declare methods with pointer receivers.

//This means the receiver type has the literal syntax *T for some type T.
//(Also, T cannot itself be a pointer such as *int.)

//For example, the Scale method here is defined on *vertexDemo.
func (v *vertexDemo) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

//Try removing the * from the declaration of the Scale function on line 29 and
//observe how the program's behavior changes.

//With a value receiver, the Scale method operates on a copy of the original Vertex value.
//(This is the same behavior as for any other function argument.)
//The Scale method must have a pointer receiver to change the Vertex value declared in the main function.

func method() {
	v := vertexDemo{3, 4}
	fmt.Println(v.Abs())
	v.Scale(2)
	fmt.Println(v.Sum())

}

//A method is a function with a special receiver argument.

//The receiver appears in its own argument list between the func keyword and the method name.

//In this example, the Abs method has a receiver of type Vertex named v.

// You can declare a method on non-struct types, too

/*
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
*/
