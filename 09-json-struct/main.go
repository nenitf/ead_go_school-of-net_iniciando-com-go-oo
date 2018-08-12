package main

import (
	"encoding/json"
	"fmt"
)

/*

Existem dois tipos de modificadores de acessos em structs:
Exported ou unexported

Igual as variaveis de pacote, isso é definido através da letra Maiuscula - exportavel - ou Minuscula - não exportavel

*/

// struct base
type Car struct {
	Name  string
	Year  int
	color string
}

type SuperCar struct {
	Car    Car
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Nome: %s\n Ano: %d\n Cor: %s", c.Name, c.Year, c.color)
}

func main() {
	sCar := SuperCar{
		Car{
			"Corsa",
			2002,
			"azul",
		},
		true,
	}

	fmt.Println(sCar.Car.info()) // imprime a cor pois está sendo enviada de fora para dentro

	// resultado com tabela asci
	result, _ := json.Marshal(sCar)

	// conversão asci em string
	fmt.Println(string(result)) // não é possivel ver sCar.Car.color - letra minuscula

}
