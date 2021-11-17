package main

import "fmt"

func main() {
	/*
	A declaração abaixo resulta em panic:
	```
	panic: assignment to entry in nil map
	```
	var m map[string]float64
	m["pi"] = 3.1416

	Por quê?
	Tem a variável criada na declaração `var`, tem a atribuição de valor.
	O que falta aqui é a inicialização do map.

	Todo map precisa ser inicializado, seja com a função make ou um map
	literal antes de inserir elementos. Todo map tem nil por padrão até ser
	inicializado, e enquanto for assim, elementos não podem ser inseridos
	nele.
	*/

	// map inicializado com uso de make:
	m := make(map[string]float64)
	m["pi"] = 3.1416

	// map inicializado de forma literal:
	var om = map[string]float64{
		"pi": 3.1416,
		"aaa": 1.2345,
	}
	fmt.Println(om)

	/*
	Referência de leitura:

	- https://yourbasic.org/golang/maps-explained/
	*/
}
