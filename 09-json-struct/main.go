package main

import (
	"encoding/json"
	"fmt"
)

/*

Existem dois tipos de modificadores de acessos em structs:
Exported ou unexported
Exported: visivel para somente o mesmo pacote ou própria classe

Igual as variaveis de pacote, isso é definido através da letra Maiuscula - exportavel - ou Minuscula - não exportavel

*/

// struct base
type Car struct {
	Name  string `json:"teste"` // tag: altera o nome do campo
	Year  int    `json:"-"`     // tag: omite para json
	color string // unexported
}

type SuperCar struct {
	Car    Car
	CanFly bool
}

func (c Car) info() string {
	return fmt.Sprintf("Nome: %s\n Ano: %d\n Cor: %s", c.Name, c.Year, c.color)
}

func main() {

	car := Car{
		"Corsa",
		2002,
		"azul",
	}
	sCar := SuperCar{
		car,
		true,
	}

	fmt.Println(sCar.Car.info()) // imprime a cor pois está sendo enviada de fora para dentro

	// resultado com tabela ascii
	result, _ := json.Marshal(sCar)

	// conversão asci em string
	fmt.Println(string(result)) // não é possivel ver sCar.Car.color - letra minuscula

	result2, _ := json.Marshal(car)
	fmt.Println(string(result2))
}
