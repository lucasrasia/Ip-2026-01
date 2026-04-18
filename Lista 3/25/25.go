package main
import (
	"fmt"
	"math"
)

func main(){
	var r float64
	b:=15.0
	for i:=0; i<=14; i++{ 
		if i%2==0{
			r=r+(math.Pow(2, float64(i)))/math.Pow(b, 2)
		} else{
			r=r-(math.Pow(2, float64(i)))/math.Pow(b, 2)
		}
		b=b-1
	}
	fmt.Println(r)
}