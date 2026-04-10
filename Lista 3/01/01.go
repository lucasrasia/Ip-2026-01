package main
import "fmt"
func main(){
	var a, b int
	r:=1
	fmt.Print("Digíte a base: ")
	fmt.Scan(&a)
	fmt.Print("Dígite o expoente: ")
	fmt.Scan(&b)
	for i:=1; i<=b; i++{
		r=r*a
	}
	fmt.Printf("Resultado: %v\n", r)
}
