package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	numero := 3

	fmt.Print("O numero ", numero, " se escreve assim: ")
	switch numero {
	case 1:
		fmt.Println("Um.")
	case 2:
		fmt.Println("Dois.")
	case 3:
		fmt.Println("Três.")
	}

	fmt.Println("Você é da família do Unix?")
	switch runtime.GOOS {
	case "darwin":
		fallthrough
	case "freebsd":
		fallthrough
	case "linux":
		fmt.Printf("Sim, eu sou um %s \r\n", runtime.GOOS)
	default:
		fmt.Println("Deixa essa pergunta para lá")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Ok, você pode dormi até mais tarde")
	default:
		fmt.Println("Vamos lá, é dia de trabalhar")
	}

	numero = 10

	switch {
	case numero < 10:
		fmt.Println("Sim")
	case numero >= 10 && numero < 100:
		fmt.Println("serve 2 dígitos")
	case numero >= 100:
		fmt.Println("O número é muito grande")

	}

}
