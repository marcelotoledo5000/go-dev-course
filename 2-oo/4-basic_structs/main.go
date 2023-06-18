package main

import "fmt"

// make a new type
type SuperNumber int

// make a new struct, like a class
type Car struct {
	Name  string
	Year  int
	Color string
}

// attach a new method to struct
func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor %s", c.Name, c.Year, c.Color)
}

func main() {
	x := []SuperNumber{}
	x = append(x, 100, 101, 102)
	fmt.Println(x)

	mustang := Car{"Mustang", 2023, "Black"}
	ferrari := Car{"Ferrari", 2023, "Red"}

	fmt.Println(mustang)
	fmt.Println(ferrari.Year)
	fmt.Println(ferrari.Name, ferrari.Color, ferrari.Year)
	fmt.Println(mustang.info())
}
