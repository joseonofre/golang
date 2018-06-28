package matematica

//Calculo execulta qualquer tipo de calculo, basta enviar a função desejada
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

// Multiplicador multiplica x * y
func Multiplicador(x int, y int) int {
	return x * y
}

//Divisor é a divisão de 2 números
func Divisor(numeroA int, numeroB int) (resultado int) {
	resultado = numeroA / numeroB
	return
}

// DivisorComResto retorna um inteiro com a divisao e outro com o resto
func DivisorComResto(numeroA int, numeroB int) (resultado int, resto int) {
	resultado = Divisor(numeroA, numeroB)
	resto = numeroA % numeroB
	return
}
