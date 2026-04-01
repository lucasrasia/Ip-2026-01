package main
import"fmt"

func maior(a, b, c int)int{
	return max(a, b, c)
}
func main(){
	var a, b, c int
	fmt.Print("Digíte 3 valores: ")
	fmt.Scan(&a, &b, &c)
	m:=maior(a, b, c)
	fmt.Printf("O maior é %v\n", m)
}
