package main

import "fmt"

func main() {
	// Creating a channel
	msg := make(chan string)

	go func() {
		msg <- "Hello World"
	}()

	result := <-msg
	fmt.Println(result)
}
