package main

import (
	"fmt"
)

func Foo(a [2]int) {
	a[0] = 8
}

func OtoFoo(a []int) {
	if len(a) > 0 {
		a[0] = 8
	}
}

func main() {
	/*
	Este aqui tem um detalhe muito interessante.

	A princípio, olhei para o código e pensei "pô, mas não vai atualizar de
	jeito nenhum assim porque estamos passando uma cópia e não o mesmo
	array. E de fato é isso o motivo do array abaixo não ser atualizado:
	*/

	a := [2]int{1, 2}
	Foo(a)
	fmt.Println(a)

	/*
	Numa explicação mais detalhada: a função Foo espera um array, e, ao
	chamá-la, passo a variável a. Só que, por questões de escopo (que são
	comuns em muitas linguagens de programação), a operação aqui trabalha
	com os dados dentro do seu escopo. Ou seja, dentro de Foo, só se sabe o
	que está dentro de Foo e ponto final, assim como dentro de main etc.
	Logo, da forma como o código foi escrito, o array nunca será atualizado
	pela função Foo.

	Podemos resolver isso passando um ponteiro do array para a função Foo,
	por exemplo. Aí sim, o escopo de Foo passa a receber um valor de fora,
	e, ao fazer sua operação, temos de fato a atualização do array.

	func Foo(a *[2]int) {
		a[0] = 8
	}

	a := [2]int{1, 2}
	Foo(&a)
	fmt.Println(a) // [8 2]

	MAS, uma coisa que o texto diz e que eu não sabia, é usar um slice ao
	invés de um array.
	*/

	s := []int{1, 2}
	OtoFoo(s)
	fmt.Println(s) // [8 2]

	/*
	Mas, POR QUÊ?

	Segundo o texto original, um slice não armazena dados, apenas descreve
	uma parte de um array. Isso quer dizer que ao passarmos um slice para
	uma função, estamos na verdade passando a referência do segmento de um
	array?

	Parece que sim, e isso eu não sabia. Ao mudar o valor de um elemento num
	slice, estamos mudando o valor do array, e outros slices deste mesmo
	array vão ter a mesma atualização. Malandro, que pegadinha sinistra!
	*/

	/*
	Referências de leitura:

	- https://yourbasic.org/golang/slices-explained/
	*/
}
