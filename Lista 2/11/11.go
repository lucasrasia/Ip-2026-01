package main
import "fmt"

func main() {
	var x float64
	fmt.Print("Digite o valor de x: ")
	if x == 2 {
		fmt.Println("Erro: divisao por zero, x nao pode ser 2")
	} else {
	f := 8.0 / (2.0 - x)
	fmt.Printf("f(%.2f) = %.6f\n", x, f)
	}	
}
