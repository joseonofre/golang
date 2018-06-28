package model

//Imovel é uma struct que armazena os dados de um imovel
type Imovel struct {
	X     int    `json:"cordenada_x"`
	Y     int    `json:"cordenada_y"`
	Nome  string `json:"nome"`
	valor int
}

// SetValor atribui o valor do imóvel
func (i *Imovel) SetValor(valor int) {
	i.valor = valor
}

// GetValor retorna o valor do imóvel
func (i *Imovel) GetValor() int {
	return i.valor
}
