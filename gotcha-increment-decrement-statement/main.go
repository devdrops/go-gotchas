package main

import (
	"fmt"
)

func main() {
	/*
	O trecho de código abaixo gera dois erros de sintaxe:

	i := 0
	fmt.Println(++i)
	fmt.Println(i++)

	syntax error: unexpected ++, expecting expression
	syntax error: unexpected ++, expecting comma or )

	Por quê?

	Em Go, as operações de incremento e decremento não podem ser usadas
	como expressões (algo passado para uma função), apenas como declarações.
	E também, só funcionam se for pós (o uso como pré também cai no mesmo
	erro de sintaxe), como no trecho de código abaixo:

	i := 0
	++i
	fmt.Println(i)

	syntax error: unexpected ++, expecting }

	A forma de resolver isso é fazer cada etapa de uma vez: primeiro a
	operação (incremento ou decremento), depois a próxima tarefa:
	*/
	i := 0
	i++
	fmt.Println(i) // 1
	i--
	fmt.Println(i) // 0
	/*
	Referências de leitura:

	- https://go.dev/doc/faq#inc_dec
	*/
}
