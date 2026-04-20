package main
import (
	"fmt"
	"math"
)
func main(){
	var r float64
	for i:=0; i<64; i++{
		r+=math.Pow(2, float64(i))
	}
	fmt.Println("Resultado:", r)
}