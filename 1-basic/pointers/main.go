package main

import "fmt"

func main() {
	x := 10
	fmt.Println(x)  // print value
	fmt.Println(&x) // print memory address

	y := &x // referencing
	fmt.Println(y)
	// To access the value or object located in a memory location stored
	// in a pointer or another value interpreted as such; to access
	// a value being referenced by something else.
	fmt.Println(*y) // print value

	*y = 20
	fmt.Println(x)

	var z *int = &x // pointer
	fmt.Println(z)

	b := 10
	fmt.Println(xpto(b))
	fmt.Println(b)
	fmt.Println(xpto_pointer(&b))
	fmt.Println(b)
}

func xpto(a int) int {
	a = a + 1
	return a
}

func xpto_pointer(a *int) int {
	*a = *a + 1
	return *a
}
