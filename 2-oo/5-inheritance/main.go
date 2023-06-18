package main

import "fmt"

// make a new struct, like a class
type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car
	Name   string // optional to overwrite
	CanFly bool
}

// attach a new method to struct
func (c Car) info() string {
	return fmt.Sprintf("Car: %s\nYear: %d\nColor %s\n", c.Name, c.Year, c.Color)
}

func main() {
	mustang := Car{"Mustang", 2023, "Black"}
	ferrari := Car{"Ferrari", 2023, "Red"}
	bugatti := Car{"Bugatti", 2023, "Yellow"}

	fmt.Println(mustang)
	fmt.Println(ferrari.Year)
	fmt.Println(ferrari.Name, ferrari.Color, ferrari.Year)
	fmt.Println(mustang.info())

	delorean := SuperCar{
		Car:    Car{"Delorean", 1985, "Silver"},
		CanFly: true,
	}
	fmt.Println(delorean)
	fmt.Println(delorean.Name)
	fmt.Println(delorean.info())

	supBugatti := SuperCar{
		Car:    bugatti,
		CanFly: false,
		Name:   "Super Bugatti",
	}
	fmt.Println(supBugatti)
	fmt.Println(supBugatti.Name)
	fmt.Println(supBugatti.Car.Name)
	fmt.Println(supBugatti.info())
}
