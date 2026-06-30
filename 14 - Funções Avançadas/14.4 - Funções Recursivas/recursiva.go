package main

import "fmt"

/*
	- As funções recursivas, sempre terão um caso base onde ela irá parar a sua execução.
	- Para chegar no seu caso baso, ela chama ela mesmo.
*/
func fibonacci(posicao uint) uint {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}

func main() {
	fmt.Println("Funções recursivas")

	posicao := 4

	resultado := fibonacci(uint(posicao))
	fmt.Println(resultado)
}
