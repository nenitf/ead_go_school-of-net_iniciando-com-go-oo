package main

import "fmt"

// maps são como arrays-associativos[php], dicionarios[python]
func main() {
	// map[tipo-indice]tipo-dado
	m := make(map[string]int)    // indice string valor inteiro
	m2 := map[string]int{}       // sem utilização do make
	m3 := map[string]int{"x": 5} // sem make com inicialização
	fmt.Println(m2)
	fmt.Println(m3)

	m["chave1"] = 10
	m["chave2"] = 11
	fmt.Println(m)
	fmt.Println(m["chave1"])

	// remover um valor
	delete(m, "chave2")
	fmt.Println(m)
	fmt.Println(m["chave2"])  // foi deletado MAS retorna como zero
	fmt.Println(m["chave99"]) // não existe e TBM retorna como zero
	fmt.Println(m)            // mostra apenas a chave1

	// blankIdentifier recebe o valor, e a variavel "existe" recebe boolean
	// relembrando: o blank é utilizado quando não queremos erro ao não utilizar uma variavel recebida durante o duplo retorno
	_, existe1 := m["chave1"] // true
	fmt.Println(existe1)
	_, existe2 := m["chave2"] // false
	fmt.Println(existe2)
	_, existe3 := m["chave99"] // false
	fmt.Println(existe3)

	if _, exists := m["chave2"]; exists {
		fmt.Println("o valor existe")
	} else {
		fmt.Println("o valor n existe")
	}
}
