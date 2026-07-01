package main

import "fmt"

/*
	Closure é uma função que "lembra" das variáveis que existiam no lugar onde ela foi criada, mesmo depois que esse lugar já terminou de executar.
*/
func closure() func() {
	texto := "Dentro da função Closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func contador() func() int {
	contador := 0

	return func() int {
		contador++
		return contador
	}
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoClosure := closure()
	funcaoClosure()

	c := contador()

	fmt.Println(c()) //1
	fmt.Println(c()) //2
	fmt.Println(c()) //3
	fmt.Println(c()) //4
	fmt.Println(c()) //5
	fmt.Println(c()) //6
	fmt.Println(c()) //7
}
