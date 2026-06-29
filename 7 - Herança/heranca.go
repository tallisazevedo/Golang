package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"tallis", "azevedo", 22, 188}
	fmt.Println(p1)

	e1 := estudante{p1, "CC", "UVV"}
	fmt.Println(e1.nome)
	fmt.Println(e1)
}
