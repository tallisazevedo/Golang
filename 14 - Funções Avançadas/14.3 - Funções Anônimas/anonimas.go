package main

import "fmt"

func main() {

	/*
		- A função anônima é criada sem nome, ela é executada pelo "()" que vem logo após a sua declaração.
		- Ela também pode ter parâmetros e retornos.
	*/
	func() {
		fmt.Println("Olá mundo")
	}()

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s", texto)
	}("Aqui é onde passamos o parâmetro")

	fmt.Println(retorno)
}
