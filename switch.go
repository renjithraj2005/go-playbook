package main

import (
	"fmt"
	"runtime"
	"time"
)

func switchSample() {
	fmt.Println("Go runs on ")
	//Switch cases evaluate cases from top to bottom, stopping when a case succeeds.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac")
	case "linux":
		fmt.Println("linux")
	default:
		fmt.Printf("%s.", os)
	}
	//Switch with no condition
	//Switch without a condition is the same as switch true.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
