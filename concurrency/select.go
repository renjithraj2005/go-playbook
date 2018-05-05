package concurrency

import "fmt"

//The select statement lets a goroutine wait on multiple communication operations.

//A select blocks until one of its cases can run, then it executes that case.
//It chooses one at random if multiple are ready.
//The default case in a select is run if no other case is ready.

//Use a default case to try a send or receive without blocking:
func fibonaccinew(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func Select() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonaccinew(c, quit)
}
