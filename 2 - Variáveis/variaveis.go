package main

import "fmt"

func main() {
	var variavel string = "Tallis"
	fmt.Println(variavel)

	nome := "Tallis"
	fmt.Println(nome)

	var (
		nomeCompleto string = "Tallis Azevedo"
		idade        int    = 22
	)
	fmt.Println(nomeCompleto, idade)

	const umaConstante = "Declarando uma constante"

	numero1, numero2 := 1, 2
	fmt.Println(numero1, numero2)

	numero1, numero2 = numero2, numero1
	fmt.Println(numero1, numero2)
}
