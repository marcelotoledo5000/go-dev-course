package main

import "fmt"

type Vehicle interface {
	start() string
}

type Car struct {
	Name string
}

type Motorcycle struct {
	Name string
}

func (c Car) start() string {
	return "The car " + c.Name + " has been started!"
}

func (mc Motorcycle) start() string {
	return "The motorcycle " + mc.Name + " has been started!"
}

// interface
func startVehicle(v Vehicle) string {
	return v.start()
}

func main() {
	car := Car{"Mercedes"}
	mc := Motorcycle{"Ninja"}

	fmt.Println(startVehicle(car)) // fmt.Println(car.start())
	fmt.Println(startVehicle(mc))  // fmt.Println(mc.start())
}
