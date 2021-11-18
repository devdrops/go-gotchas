package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.X * p.X + p.Y * p.Y)
}

func main() {
	/*
	O código abaixo causa panic:
	```
	panic: runtime error: invalid memory address or nil pointer dereference
	```
	var p *Point
	fmt.Println(p.Abs())

	Por quê?

	O ponteiro p está declarado mas não está inicializado, logo é nil. Daí o
	panic.

	Para corrigir isso, basta inicializar criando algo do tipo Point, usando
	a função new() por exemplo, que faz a alocação de memória retornando o
	ponteiro (valor do endereço de memória) do tipo então criado.
	*/
	// Criando através da função new()
	var p *Point = new(Point)
	fmt.Println(p.Abs())

	/*
	Referências de leitura:

	- https://pkg.go.dev/builtin#new
	- https://yourbasic.org/golang/pointers-explained/
	*/
}
