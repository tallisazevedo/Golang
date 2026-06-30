package main

import "fmt"

/*
	Retorno nomeado faz com que seja documentado o retorno da função. Dessa forma, permitindo fazer um "naked return" (ou melhor, retorno vazio, sem argumentos).
*/
func calculosMatematicos(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}

func main() {

	soma, subtracao := calculosMatematicos(5, 2)
	fmt.Println(soma, subtracao)

}
