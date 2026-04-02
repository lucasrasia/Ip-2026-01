package main
import "fmt"
func main(){
	var n int
	fmt.Print("Número de 3 dígitos: ")
	fmt.Scan(&n)
	if n<100 || n>999 {
		fmt.Println("Erro! Dígite um número de 3 dígitos")
	} else {
		n1:=n/10
		n2:= n1%10
		fmt.Println(n2)
	}
}
