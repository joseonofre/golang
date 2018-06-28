package main

import (
	"fmt"
)

var (
	//Endereco é um valor importante e tem que ser publico
	Endereco string //Variáveis com começo em maiúsculo são variáveis públicas
	//Telefone é um valor importante para nossa aula
	Telefone   string
	quantidade int64
	comprou    bool
	valor      float64
	palavras   rune
)

func main() {
	teste := "valor de teste"
	fmt.Printf("endereco: %s \r\n", Endereco)
	fmt.Printf("quantidade: %d \r\n", quantidade)
	fmt.Printf("comprou: %v \r\n", valor)
	fmt.Printf("palavras: %v \r\n", palavras)
	fmt.Printf("teste: %v \r\n", teste)
}
