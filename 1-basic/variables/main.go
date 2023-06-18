package main

import "fmt"

// var b int
// b = 15

// var txt, other string = "Hello", "World"

const (
	aa        = 55
	bb int    = 66
	Cc string = "abc"
)
const xvz int = 1234

func main() {
	a := 10
	b := "Hello"
	c := 10.45
	d := false
	e := 'W'
	f := `Uouuu`

	const xpto = "ABC"

	fmt.Printf("Current Values: \n")
	fmt.Printf("%v \n", a)
	fmt.Printf("%v \n", b)
	fmt.Printf("%v \n", c)
	fmt.Printf("%v \n", d)
	fmt.Printf("%v \n", e)
	fmt.Printf("%v \n", f)
	fmt.Printf("Current Types: \n")
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)
	fmt.Printf("%T \n", c)
	fmt.Printf("%T \n", d)
	fmt.Printf("%T \n", e)
	fmt.Printf("%T \n", f)
}
