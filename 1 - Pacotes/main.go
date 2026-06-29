package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("tallis@teste.com")

	if erro == nil {
		fmt.Println("E-mail válido")
	}
}
