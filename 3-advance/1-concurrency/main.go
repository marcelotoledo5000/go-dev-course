package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	var str string
	fmt.Scanln(&str)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, " -> ", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
}
