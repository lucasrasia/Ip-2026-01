package main

import "fmt"

func main() {
	var conta string
	var tipo string
	var consumo float64
	var valor float64

	fmt.Print("Digite a conta do cliente: ")
	fmt.Scan(&conta)
	fmt.Print("Digite o tipo (residencial/comercial/industrial): ")
	fmt.Scan(&tipo)
	fmt.Print("Digite o consumo em m³: ")
	fmt.Scan(&consumo)

	if tipo == "residencial" {
		valor = 5.0 + 0.05*consumo
	} else if tipo == "comercial" {
		if consumo <= 80 {
			valor = 500.0
		} else {
			valor = 500.0 + 0.25*(consumo-80)
		}
	} else if tipo == "industrial" {
		if consumo <= 100 {
			valor = 800.0
		} else {
			valor = 800.0 + 0.04*(consumo-100)
		}
	} else {
		fmt.Println("Tipo inválido")
		return
	}

	fmt.Printf("Conta: %s\nValor a pagar: R$ %.2f\n", conta, valor)
}
