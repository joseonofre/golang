package main

import (
	"fmt"

	"github.com/joseonofre/pacotes/operadora"
	"github.com/joseonofre/pacotes/prefixo"
)

//NomeDoUsuario é público
var NomeDoUsuario = "Jose"

func main() {
	fmt.Printf("Nome do usuário: %s\r\n", NomeDoUsuario)
	fmt.Printf("Profixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Nome da operadora: %s \r\n", operadora.NomeOperadora)
	fmt.Printf("Nome da operadora: %s \r\n", prefixo.TesteComPrefixo)
}
