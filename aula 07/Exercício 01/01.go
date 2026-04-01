package main
import"fmt"
func soma(a, b float64)float64{ //soma é o nome, no return eu defino o que a função vai fazer
		return a+b
	}
func main() {
	var a, b float64  //precisa definir a variavel
	fmt.Print("Valor 1: ")
	fmt.Scan(&a)
	fmt.Print("Valor 2: ")
	fmt.Scan(&b)
	s:=soma(a, b)
	fmt.Println("A soma é: ", s)
}