package main

import (
	"fmt"
)

func main() {
	/*
	O trecho de código abaixo tem um problema: queremos copiar os elementos
	de src para dst, porém, quando o fazemos e escrevemos o conteúdo de dst,
	o slice aparece como vazio, sem os elementos copiados.

	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)
	fmt.Println("dst: ", dst)

	O que houve aqui?
	A função copy exige que a variável de destino tenha no mínimo o mesmo
	tamanho da variável de origem. Como a variável dst foi apenas criada,
	sem elementos inseridos nem tamanho definido, a operação de cópia acaba
	não acontecendo.

	Para corrigir isso, podemos usar a função make para "inicializar" a
	variável dst com a mesma quantidade de elementos que src tem:
	*/
	var src, dst []int
	src = []int{1, 2, 3}
	dst = make([]int, len(src))
	copy(dst, src)
	fmt.Println("dst: ", dst)
	fmt.Println("cap dst: ", cap(dst))
	/*
	Podemos também usar a função append para chegar no mesmo resultado. A
	diferença aqui é que append pode criar um slice com a capacidade maior
	de elementos do que len(osrc).
	*/
	var osrc, odst []int
	osrc = []int{1, 2, 3}
	odst = append(odst, osrc...)
	fmt.Println("odst: ", odst)
	fmt.Println("cap odst: ", cap(odst))
	/*
	Referências de leitura:

	- https://yourbasic.org/golang/copy-explained/
	- https://pkg.go.dev/builtin#append
	- https://pkg.go.dev/builtin#copy
	*/
}
