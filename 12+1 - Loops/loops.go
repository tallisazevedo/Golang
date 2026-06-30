package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando I")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 10; j++ {
		fmt.Println(j)
		time.Sleep(time.Second)
	}

	slice := []string{"tallis", "rafaela", "gigi"}

	slice = append(slice, "Binho")

	/*
		- Utiliza-se o RANGE quando queremos percorrer um ARRAY ou um SLICE
		- Por padrão, ele sempre vai passsar o INDICE e o VALOR.
	*/
	for indice, nome := range slice {
		fmt.Println(indice, nome)
	}

	/*
		- Alternativa para não pegar o indice
	*/
	for _, nome := range slice {
		fmt.Println(nome)
	}

	/*
		- RANGE permite iterar sobre uma string
		- Caso declara dessa forma, ela irá exibir o código da tabela ASCII
	*/
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	/*
	 - Declarando dessa forma ele exibirá a letra.
	*/
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Tallis",
		"Sobrenome": "Azevedo",
	}

	/*
		- Para iterar sobre um Map, ele atribui chave e valor
	*/
	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
