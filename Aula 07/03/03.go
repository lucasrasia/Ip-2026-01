package main
import "fmt"
func media(a, b, c float64)float64{
	return (a+b+c)/3
}
 func main(){
	var a, b, c float64
	fmt.Print("Digíte 3 valores: ")
	fmt.Scan(&a, &b, &c)
	m:=media(a, b, c)
	fmt.Printf("A média é: %.2f\n", m)
 }
