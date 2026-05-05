package main
import(
	"fmt"
	"math"
)
func main(){
	var A[]float64
	var a float64
	for i:=1;i<=15;i++{
		fmt.Printf("Valor %v: ", i)
		fmt.Scan(&a)
		if a<0{
			a=-1
			A=append(A, a)
			continue
		}
		r:=math.Sqrt(a)
		A=append(A, r)
	}
	for _, v := range A {
    fmt.Printf("%.2f\n", v)
	}
}