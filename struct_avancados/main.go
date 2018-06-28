package main

import (
	"encoding/json"
	"fmt"

	"github.com/joseonofre/struct_avancados/model"
)

func main() {
	casa := model.Imovel{}
	casa.X = 200
	casa.Y = 10
	casa.Nome = "Minha Casa"
	casa.SetValor(10000)

	fmt.Printf("Casa: %+v \r\n", casa)
	fmt.Printf("O valor da casa é: %d \r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	fmt.Println("Casa em JSON:", string(objJSON))
}
