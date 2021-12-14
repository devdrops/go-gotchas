package main

import (
	"fmt"
	"math/big"
)

func main() {
	/*
	O código abaixo gera o seguinte erro:

	const n = 9876543210 * 9876543210
	fmt.Println(n)

	constant 97546105778997104100 overflows int

	O que houve?

	Este cálculo é simples, mas há um detalhe. A constante n é declarada
	sem o tipo de forma explícita, logo o tipo é inferido. assumindo int.
	Porém, a conversão é necessária antes da chamada à fmt.Println() para
	que a conversão para o tipo esperado (interface{}) funcione
	corretamente. No final, o resultado da multiplicação excede o tamanho
	máximo do tipo int, o que causa o problema.

	Isso pode ser resolvido aplicando um tipo à constante, como float64 por
	exemplo. No entanto, é válido aqui explorar o conteúdo do package
	math/big para operar melhor com esse tipo de valores.
	*/

	v1 := big.NewInt(9876543210)
	v2 := big.NewInt(9876543210)
	on := big.NewInt(0)
	fmt.Println(on.Mul(v1, v2)) // 97546105778997104100
	/*
	Referências de leitura:

	- https://pkg.go.dev/math/big
	*/
}
