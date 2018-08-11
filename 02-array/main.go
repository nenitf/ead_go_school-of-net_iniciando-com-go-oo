package main

import "fmt"

// array com tamanho definido

func main() {
	// array com 10 posições inteiras, x[0] até x[9]
	var x [10]int
	// array com 10 posições string
	var y [10]string

	fmt.Println(x)      // inicia om as 10 posições zeradas
	fmt.Println(y)      // inicia com as posições vazias
	fmt.Println(len(x)) // length do array

	x[2] = 10
	x[5] = 8
	fmt.Println(x)
	fmt.Println(x[2])
}
