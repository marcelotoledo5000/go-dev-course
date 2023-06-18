package main

import (
	"fmt"
)

func main() {
	// map definition with string key and int value.
	m := make(map[string]int)
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30

	fmt.Println(m)
	delete(m, "b")
	fmt.Println(m)

	_, existsA := m["a"]
	fmt.Println(existsA)

	valueB, existsB := m["b"]
	fmt.Println(valueB, existsB)

	valueC, existsC := m["c"]
	fmt.Println(valueC, existsC)

	// declare an empty map
	var x = map[string]int{}
	fmt.Println(x)

	// declare a map with values
	newMap := map[string]int{"x": 10, "y": 20}
	fmt.Println(newMap)

	if value, exists := m["b"]; exists {
		fmt.Println(value)
	} else {
		fmt.Println("Ops!")
	}
}
