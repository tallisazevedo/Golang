package main

import "fmt"

// Podemos utilizar as interfaces como um tipo genérico. A função aceitará qualquer tipo
func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	generica("String")
	generica(1)
	generica(true)
}
