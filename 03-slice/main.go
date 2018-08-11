package main

import "fmt"

// diferente do array, slices não precisam ter tamanho definidos

func main() {
	// cria slice de inteiros sem qtd de posição inicial
	slice1 := make([]int, 0)
	fmt.Println(slice1) // array VAZIO
	// fmt.Println(len(slice1))
	// cria slice de inteiros com 5 posições iniciais - zeradas
	slice2 := make([]int, 5)
	fmt.Println(slice2)

	// slice1 recebe o acrescime de 1 na ultima posição do slice1
	slice1 = append(slice1, 1)
	fmt.Println(slice1)

	slice2 = append(slice2, 2, 3, 5, 6)
	fmt.Println(slice2)

	fmt.Println("---------------------------")

	// cria slice com 2 posições e capacidade de 5
	// cada vez que é dado append ultrapassando a capacidade, a msm dobra e é criada uma nova variavel - apontamento de memoria - com os valores corretos
	slice3 := make([]int, 2, 5)
	for i := 0; i < 20; i++ {
		slice3 = append(slice3, i)
		fmt.Println("len: ", len(slice3), " capacidade: ", cap(slice3))
	}

	// quando alguma var recebe um slice ele recebe o msm ponteiro
	slice4 := slice3
	slice4[0] = 99
	fmt.Println("valores slice3: ", slice3, " len: ", len(slice3), " cap: ", cap(slice3))
	fmt.Println("valores slice4: ", slice4, " len: ", len(slice4), " cap: ", cap(slice4))

	// vou estoura a capacidade do slice3 para ser criado um novo ponteiro e tornar ambas variaveis independentes
	slice3 = append(slice3, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0)

	fmt.Println("len slice3: ", len(slice3), " cap: ", cap(slice3))
	fmt.Println("len slice4: ", len(slice4), " cap: ", cap(slice4))

	fmt.Println("-------------")

	// array com [valor]
	arrayString := [4]string{
		"felipe",
		"jay",
		"pimenta",
		"fralda", // terminar sempre com virgula
	}

	fmt.Println(arrayString)

	// array sem [valor]
	sliceString := []string{
		"felipe",
		"jay",
		"pimenta",
		"fralda", // terminar sempre com virgula
	}

	fmt.Println(sliceString)

	fmt.Println(sliceString[1])
	fmt.Println(sliceString[:])   // infinito
	fmt.Println(sliceString[:2])  // infinito até a segunda [1], resultado: [0] e [1]
	fmt.Println(sliceString[2:])  // da segunda posição [1] até infinito, resultado: [3] e [4]
	fmt.Println(sliceString[1:2]) // posições depois da primeira até a segunda, resultado: [1] - segunda
	fmt.Println(sliceString[2:4]) // posições depois da segunda [1] até a quarta [3], resultado: [2] e [3]

}
