package main

// pegar um json e criar um objeto

import (
	"encoding/json"
	"fmt"
)

type Car struct {
	Name  string
	Year  int
	color string // não recebe valores
}

func main() {
	var car Car
	data := []byte(`{"Name":"Gol", "Year": 2017, "Color":"Azul"}`)

	json.Unmarshal(data, &car) // deve ser passado o apontamento para car
	fmt.Println(car)
}
