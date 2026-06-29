package main

import "fmt"

// Criação padrão de uma função
func Somar(n1 int, n2 int) int {
	return n1 + n2
}

// Declarando variável com mais de um retorno
func CalculosMatematicos(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 - n2
}

func main() {
	// Declaração uma função para uma variável (Ela também podem ter retorno caso seja necessário)
	var f = func(txt string) {
		fmt.Println(txt)
	}

	// Utilizando a função Somar()
	soma := Somar(10, 13)
	fmt.Println(soma)

	// Utilizando a função f()
	f("Esse será meu texto exibido")

	// Utilizando a função CalculosMatematicos(), porém, quero apenas a subtração
	_, sub := CalculosMatematicos(10, 13)
	fmt.Println(sub)

}
