package main

//import "fmt"
//import "math" // This can be optimised as below
import (
	"fmt"
	"go-playbook/function"
	"math"
	"math/rand"
)

var c bool // Variables

// Comment in Golang Ref https://golang.org/doc/effective_go.html#commentary
func add(a int, b int) int {
	return a + b
}

// go syntax declaration https://blog.golang.org/gos-declaration-syntax
func addOtherways(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Random number", rand.Intn(10), math.Sqrt(2))
	fmt.Println("exported name", math.Pi)
	fmt.Println("function sample", add(1, 3))
	fmt.Println("function sample another way", addOtherways(1, 3))
	a, b := function.Swap(1, 2)
	fmt.Println(a, b)
	fmt.Println(function.Split(25))
	fmt.Println("variables", c)

	//Variables with initializers
	var python, java = false, "no!"
	fmt.Println(python, java)

	//short variable declaration
	k := "3"
	fmt.Println(k)
	fmt.Printf("Type: %T Value: %v\n", k, k)

	//Variables declared without an explicit initial value are given their zero value.

	// Zero value of int is 0, bool is false and String is ""

	// Type Conversion : The expression T(v) converts the value v to the type T.
	var i = 10
	var typeConversionTest = uint(i)
	fmt.Printf("Type : %T Value %v\n", typeConversionTest, typeConversionTest)

	//Try the below commented code to know what happen ?
	//fmt.Printf(function.Pi)
	fmt.Printf("Constant %v\n", function.Pi)
	sum := 1

	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println("---------------------------")
	//defersample()
	//switchSample()

	//pointer()

	// Go return error even if the variable is unused
	//ex: j declared and not used

	//structsample()
	//arraySlice()
	//sliceOfSlice()
	//rangeSample()
	//typeAssertions()
	//stringers()
	//errorExample()
	//reader()
	//concurrency.Routine()
	//concurrency.Select()
	//orm.InitORM()
	//fetch() // a sample http client implementation in go
}
