package main

import "fmt"

func main() {
	var variavel1 int = 10
	var variavel2 int = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)
	/* O valor da variável "variavel2" não muda, pois quando fazemos a atribuição de uma variável à outra, estamos passando uma cópia do valor */

	var variavel3 int
	var ponteiro *int /* Estamos criando um ponteiro, utilizando "*" antes do tipo */
	fmt.Println(variavel3, ponteiro)

	variavel3 = 100
	ponteiro = &variavel3 /* A variável "ponteiro", que é de fato um PONTEIRO, nesse caso, está apontando para o endereço de memória onde está armazenado a variável "variavel3" */

	fmt.Print(variavel3, ponteiro) /* A variável "ponteiro" irá imprimir o endereço de memória onde a variável "variavel3" está armazenada. */

	fmt.Println(variavel3, *ponteiro) /* A variável ponteiro irá imprimir o valor armazenado dentro do endereço de memória da variável "variavel3", pois o "*" antes da variável, está fazendo um DESFERENCIAÇÃO, que resume em buscar o valor presente dentro do endereço de memória */
}
