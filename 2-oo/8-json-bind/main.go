package main

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	Color string
}

// More about this subject:
// https://medium.com/@aalves/golang-transformando-dados-marshal-unmarshal-decode-encode-porque-e-quando-2048ecbe1075
func main() {
	var car Car
	strValues := `{ "Name": "Fit", "Year": 2015, "Color": "Blue" }`
	data := []byte(strValues)
	fmt.Println(data)

	err := json.Unmarshal(data, &car)
	fmt.Println(err, car.Name, car.Year, car.Color)
	fmt.Println(car)
}
