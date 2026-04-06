package main

import "fmt"

func main() {
	var preco float64
	var codigo int
	var precoFinal float64

	fmt.Print("Digite o preço normal: ")
	fmt.Scan(&preco)
	fmt.Print("Digite o código da condição (1-4): ")
	fmt.Scan(&codigo)

	if codigo == 1 {
		precoFinal = preco * 0.9
	} else if codigo == 2 {
		precoFinal = preco * 0.95
	} else if codigo == 3 {
		precoFinal = preco
	} else if codigo == 4 {
		precoFinal = preco * 1.1
	} else {
		fmt.Println("Código inválido")
		return
	}

	fmt.Printf("Preço final: R$ %.2f\n", precoFinal)
}