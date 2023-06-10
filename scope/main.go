package main

import "fmt"

var y int = 20

func main() {
	x := 10
	fmt.Println(x)
	fmt.Println(y * x)
	printNum()
	fmt.Println(z)
	printZ()
}

func printNum() {
	fmt.Println(y)
}
