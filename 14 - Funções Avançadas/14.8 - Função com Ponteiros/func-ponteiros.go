package main

import "fmt"

func inverterSinal(n int) int {
	return n * -1
}

func inverterSinalComPonteiro(n *int) {
	// Utilizamos o "*" na frente da variável pois queremos fazer a desferenciação para utilizar o valor dentro do endereço de memória
	*n = *n * -1
}

func main() {
	/*
		Nesse caso, a variável número, continuará guardando o valor 20 dentro do seu endereço de memória, pois o que passamos dentro dos parâmetros da função, foi uma cópia do seu valor.
	*/
	numero := 20

	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	/*
		Nesse novo caso, como estamos utilizando uma função que utiliza ponteiro para apontar para o endereço de memória da variável, o valor da variável irá mudar dentro do seu endereço de memória.
	*/
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalComPonteiro(&novoNumero) // Passamos usando o "&" pois a função espera um endereço de memória
	fmt.Println(novoNumero)
}
