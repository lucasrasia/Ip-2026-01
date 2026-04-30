package main 
import (
	"fmt"
	"strconv"
)

func main(){
	var num[]int
	var a string
	for i:=1;i<=10;{
		fmt.Print("Digite: ")
		fmt.Scan(&a)
		n, err:=strconv.Atoi(a)
		if err!=nil{
			fmt.Print("ERRO! Digite um número inteiro!\n")
			continue
		}
		num=append(num, n)
		i++
	}
	var soma_par, soma_imp int
	var par[] int
	var imp[] int
	for _, i:=range num{
		if i%2==0{
			par=append(par, i)
			soma_par+=i
		} else{
			imp=append(imp, i)
			soma_imp+=i
		}
	}
	fmt.Println()
	fmt.Println("Números pares:", par, " Somatório:", soma_par)
	fmt.Println("Números ímpares:", imp, " Somatório:", soma_imp)
}