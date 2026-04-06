package main

import "fmt"

func main() {
	var preco float64
	var categoria string
	var dia int
	var precoFinal float64

	fmt.Print("Digite o preço normal do DVD: ")
	fmt.Scan(&preco)
	fmt.Print("Digite a categoria (comum/lancamento): ")
	fmt.Scan(&categoria)
	fmt.Print("Digite o dia da semana (1-domingo, 2-segunda, ..., 7-sábado): ")
	fmt.Scan(&dia)

	if categoria == "lancamento" {
		preco *= 1.15
	}

	if dia == 2 || dia == 3 || dia == 5 {
		precoFinal = preco * 0.6
	} else {
		precoFinal = preco
	}

	fmt.Printf("Preço final: R$ %.2f\n", precoFinal)
}