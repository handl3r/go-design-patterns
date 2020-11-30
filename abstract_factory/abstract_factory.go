package abstract_factory

import "fmt"

// Abstract product
type Runnable interface {
	Run()
}

type Eatable interface {
	Eat()
}

// Concrete product
type LuxuryCar struct {
	name string
}

func NewLuxuryCar() *LuxuryCar {
	return &LuxuryCar{
		name: "Lux Car",
	}
}

func (car *LuxuryCar) Run() {
	fmt.Println("LuxuryCar is running")
}

type NormalCar struct {
	name string
}

func NewNormalCar() *NormalCar {
	return &NormalCar{
		name: "Nor Car",
	}
}

func (bus *NormalCar) Run() {
	fmt.Println("NormalCar is running")
}

type LuxuryAnimal struct {
	name string
}

func NewLuxuryAnimal() *LuxuryAnimal {
	return &LuxuryAnimal{
		name: "Lux An",
	}
}

func (c *LuxuryAnimal) Eat() {
	fmt.Println("LuxuryAnimal is eating")
}

type NormalAnimal struct {
	name string
}

func NewNormalAnimal() *NormalAnimal {
	return &NormalAnimal{
		name: "Nor An",
	}
}

func (d *NormalAnimal) Eat() {
	fmt.Println("NormalAnimal is eating")
}

// Abstract Factory
type AbstractFactory interface {
	CreateRunnable() Runnable
	CreateEatable() Eatable
}

// Concrete Factory
type LuxuryFactory struct {
	name string
}

func NewLuxuryFactory() *LuxuryFactory {
	return &LuxuryFactory{
		name: "Lux Fac",
	}
}

func (lf *LuxuryFactory) CreateRunnable() Runnable {
	return NewLuxuryCar()
}

func (lf *LuxuryFactory) CreateEatable() Eatable {
	return NewLuxuryAnimal()
}

type NormalFactory struct {
	name string
}

func NewNormalFactory() *NormalFactory {
	return &NormalFactory{
		name: "Nor Fac",
	}
}

func (nf *NormalFactory) CreateRunnable() Runnable {
	return NewNormalCar()
}

func (nf *NormalFactory) CreateEatable() Eatable {
	return NewNormalAnimal()
}

//

func NewFactory(typeFac string) AbstractFactory {
	switch typeFac {
	case "Luxury":
		return NewLuxuryFactory()
	case "Normal":
		return NewNormalFactory()
	default:
		return nil
	}
}
