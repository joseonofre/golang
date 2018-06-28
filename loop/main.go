package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("Qual o valor de i? %d \r\n", i)
	}

	valor := 0
	teste := true
	for teste {
		valor++
		if valor%5 == 0 {
			teste = false
		}
		fmt.Println("O número é: ", valor)
	}

	for {
		valor--
		if valor == 0 {
			break
		}
		fmt.Println("O número é: ", valor)
	}

	texto := "Eu adoro escrever programas em go"
	for indice, letra := range texto {
		fmt.Printf("Texto[%d] = %q \r\n", indice, letra)
	}
}
