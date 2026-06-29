package main

import "fmt"

func main() {
	/* OPERADORES ARITMÉTICOS */
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 5
	multiplicacao := 10 * 2
	restoDaDivisao := 10 % 3

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	/* OPERADORES DE ATRIBUIÇÃO */
	var variavel1 string = "Atribuição com tipo explicíto"
	variavel2 := "Atribuição com tipo implícito"
	fmt.Println(variavel1, variavel2)

	/* OPERADORES RELACIONAIS */
	fmt.Println(1 > 2)  // MAIOR QUE
	fmt.Println(1 >= 2) // MAIOR OU IGUAL A
	fmt.Println(1 == 2) // IGUAL A
	fmt.Println(1 < 2)  // MENOR QUE
	fmt.Println(1 <= 2) // MENOR OU IGUAL A
	fmt.Println(1 != 2) // DIFERENTE QUE

	/* OPERADORES LÓGICOS */
	fmt.Print("-------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) // AND - Só retorna TRUE se as duas condições forem verdadeiras
	fmt.Println(verdadeiro || falso) // OR - Retorna TRUE se uma das duas condições for verdadeira
	fmt.Println(!verdadeiro)         // NOT - Inverte o valor da condição, se for TRUE, retorna FALSE
	fmt.Println(!falso)              // NOT - Inverte o valor da condição, se for TRUE, retorna FALSE

	/* OPERADORES UNÁRIOS */
	fmt.Print("-------------")
	numero := 10
	numero++ // Itera +1 para a variável
	numero += 10
	fmt.Println(numero)

	numero--
	numero -= 9
	fmt.Println(numero)

	numero *= 10
	numero /= 10
	numero %= 10

	fmt.Println(numero)
}
