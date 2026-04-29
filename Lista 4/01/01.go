package main
import "fmt"
func main(){
	var n[]int
	var a int
	for i:=1; i<=10; i++{
		fmt.Printf("Digite o %v valor: ", i)
		fmt.Scan(&a)
		n=append(n, a)
	}
	var contador int
	for v:=range n {
		if n[v]>50{
			fmt.Printf("%v posição %v\n", n[v], v+1)
			contador++
		}
	}
	if contador==0{
		fmt.Println("Não foi digiado nenhum número maior que 50")
	}
}