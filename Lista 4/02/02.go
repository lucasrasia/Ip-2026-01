package main
import "fmt"
func main(){
	var arr1[]int
	var arr2[]int
	var a int
	fmt.Println("Vetor 1")
	for i:=1;i<=10;i++{
		fmt.Scanln(&a)
		arr1=append(arr1, a)
	} 
	fmt.Println("Vetor 2")
	for i:=1;i<=5;i++{
		fmt.Scanln(&a)
		arr2=append(arr2, a)
	}
	var r1[]int
	var r2[]int
	var soma int
	for _, i:=range arr2{
		soma+=i
	}
	for _, i:=range arr1{
		if i%2==0{
			soma2:=soma
			soma2+=i
			r1=append(r1, soma2)
		} else{
			r2=append(r2, soma2)
		}
	}
	fmt.Println("Vetor resutante 1:", r1)
	fmt.Println("Vetor resutante 2:", r2)
}