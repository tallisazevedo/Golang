package main

import "fmt"

/*
	- Função variática permite que passamos mais de um argumento para o parâmetro
	- Cada função só pode ter um parâmetro variático
	- O parâmetro variático deve ser sempre o último declarado
	- Ele vira um Slice, então pode ser iterado
*/
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println("Funções variáticas")
	total := soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 2)
	fmt.Println(total)

	escrever("Eu amo a minha namorada", 1, 2, 3, 4, 5, 6, 7)
}
