package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var mut sync.Mutex

func main() {
	go runProcess("P1", 20)
	go runProcess("P2", 20)

	var str string
	fmt.Scanln(&str)
	fmt.Println("Final result: ", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		mut.Lock()
		result++
		fmt.Println(name, " -> ", i, "Partial result: ", result)
		mut.Unlock()
	}
}
