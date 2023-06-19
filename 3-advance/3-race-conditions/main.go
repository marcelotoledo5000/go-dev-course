package main

import (
	"fmt"
	"math/rand"
	"time"
)

var result int

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	var str string
	fmt.Scanln(&str)
	fmt.Println("Final result: ", result) // Should be 40 but currently is 20.
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		z := result
		z++
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		result = z
		fmt.Println(name, " -> ", i, "Partial result: ", result)
	}
}
