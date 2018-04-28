package main

import (
	"fmt"
	"math"
)

type Interface interface {
	M()
}

type Struct struct {
	S string
}

func (t *Struct) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

type Float float32

func (t Float) M() {
	fmt.Println(t)
}

func interfaceValues() {
	var i Interface
	i = &Struct{"Hello"}
	describe(i)
	i.M()

	i = Float(math.Pi)
	describe(i)
	i.M()

	// Interface values with nil underlying values

	var j Interface
	var k *Struct
	j = k
	describe(j)
	j.M()

}

func describe(i Interface) {
	fmt.Printf("(%v, %T)\n", i, i)
}
