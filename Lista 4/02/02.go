package main
import "fmt"
func main(){
	var arr1[]int
	var arr2[]int
	var a int
	fmt.Println("Vetor 1") //título vetor 1
	for i:=1;i<=10;i++{
		fmt.Scanln(&a)
		arr1=append(arr1, a)
	} 
	fmt.Println("Vetor 2") //título vetor 2
	for i:=1;i<=5;i++{
		fmt.Scanln(&a)
		arr2=append(arr2, a)
	}
	var r1[]int  // vetores resultantes 1 e 2
	var r2[]int
	var soma int
	for _, i:=range arr2{  //soma dos termos do vetor 2
		soma+=i
	}
	for _, i:=range arr1{
		soma+=i
		if i%2==0{
			r1=append(r1, soma)
		} else{
			r2=append(r2, soma)
		}
		soma-=i
	}
	fmt.Println("Vetor resutante 1:", r1)
	fmt.Println("Vetor resutante 2:", r2)
}