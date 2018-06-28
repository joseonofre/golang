package main

import (
	"fmt"

	"github.com/joseonofre/function/matematica"
)

func main() {
	resultado := matematica.Calculo(matematica.Multiplicador, 2, 2)
	fmt.Printf("O Resultado da multiplicacao é: %d \r\n", resultado)
	resultado = matematica.Calculo(matematica.Soma, 10, 10)
	fmt.Printf("O Resultado da soma é: %d \r\n", resultado)
	resultado = matematica.Calculo(matematica.Divisor, 20, 2)
	fmt.Printf("O Resultado da divisão é: %d \r\n", resultado)
	resultado, resto := matematica.DivisorComResto(20, 10)
	fmt.Printf("O resultado com resto foi: %d e o resto da divisão foi: %d \r\n", resultado, resto)
}
