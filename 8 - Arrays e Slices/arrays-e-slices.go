package main

import "fmt"

func main() {
	fmt.Println("Array e Slices")

	/* ARRAYS */
	var arr1 [5]string

	arr1[0] = "Primeira posição"
	arr1[1] = "Segunda posição"
	arr1[2] = "Terceira posição"

	arr2 := [5]string{"Posição 1", "Posição 2", "Posição 3", "Posição 4", "Posição 5"}

	fmt.Println(arr1)
	fmt.Println(arr2)

	/* SLICES */
	slice := []string{"Tallis", "Gigi", "Binho", "Mirela", "Gedson"}
	fmt.Println(slice)

	slice = append(slice, "Rafaela") // Adicionando a string "Rafaela" no meu slice

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i]) // Exibindo cada item do meu slice individualmente
	}

	slice2 := arr2[1:3] // Estou pegando o indice 1 e o indice 3 do arr2 (Funciona como um ponteiro, ele está apontando para os dois indice. Caso eu altere o conteúdo armazenado no indice, também mudará no slice)
	fmt.Println(slice2)

	arr2[1] = "Posição alterada"
	fmt.Println(slice2)

}
