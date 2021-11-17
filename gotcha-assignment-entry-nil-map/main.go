package main

/*
A declaração abaixo resulta em panic:

```
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()
        /code/main.go:5 +0x2e
        exit status 2
```

Por quê?
Temos a variável criada na declaração `var`, temos a atribuição de valor. O que
falta aqui é o uso de **make**.

Todo map precisa ser inicializado, seja com a função make ou um map literal
antes de inserir elementos. Todo map tem nil por padrão até ser inicializado, e
enquanto for assim, elementos não podem ser inseridos nele.

Referência de leitura:

- https://yourbasic.org/golang/maps-explained/
*/

func main() {
	// Código original:
	//var m map[string]float64
	//m["pi"] = 3.1416

	// Código com uso de make:
	m := make(map[string]float64)
	m["pi"] = 3.1416
}
