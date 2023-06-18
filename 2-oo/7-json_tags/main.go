package main

import (
	"encoding/json"
	"fmt"
)

// struct with tags
type Car struct {
	Name  string `json:"Model"`
	Year  int    `json:"-"`
	Color string
}

func main() {
	car := Car{"Fit", 2015, "Blue"}
	result, _ := json.Marshal(car)
	fmt.Println(car)
	fmt.Println(string(result))
}
