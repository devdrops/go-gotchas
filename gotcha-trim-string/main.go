package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
	O trecho de código abaixo tem um resultado "inesperado":

	fmt.Println(strings.TrimRight("ABBA", "BA")) // ""

	É isso aí, uma string vazia. Por quê?
	As funções Trim, TrimLeft e TrimRight removem todos os caracteres cujos
	valores de Unicode se encontram no trecho dado. No exemplo dado, todos
	os A:s e B:s são afetados, deixando apenas a string vazia no final.

	Para resolver isso, podemos usar TrimSuffix (que age no sufixo da
	string):
	*/
	fmt.Println(strings.TrimSuffix("ABBA", "BA")) // "AB"
	/*
	Referências de leitura:

	- https://pkg.go.dev/strings#Trim
	- https://pkg.go.dev/strings#TrimLeft
	- https://pkg.go.dev/strings#TrimRight
	- https://pkg.go.dev/strings#TrimSuffix
	- https://yourbasic.org/golang/string-functions-reference-cheat-sheet/
	*/
}
