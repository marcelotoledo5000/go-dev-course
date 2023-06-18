package main

import "fmt"

// empty interface
type Names []interface{}

func (n *Names) load() {
	*n = Names{"Odin", "Thor", 1}
}

func (n Names) printNames() {
	fmt.Println(n)
}

func main() {
	var names Names
	names.load()
	names.printNames()
}
