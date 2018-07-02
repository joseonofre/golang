package main

import (
	"fmt"

	"github.com/joseonofre/golang/interfaces/model"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da Silva"

	queroAcordarComUmCacareja(jojo)
	queroOuvirUmaPataNoLago(jojo)
}

func queroAcordarComUmCacareja(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func queroOuvirUmaPataNoLago(p model.Pata) {
	fmt.Println(p.Quack())
}
