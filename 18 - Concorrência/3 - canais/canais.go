package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string) // Cria um canal que irá passar dados do tipo String

	go escrever("Olá, mundo!", canal) // Executa a função utilizando goroutine

	for {
		mensagem, aberto := <-canal // A sintaze "<-canal" antes da variável declara que está esperando um valor, e o seu programa fica parado até que o valor chegue. A variável aberto, guarda o estado do canal, se está aberto o fechado.
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}
	fmt.Println("Fim de programa")
}

func escrever(texto string, canal chan string) { // A função recebe um novo parâmetro para receber o canal que será transmitido o dado.
	for range 5 {
		canal <- texto // A sintaxe "<-" depois da variável, está declarando que o valor que a variável seguinte guarda, está sendo transmitida.
		time.Sleep(time.Second)
	}
	close(canal) // Após a execução o loop, o close fecha o canal.
}
