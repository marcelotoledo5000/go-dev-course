package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	color string
}

func main() {
	car := Car{"Honda Fit", 2015, "Blue"}
	result, _ := json.Marshal(car)
	// fmt.Println(result) // print value of chars (ASCII)
	fmt.Println(car)
	fmt.Println(string(result))
}
