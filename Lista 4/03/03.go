package main 
import (
	"fmt"
	"strconv"
)

func main(){
	var num[]int
	var a string
	for i:=1;i<=5;{
		fmt.Print("Digite: ")
		fmt.Scan(&a)
		n, err:=strconv.Atoi(a)
		if err!=nil{
			fmt.Print("Digite um número inteiro!\n")
			continue
		}
		num=append(num, n)
		i++
	}
	var soma_par, soma_imp int
}