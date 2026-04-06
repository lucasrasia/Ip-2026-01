package main

import "fmt"

func main() {
	var idade int
	var classe string

	fmt.Print("Digite a idade: ")
	fmt.Scan(&idade)

	if idade < 16 {
		classe = "Não-eleitor"
	} else if idade >= 18 && idade <= 65 {
		classe = "Eleitor obrigatório"
	} else {
		classe = "Eleitor facultativo"
	}

	fmt.Printf("Classe eleitoral: %s\n", classe)
}
