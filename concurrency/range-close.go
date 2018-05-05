package concurrency

//A sender can close a channel to indicate that no more values will be sent.
//Receivers can test whether a channel has been closed by assigning a second parameter
//to the receive expression: after

//v, ok := <-ch
//ok is false if there are no more values to receive and the channel is closed.

//Only the sender should close a channel, never the receiver.
//Sending on a closed channel will cause a panic.

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func ChannelClose() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
