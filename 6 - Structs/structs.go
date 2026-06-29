package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	// Criação de usuário com variável explcíta
	var u usuario
	u.nome = "Tallis"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua dos Bobos", 0}

	// Criação de usuário com variável implcíta
	u2 := usuario{"tallis", 22, enderecoExemplo}
	fmt.Println(u2)

	// Criação de usuáro sem passar todos dados declarados na Struct
	u3 := usuario{idade: 22}
	fmt.Println(u3)
}
