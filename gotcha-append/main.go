package main

import (
	"fmt"
)

func main() {
	/*
	O texto original aponta que o trecho de código abaixo resulta em um bug
	bizarro:

	a := []byte("ba")

	a1 := append(a, 'd')
	a2 := append(a, 'g')

	fmt.Println(string(a1)) // bag
	fmt.Println(string(a2)) // bag

	Ambos os slices escrevem a mesma saída, devido a um bug na forma como a
	função append funciona (ou melhor, funcionava), onde a, a1 e a2 se
	referiam ao mesmo array. Logo, ao editar um desses slices, todos seriam
	editados.

	Testei aqui a execução do mesmo código com Docker, em imagens com Linux
	Alpine como base, da versão 1.5 até 1.17, e não consegui reproduzir esse
	bug a partir da versão 1.12. Estranhamente, a saída correta obtida a
	partir de 1.12 foi reproduzida na versão 1.6. Curioso, hehehe.

	A solução apontada para o problema do texto original é usar um array
	diferente para cada ação de append():
	*/
	const prefix = "ba"

	a1 := append([]byte(prefix), 'd')
	a2 := append([]byte(prefix), 'g')

	fmt.Println(string(a1)) // bad
	fmt.Println(string(a2)) // bag

	// Assim como no outro caso, também escrevi sobre este no Twitter
	// marcando o autor do blog.

	/*
	Referências de leitura:

	- https://yourbasic.org/golang/append-explained/
	- https://twitter.com/devdrops/status/1463529279602540544
	*/
}
