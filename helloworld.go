package main

import (
	"fmt"
	"reflect"
)

func main() {
	nome := "Paulo"
	idade := 25
	versao := 1.1
	fmt.Println("Olá", nome, "sua idade é", idade)
	fmt.Println("Versão: ", versao)

	fmt.Print("O tipode var da versão é: ", reflect.TypeOf(versao))
}
