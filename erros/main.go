package main

import (
	"encoding/json"
	"fmt"

	"github.com/joseonofre/golang/erros/model"
)

func main() {
	casa := model.Imovel{}
	casa.X = 200
	casa.Y = 10
	casa.Nome = "Casa amarela"
	casa.SetValor(10000)

	fmt.Printf("Casa é: %+v \r\n", casa)

	objJSON, err := json.Marshal(casa)

	if err != nil {
		fmt.Println("[main] Houve um erro na geracao do objeto JSON: ", err.Error())
	}
	fmt.Println("A casa em JSON ", string(objJSON))
}
