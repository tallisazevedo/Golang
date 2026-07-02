package main

import (
	"fmt"
	"math"
)

// interface define um conjunto de "comportamentos" (métodos) que um tipo precisa ter, sem dizer como esse tipo vai fazer isso.
type forma interface {
	area() float64
}

func escreverArea(f forma) {
	fmt.Printf("A area da forma é %0.2f\n", f.area())
}

type retangulo struct {
	largura float64
	altura  float64
}

// Implementação da interface
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct {
	raio float64
}

// Implementação da interface
func (c circulo) area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}

func main() {

	r := retangulo{32.2, 33.3}

	escreverArea(r)

	c := circulo{10}

	escreverArea(c)
}
