package main

import (
	"fmt"

	"./car"
)

func main() {
	// ferrari := car.Car{"Ferrari", "Black"} // Warning: Car struct literal uses unkeyed fields
	ferrari := car.Car{Name: "Ferrari", Color: "Black"} // Solution to warning above.

	fmt.Println(ferrari.Start())
}
