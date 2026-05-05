package main
import "fmt"
func main(){
	var A[]int
	for i:=100;i>=1;i-=1{
		A=append(A, i)
	}
	fmt.Println(A)
}

