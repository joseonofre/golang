package main

import (
	"fmt"

	"github.com/joseonofre/golang/struct_avancados/model"
)

// os caches o go inicializa como null
var cache map[string]model.Imovel

func main() {

	cache = make(map[string]model.Imovel, 0)

	casa := model.Imovel{}
	casa.X = 200
	casa.Y = 10
	casa.Nome = "Casa amarela"
	casa.SetValor(10000)

	cache["Casa amarela"] = casa

	fmt.Println("Há uma Casa amarela no cache")
	imovel, achou := cache["Casa amarela"]

	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v - %v \r\n", imovel, achou)
	}

	apto := model.Imovel{}
	apto.Nome = "Apto Azul"
	apto.X = 19
	apto.Y = 26
	apto.SetValor(70000)

	cache[apto.Nome] = apto

	fmt.Println("Quantos itens há no chache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v \r\n", chave, imovel)
	}

	/*
		removendo itens de cache
	*/
	imovel, achou = cache["Casa amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}
	fmt.Println("Quantos itens há no chache?", len(cache))
	fmt.Printf("Ponteiro do cache: %p \r\n", &cache)
}
