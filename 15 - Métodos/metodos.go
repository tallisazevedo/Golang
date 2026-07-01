package main

import "fmt"

/*
A diferença entre uma função e um método, é que o método sempre estará ligado a alguma coisa, seja ela um struct, uma interface...
*/

type usuario struct {
	nome  string
	idade int8
}

/*
	Essa é a sintaxe utilizada para criar um método. Quando passando o "(u usuario)", estamos fazendo relação do método com a struct "usuario"
*/

func (u usuario) salvar() {
	fmt.Println("Salvando o usuário: ", u.nome)
}

func (u usuario) maioridade() bool {
	return u.idade >= 18
}

func (u *usuario) fazerAniversario() {
	// Aqui não é necessário fazer a desferenciação
	u.idade++
}

func main() {
	usuario1 := usuario{"Usuário 1", 22}

	fmt.Println(usuario1)
	usuario1.salvar()

	usuario2 := usuario{"Usuário 2", 30}
	fmt.Println(usuario2)

	fmt.Println(usuario2.maioridade())

	usuario2.fazerAniversario()
	fmt.Println(usuario2.idade)
}
