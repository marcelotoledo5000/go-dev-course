package car

type Car struct {
	Name, Color string
}

func (c Car) Start() string {
	return c.Name + " has been started!"
}
