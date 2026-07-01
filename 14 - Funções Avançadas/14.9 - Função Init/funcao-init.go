package main

import "fmt"

func init() {
	fmt.Println("Executando a função Init")
}

func main() {
	fmt.Println("Executando a função Main")
}

/*
	A função init será executada sempre antes da função "main". Diferente da função "main" que só pode ter uma por pacote. A função "init" pode ter uma por arquivo.
*/
