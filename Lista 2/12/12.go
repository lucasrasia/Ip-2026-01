package main
import "fmt"

func main(){
	var idade int
	fmt.Print("Idade: ")
	fmt.Scan(&idade) 
	if idade<0 {
		fmt.Println("Digíte uma idade válida")
		return
	}
	if idade<3 {
		fmt.Println("Recém-nascido")
	} else if idade<12 {
		fmt.Println("Criança")
	} else if idade<20 {
		fmt.Println("Adolescente")
	} else if idade<56 {
		fmt.Println("Adulto")
	} else{
		fmt.Println("Idoso")
	}
}