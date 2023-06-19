package main

import (
	"fmt"
	"time"
)

func main() {
	channel := make(chan int)

	// First routine.
	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
		}
	}()

	// Second routine.
	go func() {
		for {
			fmt.Println(<-channel)
		}
	}()
	time.Sleep(time.Second)

	fmt.Println(channel)
}
