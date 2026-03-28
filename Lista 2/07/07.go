package main
import "fmt"

func main() {
	var a, b, c int	
	fmt.Print("Digite 3 valores inteiros: ")
	fmt.Scan(&a, &b, &c)
	mi:=min(a, b, c)
	ma:=max(a, b, c)
	fmt.Println("Mínimo:", mi)
	i:= a+b+c-mi-ma
	fmt.Println("Intermediário:", i)
	fmt.Println("Máximo:", ma)
}
