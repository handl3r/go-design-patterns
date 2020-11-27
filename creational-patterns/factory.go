package creational_patterns

import "fmt"

type Factory struct {
}

func (f *Factory) GenerateProduct(kind string) Runnable {
	var transport Runnable
	switch kind {
	case "car":
		transport = NewCar()
	case "bike":
		transport = NewBike()
	default:
		transport = nil
	}
	return transport
}

// Interface product
type Runnable interface {
	Run()
}

// Concrete Type
type Car struct {
	name string
}

func NewCar() *Car {
	return &Car{
		name: "car",
	}
}

func (c *Car) Run() {
	fmt.Println("running")
}

type Bike struct {
	name string
}

func NewBike() *Bike {
	return &Bike{
		name: "bike",
	}
}

func (c Bike) Run() {
	fmt.Println("running")
}
