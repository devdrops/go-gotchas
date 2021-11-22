package main

import (
	"fmt"
)

func main() {
	/*
	Esse é bem interessante.

	O código abaixo resulta em n mantendo o valor designado 0 (zero). A
	operação feita dentro do if não resulta em alteração do valor de n.

	```
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
	```

	Por quê?

	A declaração n := 1 faz a criação de uma nova variável n dentro do
	escopo do if, que acaba "obscurendo" (shadowing) a variável criada fora
	do if.

	Se queremos usar dentro do if a mesma variável n, devemos mudar o código
	para o seguinte:

	```
	n := 0
	if true {
		n = 1
		n++
	}
	fmt.Println(n)
	```

	Com o uso de =, estamos usando a mesma exata variável fora do if.

	E ainda: como podemos detectar a situação de variáveis "obscurecidas"?
	Podemos usar `go vet` para isso, da forma abaixo:

	Primeiro, instalamos a ferramenta necessária para a análise:
	```
	 go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
	```
	(a versão da ferramenta é exigida na instalação)


	Depois, usamos o comando go vet com a ferramenta instalada:
	```
	go vet -vettool=$(which shadow) main.go
	```

	Com isso, teremos a seguinte análise:
	```
	./main.go:66:3: declaration of "n" shadows declaration at line 63
	```
	*/
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)

	/*
	Referências de leitura:

	- go help vet
	- go doc cmd/vet
	- https://pkg.go.dev/cmd/go/internal/vet
	*/
}
