package main
import (
	"fmt"
	"math"
)
func vol(r float64)float64{
	pi:=3.14159265
	return (4/3)*pi*math.Pow(r, 3)
}
func main(){
	for r:=0.5; r<=20; r+=0.5{
		fmt.Printf("Volume esfera de raio %v: %.2f\n\n", r, vol(r))
	}

}