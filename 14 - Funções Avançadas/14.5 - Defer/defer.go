package main

import "fmt"

func funcao1() {
	fmt.Println("Executando função 1")
}

func funcao2() {
	fmt.Println("Executando função 2")
}

func verificarAprovacaoAluno(nota1, nota2 int) bool {
	/*
		Utilizamos o DEFER para adiar uma execução. O DEFER sempre será executado antes de um return.
		Ou seja, por mais que a função esteja na primeira linha da função "verificarAprovacaoAluno()", ela será executada somente dentro do if antes do return, ou depois do if.
	*/
	defer fmt.Println("Média calculada. Resultado será retornado.")
	fmt.Println("Iniciando cálculo de média.")

	media := (nota1 + nota2) / 2

	if media > 6 {
		return true
	}
	return false
}

func main() {
	resultado := verificarAprovacaoAluno(7, 7)

	fmt.Println(resultado)
}
