package operadora

import (
	"strconv"

	"github.com/joseonofre/pacotes/prefixo"
)

//NomeOperadora representa o nome da operadora que vamos usar
var NomeOperadora = "XPTO Telecon"

//PrefixoDaCapitalOperadora prefixo mais o nome da operadora
var PrefixoDaCapitalOperadora = strconv.Itoa(prefixo.Capital) + " " + NomeOperadora
