package main

import "fmt"

func pointer() {
	i, j := 1, 2
	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)
	p = &j
	fmt.Println(*p)

}
