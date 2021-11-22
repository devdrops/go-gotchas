package main

import (
	"fmt"
)

func main() {
	/*
	O código abaixo vai gerar um erro de sintaxe:

	```
	fruit := []string{
		"apple",
		"banana",
		"cherry"
	}
	fmt.Println(fruit)
	```

	```
	syntax error: unexpected newline, expecting comma or }
	```

	A explicação, para quem não conhece: em slices, arrays e maps onde
	escrevemos seus elementos em mais de uma linha, cada linha de um
	elemento deve terminar com uma vírgula, como o código abaixo:
	*/

	fruit := []string{
		"apple",
		"banana",
		"cherry", // olha a vírgula aqui no final
	}
	fmt.Println(fruit)

	/*
	Referências de leitura:

	- https://go.dev/ref/spec#Semicolons
	*/
}
