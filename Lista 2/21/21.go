package main

import "fmt"

func main() {
	var id int
	var nota1, nota2, nota3, mediaEx float64
	var mediaFinal float64
	var conceito string
	var status string

	fmt.Print("Digite o número de identificação: ")
	fmt.Scan(&id)
	fmt.Print("Digite as 3 notas: ")
	fmt.Scan(&nota1, &nota2, &nota3)
	fmt.Print("Digite a média dos exercícios: ")
	fmt.Scan(&mediaEx)

	mediaFinal = (nota1 + nota2*2 + nota3*3 + mediaEx) / 7

	if mediaFinal >= 9.0 && mediaFinal <= 10.0 {
		conceito = "A"
	} else if mediaFinal >= 7.5 {
		conceito = "B"
	} else if mediaFinal >= 6.0 {
		conceito = "C"
	} else if mediaFinal >= 4.0 {
		conceito = "D"
	} else {
		conceito = "E"
	}

	if conceito == "A" || conceito == "B" || conceito == "C" {
		status = "APROVADO"
	} else {
		status = "REPROVADO"
	}

	fmt.Printf("Número: %d\nNotas: %.2f, %.2f, %.2f\nMédia exercícios: %.2f\nMédia aproveitamento: %.2f\nConceito: %s\n%s\n", id, nota1, nota2, nota3, mediaEx, mediaFinal, conceito, status)
}