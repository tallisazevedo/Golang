package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup // Criando um WaitGroup

	waitGroup.Add(2) // Declarando que meu WaitGroup será composta por 2 goroutines

	// Primeira goroutine
	go func() {
		escrever("Olá mundo!")
		waitGroup.Done() // O Done() ele avisa para o WaitGroup que uma das goroutines já foi executada. Ele faz o papel de -1.
	}()

	// Segunda goroutine
	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // O Done() ele avisa para o WaitGroup que uma das goroutines já foi executada. Ele faz o papel de -1.
	}()

	waitGroup.Wait() // O Wait() passa para a função main que ela precisa esperar as goroutines terminarem para finalizar
}

func escrever(texto string) {
	for range 5 {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
