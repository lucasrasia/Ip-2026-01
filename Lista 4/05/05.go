package main
import "fmt"
func main(){
	var A[]int
	var a int
	for i:=1;i<=10;i++{
		fmt.Printf("Valor %v: ", i)
		fmt.Scan(&a)
		A=append(A, a)
	}
	var min int
	var posicao int
	for i:=range A{
		if i==0{
			min=A[i]
		}
		if A[i]<min{
			min=A[i]
			posicao=i
		}
	}
	fmt.Println("O menor valor é", min, "e sua posição na lista é", posicao+1)
}
