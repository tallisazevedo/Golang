package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps")

	// Criando um Map. Uma estrutura c/ chave e valor
	usuario := map[string]string{
		"nome":      "Pedro",
		"Sobrenome": "Silva",
	}
	fmt.Println(usuario)

	// Criando um Map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Tallis",
			"segundo":  "Azevedo",
		},
		"curso": {
			"nome":   "Ciência da Computação",
			"Campus": "UVV",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome") /* Método utilizando para excluir uma chave do MAP. Primeiro passamos qual é o Map e depois a chave */
	fmt.Println(usuario2)

	// Adicionando uma chave ao Map
	usuario2["signo"] = map[string]string{
		"nome": "aries",
	}
	fmt.Println(usuario2)
}
