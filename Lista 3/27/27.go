package main
import(
	"fmt"
	"math"
)
func fat(x float64)float64{
	if x==1.0 || x==0.0{
		return 1
	}
	return x*fat(x-1.0)
}
func main(){
	var x float64
	fmt.Print("Digite o valor do ângulo: ")
	fmt.Scan(&x)
	var cosx float64
	for i:=0; i<=40; i=i+2{
		if i%4==0{
			cosx=cosx+(math.Pow(x, float64(i)))/fat(float64(i))
		} else{
			cosx=cosx-(math.Pow(x, float64(i)))/fat(float64(i))
		}
	}
	cosx2:=math.Cos(x)
	fmt.Println("Valor calculado:", cosx, "Valor pela função:", cosx2)
}