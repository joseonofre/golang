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

	casa := new(Imovel)
	fmt.Printf("Ponteiro da casa: %p - Casa: %+v \r\n", &casa, casa)

	chacara := Imovel{10, 10, "Chácara Linda", 100000}
	fmt.Printf("Ponteiro da chacara: %p - Chacara %+v \r\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("Ponteiro da chacara: %p - Chacara %+v \r\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X * 10
}
