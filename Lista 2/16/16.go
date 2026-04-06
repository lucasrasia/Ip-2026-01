package main
import (
	"fmt"
	"math"
)
var a, b, c float64
var delta float64
func bask1(a, b, c float64)float64{
	return (-b + math.Sqrt(delta)) / (2 * a)
}
func bask2(a, b, c float64)float64{
	return (-b - math.Sqrt(delta)) / (2 * a)
}
func main(){
	fmt.Print("Valores de A, B e C: ")
	fmt.Scan(&a, &b, &c)
	delta = b*b - 4*a*c
	b1:= bask1(a, b, c)
	b2:= bask2(a, b, c)
	if delta<0{
		fmt.Println("As raízes sao imagnárias, não há como expressar com números reais")
	} else if delta==0{
		fmt.Printf("Ás raízes são iguais: %.2f \n", b1)
	} else{
		fmt.Printf("Ás raízes são: %.2f e %.2f \n", b1, b2)
	}
}
