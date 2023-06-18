package main

import "fmt"

func main() {
	for i := 0; i < 5; i++ { // basic form
		fmt.Println(i)
	}

	x := 0
	for x < 10 { // like while
		fmt.Println(x)
		x++
	}

	for { // infinity loop
		x += 5
		fmt.Println(x)

		if x == 100 {
			break
		}
	}
}
