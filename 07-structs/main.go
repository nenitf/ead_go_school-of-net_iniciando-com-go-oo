package main

import "fmt"

// lembram classes

type SuperNumber int // criação de um tipo

// criação e struct
type Car struct {
	Name  string
	Year  int
	Color string
}

func (c Car) info() string {
	return fmt.Sprintf("Nome: %s\n Ano: %d\n Cor: %s", c.Name, c.Year, c.Color)
	// %s string
	// %d digito => int
}

func main() {
	car1 := Car{"Corsa", 2002, "azul"}
	car2 := Car{"Elba", 1996, "branco"}
	fmt.Println(car1)
	fmt.Println(car2)
	fmt.Println(car1.Name)
	fmt.Println(car1.info())
}
