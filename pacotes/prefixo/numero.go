package prefixo

import (
	"strconv"
)

//Capital representa o numero do prefixo de telefone da capital de um estado/provincia
var Capital = 11

var teste = "teste"

//TesteComPrefixo é um teste de private com public
var TesteComPrefixo = teste + " " + strconv.Itoa(Capital)
