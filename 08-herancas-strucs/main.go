package main

import "fmt"

// struct base
type Car struct {
	Name  string
	Year  int
	Color string
}

type SuperCar struct {
	Car    Car
	CanFly bool
	Name   string // nome para sobescrever dentro de car
}

func (c Car) info() string {
	return fmt.Sprintf("Nome: %s\n Ano: %d\n Cor: %s", c.Name, c.Year, c.Color)
}

func main() {
	car1 := Car{"Corsa", 2002, "azul"}
	/*
		// nomeado
		sCar := SuperCar{
			Car: Car{
				"Corsa",
				2002,
				"azul",
			},
			CanFly: true,
			Name: "Elba"
		}

		// não nomeado
		sCar := SuperCar{
			 Car{
				"Corsa",
				2002,
				"azul",
			},
			true,
			"Elba",
		}
	*/

	// aproveitando var car1
	sCar := SuperCar{
		car1,
		true,
		"Elba",
	}
	fmt.Println(sCar)
	fmt.Println(sCar.Car.Name)
	fmt.Println(sCar.Car.info())
}
