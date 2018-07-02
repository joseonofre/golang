package main

import (
	"fmt"
)

func main() {
	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Printf("Qual capacidade deste array? %d \r\n", len(teste))

	cantores := [2]string{"José Onofre", "Clarice Garcia"}
	fmt.Printf("O que há neste array: \n\r%+v\r\n", cantores)

	capitais := [...]string{"Lisboa", "Brasília", "Luanda", "Maputo"}
	fmt.Printf("Qual capacidade deste array? %d \r\n", len(capitais))

	for indice, cidade := range capitais {
		fmt.Printf("Capital[%d] é: %s\r\n", indice, cidade)
	}
}
