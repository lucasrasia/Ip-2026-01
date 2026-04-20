package main
import (
	"fmt"
	"strconv"
)

func main() {
	var a string
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&a)
	n, err:=strconv.Atoi(a)
	if err!=nil{
		fmt.Println("O número deve ser inteiro")
		return
	}
	if n<=0{
		fmt.Println("O número deve ser positivo")
		return
	}
	var soma int
	for i := 1; i <= n; i++ {
		soma+=i
	}
	fmt.Printf("A soma é %v\n", soma)
}
