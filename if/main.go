package main

import (
	"fmt"
)

func main() {
	meses := 6
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Este credor deve a pouco tempo")
	}

	if situacao {
		fmt.Println("Ele está devendo")
	}

	if !situacao {
		fmt.Println("Ele está adiplente")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive na cidade teste")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Printf("Qual a situação do cliente? %s \r\n", descricao)
		return
	}

	fmt.Println("Obrigado por nos consultar")
}

//retorno assinado
func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo"
		return
	}
	status = false
	descricao = "O cliente está com o carnê em dia"
	return
}
