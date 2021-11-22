package main

import (
	"fmt"
)

func main() {
	/*
	O código abaixo resulta em um erro:

	```
	s := "hello"
	s[0] = "H"
	fmt.Println(s)
	```

	```
	cannot assign to s[0] (strings are immutable)
	```

	A mensagem é direta: strings são imutáveis, ou seja, não podemos mudar
	seus valores da forma como está ali no código.

	Para mudar o valor de s, devemos usar outro tipo de dados, rune, como o
	exemplo abaixo:
	*/
	buf := []rune("hello")
	/*
	Temos que usar aspas simples para delimitar a string aqui; do contrário,
	teremos outro erro, `cannot use "H" (type untyped string) as type rune
	in assignment`, dado que a string delimitada por aspas duplas é
	considerada uma constante não-tipada (untyped constant), e seu tipo não
	é o mesmo que usamos em buf.

	Vale lembrar que constantes não-tipadas possuem sim um tipo definido,
	como todo valor em Go. A diferença é que seu tipo é oculto, não sendo o
	mesmo que rune. Daí o problema :)
	*/
	buf[0] = 'H'
	s := string(buf)
	fmt.Println(s)

	/*
	Referências de leitura:

	- https://yourbasic.org/golang/string-functions-reference-cheat-sheet/
	- https://pkg.go.dev/builtin#string
	- https://pkg.go.dev/builtin#rune
	*/
}
