package main

/*
interfaces vazias servem para inicializar valores default para a interface
*/
import "fmt"

// cria o tipo de variavel Names
// que é uma interface sem obrigações
type Names []interface{}

// A função utiliza o o apontamento da memoria de variavel DO TIPO Names
// Names é a interface vazia
func (n *Names) load() {
	// altera o valor no apontamento da memoria
	*n = Names{
		"Wesley",
		"Davi",
		1,
	}
}

func (n Names) printNames() {
	fmt.Println(n)
}

func main() {
	// cria uma variavel DO TIPO Names
	// Names é uma interface
	var names Names
	names.load()
	names.printNames()
}
