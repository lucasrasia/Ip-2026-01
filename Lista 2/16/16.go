package main
import (
	"fmt"
	"math"
)
func bask1(a, b, c float64)float64{
	return (math.Sqrt(b*b-4*a*c)-b)/2*a
}
func baks2(a, b, c float64)float64{
	return (-b-(math.Sqrt(b*b-4*a*c)))/2*a
}
func main(){
	var a, b, c float64
	fmt.Print("Valores de A, B e C: ")
	fmt.Scan(&a, &b, &c)
}
