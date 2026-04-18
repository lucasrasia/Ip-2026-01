package main
import "fmt"
func main(){
	var a float64
	var b float64
	for i:=1; i<=37; i++{
		a=38
		b+=(a*(a+1))/float64(i)
		a=a-1
	}
	fmt.Println(b)
}