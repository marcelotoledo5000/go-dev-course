package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	x := funnel(generateMsg("Hello Gophers"), generateMsg("Hello World"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-x)
	}

	fmt.Println("Finished...")
}

func generateMsg(str string) <-chan string {
	channel := make(chan string)
	go func() {
		for i := 0; ; i++ {
			channel <- fmt.Sprintf("String: %s - Value: %d", str, i)
			time.Sleep(time.Duration(rand.Intn(255)) * time.Millisecond)
		}
	}()
	return channel
}

// fan in
func funnel(channel1, channel2 <-chan string) <-chan string {
	channel := make(chan string)

	go func() {
		for {
			channel <- <-channel1
		}
	}()

	go func() {
		for {
			channel <- <-channel2
		}
	}()

	return channel
}
