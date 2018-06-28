package main

import (
	"fmt"
)

//Imovel é uma struct que armazena os dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := Imovel{}
	fmt.Printf("A casa é %+v \r\n", casa)

	apartamento := Imovel{17, 56, "Meu Ap", 100000}
	fmt.Printf("O apartamento é %+v \r\n", apartamento)

	chacara := Imovel{
		Y:     10,
		X:     20,
		valor: 1000000,
		Nome:  "Minha chácara",
	}
	fmt.Printf("A chácara é: %+v \r\n", chacara)

	casa.Nome = "Minha casa"
	casa.valor = 40000
	casa.X = 20
	casa.Y = 400

	fmt.Printf("A casa com os dados %+v \r\n", casa)
}
