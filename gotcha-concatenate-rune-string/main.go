package main

import (
	"fmt"
)

func main() {
	/*
	O texto explica que o trecho de código abaixo tem resultados diferentes:

	fmt.Println("H" + "i") // Hi
	fmt.Println('H' + 'i') // 177

	A resposta está na diferença entre os tipos. Na primeira linha, temos a
	concatenação de dois valores do tipo untyped constant, com pseudo tipo
	string. Na segunda linha, temos a soma de dois valores também do tipo
	untyped constant, porém com pseudo tipo rune, que ali representam
	valores de Unicode ('H' é 72, e 'i' é 105), logo a operação não é uma
	concatenação mas sim uma adição, com o resultado abaixo.

	Uma das formas de resolver isso é converter os valores para string, como
	o código abaixo:
	*/
	fmt.Println("H" + "i") // Hi
	fmt.Println(string('H') + string('i')) // Hi
	fmt.Println(string(72) + string('i')) // Hi
	fmt.Println(string('H') + string(105)) // Hi
	fmt.Println(string(72) + string(105)) // Hi
	/*
	Também podemos usar a função fmt.Printf() (a opção %c escreve o
	caractere representado pelo valor Unicode):
	*/
	fmt.Printf("%c%c, world!\n", 'H', 'i') // Hi, world!
	fmt.Printf("%c%c, world!\n", 72, 'i') // Hi, world!
	fmt.Printf("%c%c, world!\n", 'H', 105) // Hi, world!
	fmt.Printf("%c%c, world!\n", 72, 105) // Hi, world!
	/*
	Referências de leitura:

	- https://pkg.go.dev/fmt#hdr-Printing
	- https://yourbasic.org/golang/fmt-printf-reference-cheat-sheet/
	*/
}
