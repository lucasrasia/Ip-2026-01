package main

import "fmt"

func main() {
	var matricula, horasExtras int
	var salMin = 788.0
	var valorHE = 10.0
	var salHE, salBruto, inss, ir, salLiquido float64

	fmt.Print("Digite a matrícula: ")
	fmt.Scan(&matricula)
	fmt.Print("Digite as horas-extras: ")
	fmt.Scan(&horasExtras)

	salHE = float64(horasExtras) * valorHE
	salBruto = 3 * salMin + salHE

	inss = 0
	if salBruto > 1500 {
		inss = 0.12 * salBruto
	}

	ir = 0
	if salBruto > 2000 {
		ir = 0.20 * salBruto
	}

	salLiquido = salBruto - inss - ir

	fmt.Printf("Matrícula: %d\nSalário bruto: R$ %.2f\nSalário líquido: R$ %.2f\n", matricula, salBruto, salLiquido)
}