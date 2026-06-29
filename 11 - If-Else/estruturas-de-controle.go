package main

import "fmt"

func main() {
	fmt.Println("Estruturas de  Controle")

	numero := 10

	/*
		Estrutura de controle com IF e ELSE
	*/
	if numero > 15 {
		fmt.Println("Maior que 15s")
	} else if numero >= 10 {
		fmt.Println("Número é maior ou igual a 10")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	/*
		IF INIT
		- Ele proporciona capacidade de iniciar uma nova variável atribuir um valor e utilizar as estruturas de controle para avaliar a variável criada
		- A variável só pode ser usada no escopo da estrutura de controle
	*/
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Número é maior que zero")
	} else {
		fmt.Println("Menor que 0")
	}
}
