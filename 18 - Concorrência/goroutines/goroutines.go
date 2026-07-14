package main

import (
	"fmt"
	"time"
)

func main() {
	/* O GoRoutine funciona da seguinte forma: Quando eu passo o "Go" na linha, ele não espera a função terminar a sua execução. Ele joga a execução para o lado (em uma goroutine) e imediatamente passa para a linha seguinte. */
	go escrever("Olá mundo!") // goroutine
	escrever("Programando em Go!")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
