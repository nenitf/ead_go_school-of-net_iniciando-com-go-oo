package main

/*
 * Tornar o código mais adaptativo
 */

import "fmt"

// é considerado vehicle todos que implementarem o contrato
type vehicle interface {
	start() string
}

type car struct {
	name string
}

// torna car um vehicle
func (c car) start() string {
	return "The car " + c.name + " has been started"
}

type motorcycle struct {
	name string
}

// torna motorcycle um vehicle
func (m motorcycle) start() string {
	return "The motorcycle " + m.name + " has been started"
}

// vehicle é qualquer struct que implemente a interface vehicle
// a implementação é implicita, diferente de "class x implements i"
func startVehicle(v vehicle) string {
	return v.start()
}

func main() {
	c := car{"Gol"}
	m := motorcycle{"XPTO"}
	fmt.Println(startVehicle(c))
	fmt.Println(startVehicle(m))
}
