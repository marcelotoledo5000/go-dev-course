package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func init() {
	// Allows the use of all available cores
	// runtime.GOMAXPROCS(runtime.NumCPU())
	// Limiting processing usage.
	runtime.GOMAXPROCS(1)
}

func main() {
	fmt.Println(runtime.NumCPU())
	waitGroup.Add(2)
	go runProcess("P1", 20)
	go runProcess("P2", 20)
	waitGroup.Wait()
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, " -> ", i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()
}
