package concurrency

import (
	"fmt"
	"time"
)

//A goroutine is a lightweight thread managed by the Go runtime.
//for ex: go f(x, y, z) starts a new goroutine running f(x, y, z)
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// Goroutines run in the same address space, so access to shared memory must be synchronized.
//The sync package provides useful primitives, although you won't need them much in Go as there are other primitives.
func Routine() {
	go say("world")
	say("hello")
}
