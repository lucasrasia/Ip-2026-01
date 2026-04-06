package main

import (
	"fmt"
	"math"
)

func main() {
	var opcao int
	var raio, altura float64
	var volume, area float64

	fmt.Print("Digite a opção (1-cone, 2-cilindro, 3-esfera): ")
	fmt.Scan(&opcao)

	if opcao == 1 {
		fmt.Print("Digite o raio e a altura do cone: ")
		fmt.Scan(&raio, &altura)
		volume = (1.0 / 3.0) * math.Pi * raio * raio * altura
		area = math.Pi * raio * math.Sqrt(raio*raio + altura*altura)
		fmt.Printf("Volume: %.2f\nÁrea: %.2f\n", volume, area)
	} else if opcao == 2 {
		fmt.Print("Digite o raio e a altura do cilindro: ")
		fmt.Scan(&raio, &altura)
		volume = math.Pi * raio * raio * altura
		area = 2 * math.Pi * raio * altura
		fmt.Printf("Volume: %.2f\nÁrea: %.2f\n", volume, area)
	} else if opcao == 3 {
		fmt.Print("Digite o raio da esfera: ")
		fmt.Scan(&raio)
		volume = (4.0 / 3.0) * math.Pi * raio * raio * raio
		area = 4 * math.Pi * raio * raio
		fmt.Printf("Volume: %.2f\nÁrea: %.2f\n", volume, area)
	} else {
		fmt.Println("Opção inválida")
	}
}