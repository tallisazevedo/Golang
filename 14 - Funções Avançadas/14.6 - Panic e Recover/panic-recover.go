package main

import "fmt"

func RecuperarPanico() {
	/*
	 A função RECOVER funciona como um tratamento para o PANIC.
	 Ela retorna um erro(caso tenha estourado um panic) ou nil(caso nenhum panic tenha sido estourado).
	 Ela permite fazer o tratamento do erro e seguir com a execução.
	*/
	if r := recover(); r != nil {
		fmt.Println("A EXECUÇÃO FOI RECUPERADA COM SUCESSO")
	}
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer RecuperarPanico()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	/*
		A função "PANIC" quebra toda a execução do programa. O programa entre em pânico.
		Porém, antes dela parar tudo, ela chamar todos as funções que estão adiadas pelo DEFER.
		Resumindo, ela emite um erro GRAVE.
	*/
	panic("A MÉDIA É EXATAMENTE 6")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("FIM DA EXECUÇÃO DO PROGRAMA!")
}
