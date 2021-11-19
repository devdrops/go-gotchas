package main

import (
	"fmt"
	"time"
)

func main() {
	/*
	Aqui temos um detalhe.

	O texto original deste desafio explica que o código abaixo resulta no
	seguinte erro de compilação:

	```
	multiple-value time.Parse() in single-value context
	```

	Este é o código proposto:
	*/

	t := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
	fmt.Println(t)

	/*
	O texto segue explicando um detalhe comum, a confusão de tentar atribuir
	a somente uma variável o retorno de dois valores distintos. O método
	time.Parse() retorna um Time e um error. Logo, os caminhos aqui podem
	ser os seguintes:

	// Analisando o retorno de erro
	t, err := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(t)

	// Ignorando o retorno de erro com _ (underscore)
	t, _ := time.Parse(time.RFC3339, "2018-04-06T10:49:05Z")
	fmt.Println(t)

	O que chamou a atenção é que a mensagem que obtive foi diferente da que
	é descrita no texto. Aqui, ao executar o código conforme o texto
	original, recebi a seguinte mensagem:

	```
	assignment mismatch: 1 variable but time.Parse returns 2 values
	```

	A versão de Go que usei: go version go1.17.2 linux/amd64.

	Ou seja, acredito que a particularidade aqui é que a mensagem era fora
	do que conhecemos, um pouco estranha mas que podemos entender. O
	importante é notar que essa mensagem não existe mais (até avisei o
	autor dos desafios no Twitter, para ver se podemos atualizar a lista de
	desafios).

	*/

	/*
	Referências de leitura:

	- https://yourbasic.org/golang/underscore/
	- https://pkg.go.dev/time#Parse
	- https://twitter.com/devdrops/status/1461713172562382848
	*/
}
