package main

import "fmt"

func main() {
	num := mult(5)
	fmt.Println(mult(num))
	fmt.Println(namedReturn("Marc"))
	fmt.Println(moreReturns("Marc", "Toledo"))
	surname, name := moreReturns("Marcelo", "Toledo")
	fmt.Println(surname+",", name)
	fmt.Println(variadicFunc(10, 5, 2, 23))

	// closures
	// fmt.Println(name); {
	// 	y := "Hello"
	// 	...
	// }

	// anonimous functions
	z := 0
	add := func() int {
		z += 2
		return z
	}

	fmt.Println(add())
	fmt.Println(add())
	fmt.Println(funcInsideFunct())
}

func mult(a int) int {
	return a * a
}

func namedReturn(a string) (x string) {
	x = a
	return
}

func moreReturns(a string, b string) (string, string) {
	return b, a
}

func variadicFunc(x ...int) int {
	var result int
	for _, v := range x { // _ -> index, v -> value
		result += v
	}

	return result
}

func funcInsideFunct() func() int {
	x := 13
	return func() int {
		return x + 0
	}
}
