package main
import (
	"fmt"
	"math"
)
func serie(a float64)float64{
	return a - math.Pow(a, 3)/6 + math.Pow(a, 5)/120 - math.Pow(a, 7)/5040
}
func main(){
	var a float64
	fmt.Print("Digite o valor do ângulo em radianos: ")
	fmt.Scan(&a)
	if a<0 || a>6.3{
		fmt.Print("Digite um valor entre 0 e 2pi")
	} else{
		sena:=serie(a)
		fmt.Printf("Sen(%v)=%.2f\n", a, sena)
	}
}