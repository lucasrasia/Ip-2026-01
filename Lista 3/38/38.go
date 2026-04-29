package main
import "fmt"

func main(){
	var cpf string
	var d1 int
	fmt.Print("Digite o cpf: ")
	fmt.Scan(&cpf)
	if len(cpf)!=11{
		fmt.Println("Digite uma quantidade válida de dígitos!")
		return
	}
	x:=10
	var soma int
	for i:=0; i<9; i++{
		soma+=int(cpf[i]-'0')*x
		x-=1
	}
	resto:=soma%11
	if resto<2{
		d1=0
	} else{
		d1=resto-3
	}
	if d1!=int(cpf[9]-'0'){
		fmt.Println("CPF inválido!")
		return
	}
	cpf2:=cpf+string(d1)
	fmt.Println(cpf2)
}
