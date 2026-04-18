package main
import "fmt"
func main(){
	var b, n int
	fmt.Print("Digite a base: ")
	fmt.Scan(&b)
	fmt.Print("Digite o expoente: ")
	fmt.Scan(&n)
	r:=1
	for i:=1; i<=n; i++{
		r=r*b
	}
	fmt.Println("Resultado:", r)
}