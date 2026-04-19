package main
import (
	"fmt"
	"math"
)
func main(){
	var s float64
	for i:=0; i<=50; i++{
		impar:=2*i+1
		if i%2==0{
			s=s+(1.0/math.Pow(float64(impar), 3))
		} else{
			s=s-(1.0/math.Pow(float64(impar), 3))
		}
	}
	pi:=math.Cbrt((32*s))
	fmt.Printf("%.5f\n", pi) 	
}