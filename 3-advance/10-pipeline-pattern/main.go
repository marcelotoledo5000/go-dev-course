package main

import "fmt"

// Pipeline pattern
func main() {
	numbers := generate(6, 8, 20, 12, 34)
	result := divide(numbers)

	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
	fmt.Println(<-result)
}

func generate(numbers ...int) chan int {
	channel := make(chan int)

	go func() {
		for _, number := range numbers {
			channel <- number
		}
	}()

	return channel
}

func divide(input chan int) chan int {
	channel := make(chan int)

	go func() {
		for number := range input {
			channel <- number / 2
		}
		close(channel)
	}()

	return channel
}
