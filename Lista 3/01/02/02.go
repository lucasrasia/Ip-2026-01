package main
import "fmt"
func main(){
	soma:=0
	for i:=50; i<=70; i=i+2{
		soma=soma+i
	}
	
	fmt.Printf("Somatório: %v\n", soma)
	fmt.Printf("Média: %v \n", (soma/11)) 
}
